package main

// TODO: use https://github.com/domodwyer/mailyak as a base lib

import (
	"bytes"
	"encoding/base64"
	"errors"
	"fmt"
	// "github.com/davecgh/go-spew/spew"
	"log"
	"strings"
	// "strconv"
	"math/rand"
	"net/mail"
)

type EmailHeaders map[string]string

type EmailCtx struct {
	Headers   EmailHeaders
	BodyPlain string
	BodyHtml  string
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func (headers EmailHeaders) String() string {
	var buffer bytes.Buffer
	for k, v := range headers {
		buffer.WriteString(fmt.Sprintf("%s: %s\r\n", k, v))
	}
	return buffer.String()
}

func NewEmail() *EmailCtx {
	email := EmailCtx{}
	email.Headers = make(EmailHeaders)
	return &email
}

func encodeRFC2047(text string) string {
	// use mail's rfc2047 to encode any string
	addr := mail.Address{Name: text, Address: ""}
	return strings.Trim(addr.String(), " \"<@>")
}

func (email *EmailCtx) SetHeaders(config EmailConfig) {
	email.Headers["Return-Path"] = config.From
	email.Headers["From"] = config.From
	fmt.Printf("config.To %#v\n", config.To)
	if len(config.To) < 1 {
		errmsg := "EMAIL_TO must contain at least one value"
		log.Panic(errors.New(errmsg))
	}
	email.Headers["To"] = config.To[0]

	if len(config.To) > 1 {
		email.Headers["Cc"] = strings.Join(config.To[1:], ",")
	}
	email.Headers["Subject"] = encodeRFC2047(config.Subject)
	email.Headers["Content-Transfer-Encoding"] = "quoted-printable"
	email.Headers["MIME-Version"] = "1.0"

	return
}

func (email *EmailCtx) SetBody(html string, plain string) {
	email.BodyPlain = plain
	email.BodyHtml = html
	return
}

func (email *EmailCtx) String() string {
	var buffer bytes.Buffer
	mixBoundary := RandStringBytes(16)
	altBoundary := RandStringBytes(16)

	buffer.WriteString(email.Headers.String())
	buffer.WriteString(fmt.Sprintf("Content-Type: multipart/mixed;\r\n    boundary=\"%s\"\r\n", mixBoundary))
	buffer.WriteString("\r\n")

	buffer.WriteString(fmt.Sprintf("--%s\r\n", mixBoundary))
	buffer.WriteString(fmt.Sprintf("Content-Type: multipart/alternative;\r\n    boundary=\"%s\"\r\n", altBoundary))
	buffer.WriteString("\r\n")

	buffer.WriteString(fmt.Sprintf("--%s\r\n", altBoundary))
	buffer.WriteString(fmt.Sprintf("Content-Type: text/plain; charset=\"utf-8\"\r\n"))
	buffer.WriteString(fmt.Sprintf("Content-Transfer-Encoding: base64\r\n"))
	buffer.WriteString("\r\n")
	buffer.WriteString(base64.StdEncoding.EncodeToString([]byte(email.BodyPlain)))
	buffer.WriteString("\r\n")
	buffer.WriteString("\r\n")

	buffer.WriteString(fmt.Sprintf("--%s\r\n", altBoundary))
	buffer.WriteString(fmt.Sprintf("Content-Type: text/html; charset=\"utf-8\"\r\n"))
	buffer.WriteString(fmt.Sprintf("Content-Transfer-Encoding: base64\r\n"))
	buffer.WriteString("\r\n")
	buffer.WriteString(base64.StdEncoding.EncodeToString([]byte(email.BodyHtml)))
	buffer.WriteString("\r\n")
	buffer.WriteString("\r\n")

	buffer.WriteString(fmt.Sprintf("--%s--\r\n", altBoundary))
	buffer.WriteString("\r\n")

	buffer.WriteString(fmt.Sprintf("--%s--\r\n", mixBoundary))
	buffer.WriteString("\r\n")

	return buffer.String()
}
