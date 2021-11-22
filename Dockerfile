FROM golang:1.16-alpine

RUN mkdir -p /DDD

WORKDIR /DDD

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o /DDD

EXPOSE 8080

CMD [ "./DDD" ]