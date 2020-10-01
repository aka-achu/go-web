FROM golang:latest as builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o go-web ./main.go

FROM ubuntu:latest
WORKDIR /root/
COPY --from=builder /app/go-web  .
COPY --from=builder /app/.env  .
COPY --from=builder /app/cert/  ./cert
EXPOSE 8080
CMD ["./go-web"]
