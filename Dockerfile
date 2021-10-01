FROM golang:1.17.0-alpine3.14
RUN mkdir /app
ADD ginproject /app
WORKDIR /app
RUN go build -o main .
CMD ["/app/main"]