FROM golang:1.14

RUN mkdir /web/

COPY ./dlv /usr/local/bin/
COPY ./testapp/main.go /go/src
COPY ./start.sh  /web/
RUN chmod   u+x /web/start.sh
WORKDIR /go/src/
ENTRYPOINT ["/web/start.sh"]