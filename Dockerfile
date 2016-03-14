FROM golang

ADD . /go/src/github.com/dfernandez/geb

WORKDIR /go/src/github.com/dfernandez/geb

RUN go get ./...
RUN go install .

ENTRYPOINT /go/bin/geb
