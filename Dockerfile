FROM golang:1.8-alpine AS builder
WORKDIR /go/src/github.com/edgexfoundry/exportclient
RUN apk add --no-cache git make
COPY . .
RUN cd cmd && make

FROM scratch
COPY --from=builder /go/src/github.com/edgexfoundry/exportclient/expotclient /
EXPOSE 7070
ENTRYPOINT ["/exportclient"]
