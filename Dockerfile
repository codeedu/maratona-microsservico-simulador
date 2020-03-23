FROM golang:1.13
COPY / /

WORKDIR /

RUN go build simulator.go

ENTRYPOINT ["./simulator"]