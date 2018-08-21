package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/smtp"
)

type TransportCtx struct {
	Config  SmtpConfig
	Address string
	Auth    *smtp.Auth
	Tls     *tls.Config
	Client  *smtp.Client
}

func NewTransport(config SmtpConfig) *TransportCtx {
	ctx := TransportCtx{}
	ctx.Config = config
	ctx.Address = fmt.Sprintf("%s:%d", config.Hostname, config.Port)
	ctx.Auth = NewTransportAuth(config)
	ctx.Tls = NewTransportTls(config)
	return &ctx
}

func NewTransportAuth(config SmtpConfig) *smtp.Auth {
	switch config.AuthType {
	case "plain":
		auth := smtp.PlainAuth(
			"",
			config.Username,
			config.Password,
			config.Hostname,
		)
		return &auth

	case "login":
		auth := LoginAuth(config.Username, config.Password)
		return &auth

	default:
	}
	return nil
}

func NewTransportTls(config SmtpConfig) *tls.Config {
	// TLS config
	return &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         config.Hostname,
	}
}

func (ctx *TransportCtx) DialInsecure() {
	// no SSL/TLS
	fmt.Println("Creating SMTP client...")
	c, err := smtp.Dial(ctx.Address)
	if err != nil {
		log.Panic(err)
	}
	ctx.Client = c
}

func (ctx *TransportCtx) DialTls() {
	fmt.Printf("Creating TLS connection to %s...\n", ctx.Address)
	conn, err := tls.Dial("tcp", ctx.Address, ctx.Tls)
	if err != nil {
		log.Panic(err)
	}

	fmt.Println("Creating SMTP client...")
	c, err := smtp.NewClient(conn, ctx.Config.Hostname)
	if err != nil {
		log.Panic(err)
	}
	ctx.Client = c
}

func (ctx *TransportCtx) DialStartTls() {
	fmt.Println("Creating SMTP client...")
	c, err := smtp.Dial(ctx.Address)
	if err != nil {
		log.Panic(err)
	}
	fmt.Printf("Creating StartTLS connection to %s...\n", ctx.Address)
	c.StartTLS(ctx.Tls)

	ctx.Client = c
}

func (ctx *TransportCtx) Dial() {
	switch ctx.Config.SecurityType {
	case "tls":
		ctx.DialTls()

	case "starttls":
		ctx.DialStartTls()

	default:
		ctx.DialInsecure()
	}
}

func (ctx *TransportCtx) Authenticate() {
	err := ctx.Client.Auth(*ctx.Auth)
	if err != nil {
		log.Panic(err)
	}
}

func (ctx *TransportCtx) Quit() {
	ctx.Client.Quit()
}

func (ctx *TransportCtx) Send(email *EmailCtx) {
	return
}
