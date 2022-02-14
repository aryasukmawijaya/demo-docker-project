FROM golang:1.17.7-alpine3.15

RUN apk add --no-cache git

WORKDIR /app-go

COPY . .

RUN go mod tidy

RUN go build -o binary

ENTRYPOINT ["/app-go/binary"]