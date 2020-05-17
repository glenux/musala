# Musala

<!-- ![Build](https://github.com/glenux/musala-push/workflows/build/badge.svg?branch=master) -->
<!-- [![Gem Version](https://badge.fury.io/rb/musala-push.svg)](https://rubygems.org/gems/musala-push) -->
[![GitHub license](https://img.shields.io/github/license/glenux/musala-push.svg)](https://github.com/glenux/musala-push/blob/master/LICENSE.txt)
[![Donate on patreon](https://img.shields.io/badge/patreon-donate-orange.svg)](https://patreon.com/glenux)

Every morning, get the content of your favorite kanban board by email, WhatsApp or SMS.

Note: _musala_ means _work, occupation_ [in Lingala](https://dic.lingala.be/fr/mosala)

## Roadmap

Task sources:

* :heavy_check_mark: __Trello__
* :x: Libreboard (not yet)
* :x: Github Projects (not yet)

Delivery via:

* :heavy_check_mark: __E-mail__
* :x: Whatsapp
* :x: SMS


## Installation

### With go

Make sure you have Go installed, then type:

    $ go install github.com/glenux/musala-push/...

It will install Musala Mail binary in `$GOPATH/bin`


### With docker

Make sure you have Docker installed, then type:

    $ docker build -t musala-push -f docker/Dockerfile .

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

    $ $GOPATH/bin/musala-push

### Using with docker

    $ docker build -f docker/Dockerfile -t musala-push .
    $ docker run  \
        -e EMAIL_FROM=
        -e EMAIL_TO=
        -e EMAIL_SUBJECT=
        -e EMAIL_FROM=
        -e EMAIL_TO=
        -e EMAIL_SUBJECT=
        -v /var/lib/musala-push/musala-push.cron:/app/musala-push.cron \
        -it musala-push:latest

## Contributing

1. Fork it ( http://github.com/glenux/musala-push/fork )
2. Create your feature branch (`git checkout -b my-new-feature`)
3. Commit your changes (`git commit -am 'Add some feature'`)
4. Push to the branch (`git push origin my-new-feature`)
5. Create new Pull Request


## Credits

* [Glenn Y. ROLLAND](https://github.com/glenux) - author & maintainer: 
* You? Fork the project and become a contributor!

Got questions? Need help? Tweet at [@glenux](https://twitter.com/glenux)


## License

Musala Push is Copyright © 2018-2019 Glenn ROLLAND. It is free software, and may be redistributed under the terms specified in the LICENSE.txt file.

