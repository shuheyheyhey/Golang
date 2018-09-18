FROM golang
MAINTAINER shuhei.yukawa

ENV WD ./
WORKDIR $WD

ADD WD/src /go/src/

# Install dependencies
RUN go get github.com/gorilla/mux
RUN go get github.com/aws/aws-sdk-go

EXPOSE 6000
CMD ["/usr/local/go/bin/go", "run", "/go/src/main.go"]
