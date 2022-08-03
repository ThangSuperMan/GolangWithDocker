# syntax=docker/dockerfile:1
FROM golang:alpine
RUN mkdir /app
WORKDIR /app 
COPY . /app

RUN go get -d github.com/gorilla/mux

# . Stand for WORKDIR
# RUN go build -o main .
# CMD ["/app/main"]

CMD ["go","run","main.go"]
