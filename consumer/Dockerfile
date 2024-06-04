FROM golang:1.20

WORKDIR /app/

RUN apt-get update && apt-get install -y librdkafka-dev

COPY . .

RUN GOOS=linux go build -ldflags="-w -s" -o server cmd/consumer/main.go

CMD ["./server"]
