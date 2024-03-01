FROM golang:alpine

WORKDIR /src/usr/app

COPY ./go.* ./

RUN go mod download

COPY . .

RUN go build -o goapp .

EXPOSE 8000

CMD ["/src/usr/app/goapp"]
