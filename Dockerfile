FROM golang:1.23.2-alpine

RUN apk add --no-cache git gcc musl-dev sqlite-dev build-base

RUN go install github.com/air-verse/air@latest

ENV CGO_ENABLED=1

ENV GOOS=linux
ENV GOARCH=arm64

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

CMD ["air"]
