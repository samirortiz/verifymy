FROM golang:latest

RUN mkdir /app
WORKDIR /app
COPY . .
RUN CGO_ENABLED=1 go build -a -ldflags "-linkmode external -extldflags '-static' -s -w" -o main main.go

CMD ["/app/main"]