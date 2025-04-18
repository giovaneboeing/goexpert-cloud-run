FROM golang:1.23 as builder
WORKDIR /app
COPY . .

#RUN go test ./...
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o server -ldflags="-w -s" cmd/cloudrun/main.go
RUN apt-get update && apt-get install -y ca-certificates

FROM scratch
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

COPY --from=builder /app/server /server
CMD ["/server"]