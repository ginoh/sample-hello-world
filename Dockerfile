FROM golang as builder

WORKDIR /go/src/sample-hello-world

COPY hello-world.go go.mod ./
RUN go mod download && \
    go install

FROM ubuntu

EXPOSE 8080

COPY --from=builder /go/bin/sample-hello-world /usr/local/bin
CMD ["sample-hello-world"]
