FROM golang:1.22.0-alpine

WORKDIR /app

COPY . .

RUN go build -o my-first-api .

EXPOSE 8080

CMD ["./my-first-api"]
