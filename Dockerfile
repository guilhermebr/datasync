FROM golang:1.4.2-wheezy
MAINTAINER Guilherme Rezende <guilhermebr@gmail.com> (@gbrezende)


WORKDIR /go/src/github.com/guilhermebr/datasync
COPY . /go/src/github.com/guilhermebr/datasync

RUN go get github.com/tools/godep
RUN godep get
RUN go install .

RUN go test
WORKDIR /go/src/github.com/guilhermebr/datasync/storages
RUN go test -cashost=172.17.42.1 -eshost=172.17.42.1

WORKDIR /go/src/github.com/guilhermebr/datasync
CMD ["datasync"]
