FROM golang:1.23.2-alpine

RUN apk add --no-cache git gcc musl-dev mysql-dev build-base

RUN go install github.com/air-verse/air@latest

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN git init
ENV GOFLAGS="-buildvcs=false"

CMD ["air"]
