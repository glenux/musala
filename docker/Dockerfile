FROM golang:1.11 AS build

RUN git clone https://github.com/glenux/trello2mail /app

WORKDIR /app
RUN CGO_ENABLED=0 go build -v ./...

FROM alpine:3.7
RUN apk update \
    && apk add --no-cache dcron tzdata ca-certificates \
    && rm -rf /var/cache/apk/*

RUN mkdir /app
COPY --from=build /app/trello2mail /app/
COPY --from=build /app/docker/trello2mail.cron /app/

CMD ["/bin/sh", "-c", "/usr/bin/crontab /app/trello2mail.cron && /usr/sbin/crond -f"]

