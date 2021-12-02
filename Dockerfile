FROM golang:1.16-alpine3.14

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o mainfile

EXPOSE 1234

CMD ["./mainfile"]