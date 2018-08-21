package main

import (
// "errors"
// "fmt"
// "log"
// "os"
// "strconv"
// "net"
// "net/mail"
// "net/smtp"
)

type MailHeaders map[string]string
type MailBody []string

type EmailCtx struct {
	Headers MailHeaders
	Body    MailBody
}

func NewEmail() *EmailCtx {
	return &EmailCtx{}
}

func (email *EmailCtx) MakeHeaders(config EmailConfig) (int, error) {
	email.Headers["From"] = config.From
	email.Headers["To"] = config.To
	email.Headers["Subject"] = config.Subject
	return 0, nil
}

func (email *EmailCtx) MakeBody(content []string) (int, error) {
	email.Body = content
	return 0, nil
}

func (email *EmailCtx) Send() {
}
