FROM golang:latest
## todo  нужно сборку тоже сделать через докер файл
## работа с миграциями
WORKDIR /app
COPY . .
RUN go mod download

RUN go build -o bin/main cmd/api/main.go

RUN go install github.com/pressly/goose/v3/cmd/goose@latest