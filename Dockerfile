FROM golang:latest as builder
WORKDIR /app
ENV GOPROXY https://goproxy.cn
COPY go.mod go.sum ./
RUN go mod tidy
COPY . .

RUN CGO_ENABLED=1 GOOS=linux go build -a -installsuffix cgo -o ip .

FROM debian:latest
WORKDIR /app
ENV HTTP_OPERATION_PORT 3000
COPY --from=builder /app/ip .
EXPOSE ${HTTP_OPERATION_PORT}
ENTRYPOINT ["./ip"]
