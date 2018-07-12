FROM golang:latest
RUN mkdir /customer-go
ADD . /customer-go/
WORKDIR /customer-go
RUN go build -o main .
CMD ["/customer-go/"]

EXPOSE 8090
