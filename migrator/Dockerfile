FROM golang:1.20

WORKDIR /app/

COPY . .

RUN GOOS=linux go build -ldflags="-w -s" -o migrator main.go

CMD ["./migrator"]
