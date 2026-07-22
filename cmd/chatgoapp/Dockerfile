FROM golang:1.23.4-bookworm AS builder
WORKDIR /chatgoapp
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o /chatgoapp/exe ./cmd/chatgoapp

FROM alpine:3.23
WORKDIR /chatgoapp
COPY --from=builder /chatgoapp/exe /chatgoapp/exe
ENV PROJECT_ROOT=/chatgoapp
CMD ["/chatgoapp/exe"]