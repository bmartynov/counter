FROM golang:1.10.1-alpine3.7

ENV ENVCONSUL_VERSION=0.7.3
ENV ENVCONSUL_RELEASES=https://releases.hashicorp.com/envconsul

RUN apk update
RUN apk add git

RUN mkdir -p /tmp/build && cd /tmp/build

RUN wget ${ENVCONSUL_RELEASES}/${ENVCONSUL_VERSION}/envconsul_${ENVCONSUL_VERSION}_linux_amd64.tgz && \
    tar -xzf envconsul_${ENVCONSUL_VERSION}_linux_amd64.tgz envconsul -C /usr/bin/

RUN go get github.com/bmartynov/counter/...
RUN ls /go/bin

FROM alpine:latest
COPY --from=0 /go/bin/server /bin/server
COPY ./entrypoint.sh /docker-entrypont.sh
RUN chmod +x /docker-entrypont.sh


ENTRYPOINT ["/docker-entrypont.sh"]