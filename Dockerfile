FROM golang:1.22

WORKDIR /app

RUN apt-get update && apt-get install -y git
RUN go install github.com/air-verse/air@latest
RUN go install github.com/rubenv/sql-migrate/...@latest

COPY go.mod go.sum ./
RUN go mod download

CMD ["air"]
