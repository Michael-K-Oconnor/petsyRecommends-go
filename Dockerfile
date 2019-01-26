FROM golang:1.8 AS builder

WORKDIR /go/src/staging

COPY . /go/src/staging

RUN go get -d -v ./...
RUN go install -v ./...

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o app .

FROM scratch
WORKDIR /root/
COPY --from=builder /go/src/staging/app .

EXPOSE 8080
ENTRYPOINT ["./app"]