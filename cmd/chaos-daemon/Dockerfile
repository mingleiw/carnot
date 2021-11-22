# Build the manager binary
FROM golang:1.16 as builder

RUN apt-get update
RUN apt-get install -y libpcap-dev

WORKDIR /chaos-daemon

# Copy the go source
COPY ./ ./

RUN go mod download

# Build
RUN go build .

ENTRYPOINT ["./chaos-daemon"]
