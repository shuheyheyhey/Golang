FROM golang
MAINTAINER shuhei.yukawa
ADD . /go/src/server
RUN go install server

FROM alpine

ENV WD ./
WORKDIR $WD

COPY --from=0 /go/bin/server $WD/
ENV PORT 6000
CMD ["./server"]
