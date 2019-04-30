[![Docker Repository on Quay](https://quay.io/repository/themoah/go-http-echo/status "Docker Repository on Quay")](https://quay.io/repository/themoah/go-http-echo)

# go-http-echo

A simple golang HTTP server to echo requests

## How to use

    go run echo.go

This will start the HTTP server running on localhost, port 8080. 
To run on a different port, set the `SERVER_PORT` environment variable before running.

Once running, you can make requests from a browser (or any HTTP client) via `http://localhost:8080/`

## How to build

Golang build:

    go build -o main

Docker build(multistage):

    docker build -t themoah/go-http-echo:test .

## How to run in a docker

(Don't forget to build it first)

    docker run -p 8080:8080 -d themoah/go-http-echo:test

(logs are written to the stdout, to see them run without -d)
