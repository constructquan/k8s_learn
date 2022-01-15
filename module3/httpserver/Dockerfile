FROM golang:1.16.5-alpine
RUN mkdir /app
ADD . /app/ 
WORKDIR /app
RUN go build -o httpserver .
EXPOSE 8080
CMD  ["/app/httpserver"]

