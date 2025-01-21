FROM golang:alpine

WORKDIR /app
COPY . /app
EXPOSE 80

RUN go mod download

RUN mkdir -p logs \
    && touch logs/app.log

RUN go build -o service

CMD ["./service"]