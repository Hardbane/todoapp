FROM golang:1.26.2-bookworm AS builder
WORKDIR /application
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o /application/cmd/todoapp/exe /application/cmd/todoapp

FROM alpine:3.23
WORKDIR /application
COPY --from=builder /application/cmd/todoapp/exe /application
CMD ["/application/exe"]
