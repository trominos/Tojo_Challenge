FROM golang:latest AS builder
WORKDIR /app
COPY . .
COPY crypto /app/crypto
COPY static /app/static
COPY  web_server.go /app
RUN ls -la    # Print directory contents for debugging
RUN go build -o http_server web_server.go
RUN ls -la    # Print directory contents for debugging
EXPOSE 80
EXPOSE 443
CMD ["./http_server"]
