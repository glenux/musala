package main

import (
	"fmt"
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
	return ctx
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

func (*TransportCtx) DialInsecure() error {
}

func (*TransportCtx) DialTls() error {
}

func (*TransportCtx) DialStartTls(address) error {
}

func (*TransportCtx) Dial() *smtp.Client {
	switch config.SecurityType {
	case "tls":
		fmt.Printf("Creating TLS connection to %s...\n", address)
		conn, err := tls.Dial("tcp", address, tlsConfig)
		if err != nil {
			log.Panic(err)
		}

		fmt.Println("Creating SMTP client...")
		c, err := smtp.NewClient(conn, config.Hostname)
		if err != nil {
			log.Panic(err)
		}
		return c

	case "starttls":
		fmt.Println("Creating SMTP client...")
		c, err := smtp.Dial(address)
		if err != nil {
			log.Panic(err)
		}
		fmt.Printf("Creating StartTLS connection to %s...\n", address)
		c.StartTLS(tlsConfig)

		return c

	default:
		// no SSL/TLS
		fmt.Println("Creating SMTP client...")
		c, err := smtp.Dial(address)
		if err != nil {
			log.Panic(err)
		}
		return c
	}
}
