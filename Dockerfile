FROM golang:1.22-alpine

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go get github.com/lib/pq

RUN go build -o main ./cmd/server

EXPOSE 8080

CMD ["/app/main"]