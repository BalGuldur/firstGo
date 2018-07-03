#FROM golang
#
#ADD . /go/src/devices_go
#
#RUN go install devices_go
#
#ENTRYPOINT /go/bin/devices_go
#
#EXPOSE 8085

FROM golang

WORKDIR /go/src/devices
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["devices"]