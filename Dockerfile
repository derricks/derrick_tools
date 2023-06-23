FROM golang:1.20
WORKDIR /go/src/derrick_tools
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build
