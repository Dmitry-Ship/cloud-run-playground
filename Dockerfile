FROM golang:1.16.4

WORKDIR /app

COPY go.* ./

RUN go mod download

COPY . ./

RUN go build -v -o main

ENV PORT 3000
ENV HOST 0.0.0.0

EXPOSE 3000

CMD ["/app/main"]