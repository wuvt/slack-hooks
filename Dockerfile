FROM golang:1.12-alpine3.10 AS build

COPY . /usr/src/github.com/wuvt/slack-hooks
WORKDIR /usr/src/github.com/wuvt/slack-hooks

RUN set -ex \
        && apk add --no-cache --virtual .build-deps git \
        && go get -v . \
        && apk del .build-deps

FROM alpine:3.10

RUN apk add --no-cache ca-certificates

COPY --from=build /go/bin/slack-hooks /usr/local/bin/slack-hooks

EXPOSE 8080
USER nobody
ENTRYPOINT ["/usr/local/bin/slack-hooks"]
