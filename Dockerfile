FROM golang:1.21 AS builder

WORKDIR /FileEventLogger

COPY . .
RUN go mod download
RUN go build -o main .

FROM golang:1.21

WORKDIR /FileEventLogger

COPY --from=builder /FileEventLogger/main .

ENTRYPOINT ["./main"]
