FROM golang:1.10 AS build

RUN go get github.com/glenux/trello2mail-go/...
WORKDIR /go/src/github.com/glenux/trello2mail-go
RUN CGO_ENABLED=0 go build ./...

FROM alpine:3.7
RUN apk update \
    && apk add --no-cache dcron tzdata ca-certificates \
    && rm -rf /var/cache/apk/*

RUN mkdir /app
COPY --from=build /go/src/github.com/glenux/trello2mail-go/trello2mail /app/
COPY docker/trello2mail.cron /app/

CMD ["/bin/sh", "-c", "/usr/bin/crontab /app/trello2mail.cron && /usr/sbin/crond -f -d"]

