FROM golang as builder

EXPOSE 8080

COPY hello-world.go hello-world.go
RUN go build -o bin/hello-world

FROM ubuntu
COPY --from=builder /go/bin/hello-world /usr/local/bin
CMD ["hello-world"]
