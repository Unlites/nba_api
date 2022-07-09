FROM golang:1.16-buster

RUN go version
ENV GOPATH=/

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY ./ ./

RUN apt-get update
RUN apt-get -y install postgresql-client

RUN chmod +x wait_for_postgres.sh

RUN go build -o nba_api ./cmd/nba_api/main.go

EXPOSE 8000

CMD ["./nba_api"]