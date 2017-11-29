FROM golang:latest 
RUN mkdir /app 
ADD . /app/ 
WORKDIR /app 
RUN go build -o main . 
ENV HTTP_PORT 8080
CMD ["/app/main"]
