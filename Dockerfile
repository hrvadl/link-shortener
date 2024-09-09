FROM golang:1.21 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
WORKDIR /app/cmd/server
RUN go build -o ./server

FROM scratch AS final
COPY --from=builder /app/cmd/server/server /server
EXPOSE 3000
CMD ["/server"]
