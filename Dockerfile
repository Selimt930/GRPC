FROM golang:1.15 as builder
# Move to working directory /build
WORKDIR /build
# Copy and download dependency using go mod
COPY ./auth /build/auth
COPY ./server /build/server
COPY ./service /build/service
COPY ./go.mod /build/go.mod
COPY ./go.sum /build/go.sum
RUN CGO_ENABLED=0 GOOS=linux go build -o server /build/server/server.go

FROM alpine:latest
WORKDIR /root
COPY --from=builder /build/server .
CMD ["./server"]

