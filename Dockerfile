FROM golang:latest 
ARG http_port=:9000
ENV HTTP_PORT=$http_port
EXPOSE $http_port
RUN mkdir /app 
ADD . /app/ 
WORKDIR /app 
RUN go build -o main . 
CMD ["/app/main"]
