# Trello-2-Mail

Every morning, get the content of your favorite task list by email


## Installation

### With go

Make sure you have Go installed, then type:

    go install github.com/glenux/trello2mail-go/...

It will install Trello2Mail binary in `$GOPATH/bin`

### With docker

Make sure you have Docker installed, then type:

    docker build -t trello2mail -f docker/Dockerfile .

## Usage

## Creating a developper account

1. Create a trello account
2. Confirm your email
3. Enable developper account on <https://trello.com/app-key>
4. Get an developer API KEY

## Getting a Trello TOKEN

Open the following URL in your web browser and authenticate yourself. That will
give you the TRELLO_TOKEN that will be needed in the next step.

<https://trello.com/1/authorize?expiration=never&scope=read,write,account&response_type=token&name=Trello2Mail&key=YOUR-API-KEY>

## Normal use

Prepare your environment with the following variables

```
EMAIL_FROM:    no-reply@example.com
EMAIL_TO:      me@example.com
EMAIL_SUBJECT: "Daily mail for YYYYYY"
TRELLO_URL:    https://trello.com/b/xxxxx/yyyy
TRELLO_TOKEN:  xxxxxxxxxxxxxx
SMTP_HOSTNAME: smtp.example.com
SMTP_USERNAME: foobar@example.com
SMTP_PASSWORD: securefoobar
SMTP_PORT:   587
# SMTP_AUTH_TYPE accepts either "none", "plain" or "login"
SMTP_AUTH_TYPE: plain 
# SMTP_SECURITY_TYPE accepts either "none", "tls" or "starttls"
SMTP_SECURITY_TYPE: tls
```

Then run the program:

    $GOPATH/bin/trello2mail

### With docker

    docker run  \
        -e EMAIL_FROM=
        -e EMAIL_TO=
        -e EMAIL_SUBJECT=
        -e EMAIL_FROM=
        -e EMAIL_TO=
        -e EMAIL_SUBJECT=
        -v /var/lib/trello2mail/trello2mail.cron:/app/trello2mail.cron \
        -it trello2mail:latest

## Contributing

1. Fork it ( http://github.com/glenux/trello2mail-go/fork )
2. Create your feature branch (`git checkout -b my-new-feature`)
3. Commit your changes (`git commit -am 'Add some feature'`)
4. Push to the branch (`git push origin my-new-feature`)
5. Create new Pull Request


## Credits

Author & Maintainer: Glenn Y. ROLLAND

Contributors: none yet ;)

Got questions? Need help? Tweet at @glenux


## License

Trello2Mail is Copyright © 2018 Glenn ROLLAND. It is free software, and may be redistributed under the terms specified in the LICENSE file.
