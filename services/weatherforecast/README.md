# Weather Forecast Service

## Required environment variables

- PORT
- METEOSOURCE_API_ENDPOINT
- METEOSOURCE_API_ACCESS_KEY
- REDIS_CONNECTION_STRING


## How to run

Must have required environment variables that are aforementioned after that,

    go run main.go

## How to dockerize

We use ko to dockerize for the our go applications

    KO_DOCKER_REPO=ghcr.io/recepdmr/news-feed/services ko build