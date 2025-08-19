FROM golang:1.24.6 AS builder
WORKDIR /build
COPY go.mod .
COPY . .
RUN go build -o /build/gotp ./cmd/gotp && chmod +x /build/gotp

FROM alpine:3.22.1
RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2
WORKDIR app
COPY --from=builder /build/gotp /app/gotp
COPY ./assets/scripts /app/assets/scripts
CMD ["./gotp"]