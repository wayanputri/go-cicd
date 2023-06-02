
FROM golang:1.20-alpine

RUN mkdir /app
WORKDIR /app

COPY go.mod go.sum /app/
RUN go mod download

COPY . /app/

RUN go build -o beapi

CMD ["./beapi"]
