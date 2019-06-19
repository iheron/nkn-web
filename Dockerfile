FROM golang:1.12.4-alpine
LABEL maintainer="peng.liu@nkn.org"

ADD ./bin/ /go/src/nkn-web
WORKDIR /go/src/nkn-web

EXPOSE 8080

CMD ["./app_go"]
