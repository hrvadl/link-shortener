FROM golang:1.21
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
WORKDIR /app/cmd
RUN go build -o /link-shortener
EXPOSE 3000
CMD [ "/link-shortener" ]
