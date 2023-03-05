FROM mcr.microsoft.com/devcontainers/go:0-1-bullseye

RUN apt-get update && apt install -y protobuf-compiler && protoc --version

RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28 && go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2  && go install github.com/fullstorydev/grpcurl/cmd/grpcurl@v1.8.7 && go install github.com/google/ko@v0.12.0

RUN export PATH="$PATH:$(go env GOPATH)/bin"