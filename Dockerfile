# Builder image
FROM golang:alpine as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o build/userservice ./cmd/api


# Build final image
FROM alpine:latest
# RUN apk add --no-cache ca-certificates

WORKDIR /app

COPY --from=builder /app/build/userservice .
# COPY --from=builder /app/.env .

EXPOSE 8080

ENTRYPOINT [ "./userservice" ]