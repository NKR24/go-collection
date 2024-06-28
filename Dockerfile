FROM golang:1.22-alpine3.20

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

COPY ./config/config.json .

RUN go build -o server ./cmd/server

EXPOSE 8080

CMD ["./server"]

