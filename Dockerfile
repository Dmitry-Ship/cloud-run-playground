FROM golang:1.15.2 

WORKDIR /app

COPY go.mod .

RUN go mod download

COPY . .

RUN go build -o main .

CMD ["/app/main"]