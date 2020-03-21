FROM golang:1.13 AS builder

ENV GO111MODULE=on
ENV PORT=30000
WORKDIR /app
COPY go.mod .
COPY go.sum .

RUN go mod download
COPY . .

RUN go build 

FROM ubuntu
WORKDIR /app
COPY --from=builder /app/main .
CMD ["./main"]
