FROM golang:1.23.2-alpine

# Встановлюємо необхідні залежності для CGO та SQLite
RUN apk add --no-cache git gcc musl-dev sqlite-dev build-base

# Встановлюємо air
RUN go install github.com/air-verse/air@latest

# Увімкнення CGO
ENV CGO_ENABLED=1

# Налаштування для ARM64
ENV GOOS=linux
ENV GOARCH=arm64

# Налаштування робочого каталогу
WORKDIR /app

# Копіюємо файли go.mod і go.sum, завантажуємо залежності
COPY go.mod go.sum ./
RUN go mod download

# Копіюємо весь проект
COPY . .

# Запускаємо air
CMD ["air"]
