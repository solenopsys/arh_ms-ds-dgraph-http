FROM --platform=$BUILDPLATFORM  golang:buster as builder

ARG TARGETARCH

RUN apt-get update && apt-get install -qy build-essential software-properties-common libjemalloc2 sudo

ENV CGO_ENABLED=0 GOOS=linux GOARCH=$TARGETARCH
RUN git clone https://www.github.com/dgraph-io/dgraph/ && \
    cd dgraph && \
    make install

RUN ls -alh ${GOPATH}/bin

RUN mkdir -p /dist/bin && \
    mkdir -p /dist/tmp && \
    if [ "$TARGETARCH" = "amd64" ]  ; then mv  ${GOPATH}/bin/dgraph /dist/bin/dgraph ; else mv  ${GOPATH}/bin/linux_$TARGETARCH/dgraph /dist/bin/dgraph ; fi

FROM alpine:latest as dgraph
COPY --from=builder /dist /
ENV PATH=$PATH:/bin/
RUN chmod +x /bin/dgraph && apk --update --no-cache add bash


CMD ["/bin/dgraph", "version"]
