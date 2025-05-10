FROM golang:1.24.2 AS builder
WORKDIR /build
COPY go.mod ./
COPY cmd ./cmd 
COPY pkg ./pkg 
RUN CGO_ENABLED=0 GOOS=linux go build -o app ./cmd

FROM alpine:3.21.2
WORKDIR /applications
COPY --from=builder /build/app /applications/app
EXPOSE 8080
ENTRYPOINT ["./app"]
