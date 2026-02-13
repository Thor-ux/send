FROM golang:1.22-alpine AS builder

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -buildvcs=false -o mailer

FROM alpine:3.19

WORKDIR /app

COPY --from=builder /app/mailer /app/mailer
COPY --from=builder /app/templates /app/templates

RUN adduser -D mailer
USER mailer

EXPOSE 80

CMD ["/app/mailer"]

