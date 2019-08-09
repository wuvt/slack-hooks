FROM golang:alpine3.8 AS build

COPY . /go/src/github.com/wuvt/slack-hooks
WORKDIR /go/src/github.com/wuvt/slack-hooks

RUN set -ex \
        && apk add --no-cache --virtual .build-deps git \
        && go get -v . \
        && apk del .build-deps

FROM alpine:3.8

COPY --from=build /go/bin/slack-hooks /usr/local/bin/slack-hooks

EXPOSE 8080
USER nobody
ENTRYPOINT ["/usr/local/bin/slack-hooks"]
