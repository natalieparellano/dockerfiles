FROM golang

COPY . /tmp/dockerfiles

RUN mkdir /kaniko && cd /tmp/dockerfiles && GOOS=linux go build -o /kaniko/dockerfiles .

ENV PATH=/usr/local/bin:/kaniko
ENV HOME=/root
ENV USER=root

ENTRYPOINT /kaniko/dockerfiles
