FROM golang:1.9

WORKDIR /go/src/app

COPY . .

RUN go-wrapper install

EXPOSE 8000

ENTRYPOINT ["go-wrapper", "run"] # ["app"]
