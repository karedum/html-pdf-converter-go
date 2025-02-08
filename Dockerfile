FROM golang:alpine

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o main ./cmd

ENV PORT=8081

EXPOSE ${PORT}

CMD ["./main"]