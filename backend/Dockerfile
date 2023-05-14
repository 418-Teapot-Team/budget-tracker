FROM golang:1.19

WORKDIR /dockerapp

COPY . .

RUN go mod download && go mod verify

RUN go build -o /main

EXPOSE 80

CMD ["/main"]