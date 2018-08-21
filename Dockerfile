FROM golang:1.10 AS build
COPY . /go/src/github.com/glenux/trello2mail-go
WORKDIR /go/src/github.com/glenux/trello2mail-go
RUN CGO_ENABLED=0 go build -o trello2mail main.go

FROM alpine:3.7
RUN apk add --no-cache ca-certificates
COPY --from=build /go/src/github.com/glenux/trello2mail-go/ /usr/bin/trello2mail
CMD ["./main"]
