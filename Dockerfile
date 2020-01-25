# Build
FROM golang:alpine as builder
ENV GO111MODULE=on
RUN apk update && apk add --no-cache git
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o slingshot .

# Serve
FROM alpine:latest
ARG PORT
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/slingshot .
COPY --from=builder /app/config.yaml .
EXPOSE ${PORT}
ENTRYPOINT ["./slingshot"]
