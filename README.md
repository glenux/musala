# Musala Mail

Every morning, get the content of your favorite kanban board by email

Compatibility:

* :heavy_check_mark: __Trello__
* :x: Gitlab Issues Boards (work in progress)
* :x: Libreboard
* :x: Github Projects
* :x: Nextcloud Deck

## Installation

### With go

Make sure you have Go installed, then type:

    $ go install github.com/glenux/musala-mail/...

It will install Musala Mail binary in `$GOPATH/bin`

### With docker

Make sure you have Docker installed, then type:

    $ docker build -t musala-mail -f docker/Dockerfile .

## Usage

### Creating a developper account

1. Create a Trello account on <https://trello.com>
2. Check your mailbox and confirm your email
3. Enable developper account on <https://trello.com/app-key>
4. Get an developer API KEY

### Getting a Trello TOKEN

Open the following URL in your web browser and authenticate yourself. That will
give you the TRELLO_TOKEN that will be needed in the next step.

<https://trello.com/1/authorize?expiration=never&scope=read,write,account&response_type=token&name=Musala%20Mail&key=YOUR-API-KEY>

### Using the binary

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

    $ $GOPATH/bin/musala-mail

### Using with docker

    $ docker build -f docker/Dockerfile -t musala-mail .
    $ docker run  \
        -e EMAIL_FROM=
        -e EMAIL_TO=
        -e EMAIL_SUBJECT=
        -e EMAIL_FROM=
        -e EMAIL_TO=
        -e EMAIL_SUBJECT=
        -v /var/lib/musala-mail/musala-mail.cron:/app/musala-mail.cron \
        -it musala-mail:latest

## Contributing

1. Fork it ( http://github.com/glenux/musala-mail/fork )
2. Create your feature branch (`git checkout -b my-new-feature`)
3. Commit your changes (`git commit -am 'Add some feature'`)
4. Push to the branch (`git push origin my-new-feature`)
5. Create new Pull Request


## Credits

Author & Maintainer: [Glenn Y. ROLLAND](https://github.com/glenux)

Contributors: none yet ;)

Got questions? Need help? Tweet at [@glenux](https://twitter.com/glenux)


## License

Musala Mail is Copyright © 2018-2019 Glenn ROLLAND. It is free software, and may be redistributed under the terms specified in the LICENSE file.

