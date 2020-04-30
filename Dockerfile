FROM golang:1.13.4 

WORKDIR $GOPATH/src/github.com/DylanLovesCoffee/dogstatsd-go
COPY . .
RUN go get -d -v ./...
RUN go install -v ./...
LABEL app=dogstatsd-go
CMD ["dogstatsd-go"]