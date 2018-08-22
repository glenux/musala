package main

import (
	"strings"
	// "errors"
	"bytes"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	// "log"
	// "os"
	// "strconv"
	// "net"
	"net/mail"
	// "net/smtp"
)

type EmailHeaders map[string]string
type EmailBody string

type EmailCtx struct {
	Headers EmailHeaders
	Body    EmailBody
}

func (headers EmailHeaders) String() string {
	var buffer bytes.Buffer
	for k, v := range headers {
		buffer.WriteString(fmt.Sprintf("%s: %s\r\n", k, v))
	}
	return buffer.String()
}

func (body EmailBody) String() string {
	res := string(body)
	if false {
		spew.Dump(res)
	}
	return res
}

func NewEmail() *EmailCtx {
	email := EmailCtx{}
	email.Headers = make(EmailHeaders)
	return &email
}

func encodeRFC2047(text string) string {
	// use mail's rfc2047 to encode any string
	addr := mail.Address{text, ""}
	return strings.Trim(addr.String(), " \"<@>")
}

func (email *EmailCtx) MakeHeaders(config EmailConfig) {
	email.Headers["Return-Path"] = config.From
	email.Headers["From"] = config.From
	email.Headers["To"] = config.To
	email.Headers["Subject"] = encodeRFC2047(config.Subject)
	// email.Headers["Content-Type"] = "text/plain; charset=\"us-ascii\";"
	email.Headers["Content-Type"] = "text/plain; charset=\"utf-8\";"
	email.Headers["Content-Transfer-Encoding"] = "base64"
	email.Headers["MIME-Version"] = "1.0"

	return
}

func (email *EmailCtx) MakeBody(content string) {
	email.Body = EmailBody(content)
	if false {
		spew.Dump(email.Body)
	}
	return
}
