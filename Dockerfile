FROM golang:latest
WORKDIR /app
ADD . /app/
RUN go build -o main .

FROM alpine:latest
ARG http_port=:9000
ENV HTTP_PORT=$http_port
EXPOSE $http_port
WORKDIR /root
COPY --from=0 /app/main .
CMD ["/root/main"]


