FROM golang:latest
RUN mkdir /app
ADD . /app/
WORKDIR /app
RUN go mod init backend
RUN go build -o main .
EXPOSE 3000
CMD ["/app/main"]
