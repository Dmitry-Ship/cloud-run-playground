# stage 1: install dependencies
FROM golang:1.16.4-alpine AS base

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

# stage 2: build binary for production
FROM base as build

RUN go build -v -o main .

ENV PORT=${PORT}

FROM alpine:latest as prod

RUN apk --no-cache add ca-certificates

COPY --from=build /app .

CMD ["./main"] 

# stage 4: install dev dependencies 
FROM base as dev

RUN go get github.com/codegangsta/gin

CMD ["gin" , "run", "main.go"]  