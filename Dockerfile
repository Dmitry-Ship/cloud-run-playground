FROM golang:1.15.2 

WORKDIR /app

COPY go.mod .

RUN go mod download

COPY . .

RUN go build -o main .

ENV PORT=3000
ENV HOST=localhost

EXPOSE 3000

CMD ["/app/main"]