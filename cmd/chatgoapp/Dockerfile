FROM golang:1.26.3-bookworm AS builder
WORKDIR /chatgoapp
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o /app/cmd/chatgoapp/exe /app/cmd/chatgoapp

FROM alpine:3.23
WORKDIR /chatgoapp
COPY --from=builder /app/cmd/chatgoapp/exe /app
COPY --from=builder /app/public /app/public
ENV PROJECT_ROOT=/chatgoapp
CMD ["/chatgoapp/exe"]