##
## stage 1 : Build
##
FROM golang:1.11 AS build

# git clone https://github.com/glenux/trello2mail /app
COPY . /app

WORKDIR /app
RUN CGO_ENABLED=0 go build -v ./...


##
## Stage 2 : Runtime
## 
FROM alpine:3.7 AS runtime
RUN apk update \
    && apk add --no-cache dcron tzdata ca-certificates bash \
    && rm -rf /var/cache/apk/*

RUN mkdir /app
COPY --from=build /app/trello2mail /app/
COPY docker/trello2mail.cron /app/
COPY templates/ /app/

CMD ["/bin/sh", "-c", "/usr/bin/crontab /app/trello2mail.cron && /usr/sbin/crond -f"]

