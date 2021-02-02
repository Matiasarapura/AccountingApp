# Build from golang latest image
FROM golang:latest as builder
WORKDIR /accounting
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o accounting .

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /accounting/accounting ./
COPY --from=builder /accounting/static ./static
EXPOSE 8080
ENTRYPOINT ["./accounting"]
