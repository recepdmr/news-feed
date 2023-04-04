# Public apigateway

## Required environment variables

- PORT
- POST_SERVICE_URI

## How to run

Must have required environment variables that are aforementioned after that,

    go run main.go

## How to dockerize

We use ko to dockerize for the our go applications

    cd ./cmd && KO_DOCKER_REPO=ghcr.io/recepdmr/news-feed/apigateways/public ko build 