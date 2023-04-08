FROM golang:alpine

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

COPY ./env/production.env ./.env

RUN rm -rf ./env

RUN go build -o ./server

CMD [ "./server", "serve" ]