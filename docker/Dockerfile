##
## stage 1 : Build
##
FROM golang:1.12 AS build

RUN git clone                                    \
        --single-branch                          \
        --branch develop                         \
        https://github.com/glenux/musala         \
        /app                                     \
    && rm -f /app/contrib-trello                 \
    && git clone                                 \
        https://github.com/glenux/contrib-trello \
        /app/contrib-trello

# COPY . /app

WORKDIR /app

RUN go get                                       \
	&& CGO_ENABLED=0 make build


##
## Stage 2 : Runtime
## 
FROM alpine:3.7 AS runtime
RUN apk update \
    && apk add --no-cache dcron tzdata ca-certificates bash \
    && rm -rf /var/cache/apk/*

RUN mkdir /app
COPY --from=build /app/musala /app/
COPY docker/musala.cron /app/
COPY templates/ /app/templates

CMD ["/bin/sh", "-c", "/usr/bin/crontab /app/musala.cron && /usr/sbin/crond -f"]

