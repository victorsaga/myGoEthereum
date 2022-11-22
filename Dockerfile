FROM golang:1.18.2

WORKDIR /app

COPY . .

RUN  go mod download

RUN go build -o release .

ENV GIN_MODE release

EXPOSE 8080

CMD ["./release"]

# docker build -t myGoEthereum .
# docker run -d -p 8081:8080 myGoEthereum