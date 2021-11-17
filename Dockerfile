# stage 1: build binary
FROM golang AS builder

COPY . /tmp/dockerfiles

RUN mkdir /kaniko && cd /tmp/dockerfiles && CGO_ENABLED=0 GOOS=linux go build -o /dockerfiles .

# stage 2: copy binary
FROM gcr.io/distroless/static

COPY --from=builder /dockerfiles /dockerfiles

ENTRYPOINT ["/dockerfiles"]
