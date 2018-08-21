package main

import (
	"crypto/tls"
	// "errors"
	"fmt"
	"log"
	// "os"
	// "strconv"
	// "net"
	// "net/mail"
	"net/smtp"
)

type MailHeaders map[string]string

func (headers *MailHeaders) ParseConfig(config Config) (int, error) {
	(*headers)["From"] = config.EmailFrom
	(*headers)["To"] = config.EmailTo
	(*headers)["Subject"] = config.EmailSubject
	return 0, nil
}

func NewAuth(config Config) *smtp.Auth {

	switch config.SmtpAuthType {
	case "plain":
		auth := smtp.PlainAuth(
			"",
			config.SmtpUsername,
			config.SmtpPassword,
			config.SmtpHostname,
		)
		return &auth

	case "login":
		auth := LoginAuth(config.SmtpUsername, config.SmtpPassword)
		return &auth

	default:
	}
	return nil
}

func NewTLS(config Config) *tls.Config {
	// TLS config
	return &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         config.SmtpHostname,
	}
}

func NewSmtpClient(config Config) *smtp.Client {
	address := fmt.Sprintf("%s:%d", config.SmtpHostname, config.SmtpPort)
	tlsConfig := NewTLS(config)
	switch config.SmtpSecurityType {
	case "tls":
		fmt.Printf("Creating TLS connection to %s...\n", address)
		conn, err := tls.Dial("tcp", address, tlsConfig)
		if err != nil {
			log.Panic(err)
		}

		fmt.Println("Creating SMTP client...")
		c, err := smtp.NewClient(conn, config.SmtpHostname)
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
