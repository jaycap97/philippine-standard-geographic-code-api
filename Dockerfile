FROM golang:1.18.0-alpine3.14 as builder
WORKDIR /root/app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -a -o ./main ./cmd/main.go

FROM alpine:3.14
WORKDIR /root/app
COPY --from=builder /root/app/main .
EXPOSE 8080

CMD ["/root/app/main"]
