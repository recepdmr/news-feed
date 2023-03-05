package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/recepdmr/news-feed/services/weatherforecast/cache"
	"github.com/recepdmr/news-feed/services/weatherforecast/providers/meteosource"
	pb "github.com/recepdmr/news-feed/services/weatherforecast/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	pb.UnimplementedWeatherForecastServer
}

var cacheKeyFormat string = "weatherforecast-%v"

func (s *server) GetWeatherForecastByCityName(ctx context.Context, in *pb.GetWeatherForecastByCityNameRequest) (*pb.GetWeatherForecastByCityNameResponse, error) {

	cacheKey := fmt.Sprintf(cacheKeyFormat, in.GetCityName())

	cacheResult := cache.RedisClient.Get(ctx, cacheKey).Val()
	if cacheResult == "" {
		log.Printf("%v is not in the cache", cacheKey)
		s.FillWeatherForecast(ctx, &pb.FillWeatherForecastRequest{
			CityName: in.GetCityName(),
		})
		cacheResult = cache.RedisClient.Get(ctx, cacheKey).Val()
	} else {
		log.Printf("%v is in the cache", cacheKey)
	}

	var r *meteosource.MeteosourceWeatherForecastResponse
	err := json.Unmarshal([]byte(cacheResult), &r)

	if err != nil {
		errorMessage := fmt.Errorf("unmarshal operation thrown error %v", err.Error())
		log.Fatalf(errorMessage.Error())
		return nil, errorMessage
	}
	return &pb.GetWeatherForecastByCityNameResponse{
		Data: &pb.WeatherForecastResult{
			Result: r.Units,
		},
	}, err
}

func (s *server) FillWeatherForecast(ctx context.Context, in *pb.FillWeatherForecastRequest) (*empty.Empty, error) {

	resp, err := meteosource.GetResult(in.GetCityName())
	if err != nil {
		return nil, err
	}

	cacheKey := fmt.Sprintf(cacheKeyFormat, in.GetCityName())
	cacheValue, err := json.Marshal(resp)

	if err != nil {
		return nil, err
	}

	cache.RedisClient.Set(ctx, cacheKey, cacheValue, 0)

	return &empty.Empty{}, nil
}

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Panic("PORT is waiting")
	}

	if os.Getenv("METEOSOURCE_API_ENDPOINT") == "" {
		panic("meteosource apiEndpoint isn't configured")
	}

	if os.Getenv("METEOSOURCE_API_ACCESS_KEY") == "" {
		panic("meteosource apiAccessKey isn't configured")
	}

	if os.Getenv("REDIS_CONNECTION_STRING") == "" {
		panic("Redis connection string isn't configured")
	}

	log.Printf("Weatherforecast service is available at %v", port)
	listener, err := net.Listen("tcp", fmt.Sprintf(":%v", port))
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	pb.RegisterWeatherForecastServer(s, &server{})
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
