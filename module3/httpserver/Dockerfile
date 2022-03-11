FROM golang:1.16.5 AS build 
WORKDIR /app
ADD . /app/ 
ENV CGO_ENABLED=0 \
    GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct
RUN go get 
RUN GOOS=linux go build -installsuffix cgo -o httpserver hellohttp.go 

FROM busybox
COPY --from=build /app/httpserver /app/httpserver
COPY --from=build /app/online_gracefulstop.sh /app/online_gracefulstop.sh
EXPOSE 8080
WORKDIR /app/
CMD  ["/app/httpserver"]

