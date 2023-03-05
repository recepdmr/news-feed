package meteosource

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func GetResult(cityName string) (*MeteosourceWeatherForecastResponse, error) {

	apiEndpoint := os.Getenv("METEOSOURCE_API_ENDPOINT")
	apiAccessKey := os.Getenv("METEOSOURCE_API_ACCESS_KEY")
	req, err := http.NewRequest("GET",
		fmt.Sprintf("%v/api/v1/free/point?place_id=%v&language=en&units=auto",
			apiEndpoint,
			cityName),
		nil)

	req.Header.Add("X-API-Key", apiAccessKey)

	if err != nil {
		return nil, err
	}
	log.Printf("Requesting to %v", req.URL)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	d, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return UnmarshalMeteosourceWeatherForecastResponse(d)
}
