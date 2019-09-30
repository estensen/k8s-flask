FROM golang:alpine AS builder

WORKDIR $GOPATH/src/club/app
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY main.go .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /go/bin/app


FROM scratch
COPY --from=builder /go/bin/app /go/bin/app

CMD ["/go/bin/app"]
