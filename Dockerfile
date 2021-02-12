FROM golang:1.14
WORKDIR /go/src/derrick_tools
COPY . .
RUN go build
