# Build the manager binary
FROM golang:1.16 as builder

WORKDIR /chaosdaemon

COPY go.mod ./

RUN go mod download

# Copy the go source
COPY *.go ./

# Build
RUN go build -o chaosdaemon main.go

RUN chmod +x chaosdaemon
ENTRYPOINT ["./chaosdaemon"]
