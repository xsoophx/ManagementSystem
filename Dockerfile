FROM golang:1.23.2-alpine3.20

WORKDIR /go/src/app

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

RUN go build -o /go/bin/ManagementSystem .

CMD ["/go/bin/ManagementSystem"]