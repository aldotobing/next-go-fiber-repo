FROM golang:1.13.0

RUN apt-get update && apt-get install -y

ENV PKG_NAME=stellar-2.0-backend/
ENV PKG_PATH=$GOPATH/src/$PKG_NAME
WORKDIR $PKG_PATH/

COPY . $PKG_PATH/

RUN echo $PWD
RUN go mod vendor

WORKDIR $PKG_PATH/amqp_listener_otp/
RUN echo $PWD

RUN go build main.go
CMD ["./main"]