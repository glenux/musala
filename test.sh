#!/bin/sh 

export EMAIL_FROM="admin@example.com"
export EMAIL_TO="admin@example.com"
export EMAIL_SUBJECT="Daily trello mail"

export TRELLO_URL="https://trello.com/b/someId/someName"
export TRELLO_TOKEN="xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"

export SMTP_USERNAME="admin"
export SMTP_PASSWORD="admin"
export SMTP_PORT="465"
export SMTP_HOSTNAME="mail.example.com"
export SMTP_USE_AUTH="true"
export SMTP_USE_SSL="true"

./_build/trello2mail
