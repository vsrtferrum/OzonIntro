FROM golang:1.23-alpine AS builder

WORKDIR /app

# Копируем файлы зависимостей
COPY go.mod go.sum ./

# Скачиваем зависимости
RUN go mod download

# Копируем исходный код
COPY . .

# Копируем конфигурационный файл
COPY config/config.json ./config/config.json

# Собираем приложение
RUN GOOS=linux go build -o /main ./cmd/server.go

FROM alpine:latest
WORKDIR /root/

# Копируем собранное приложение
COPY --from=builder /main .

# Копируем конфигурационный файл
COPY --from=builder /app/config/config.json ./config/config.json

EXPOSE 8080
CMD ["./main"]