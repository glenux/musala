package main

import (
	"errors"
	"fmt"
	// "log"
	"os"
	"strconv"
	// "net"
	// "net/mail"
	// "gopkg.in/russross/blackfriday.v2"
)

type ConfigEntry struct {
	Type   string
	Ptr    interface{}
	Values []string
}

type Config struct {
	EmailFrom        string
	EmailTo          string
	EmailSubject     string
	SmtpHostname     string
	SmtpPort         uint16
	SmtpUsername     string
	SmtpPassword     string
	SmtpAuthType     string
	SmtpSecurityType string
	TrelloUrl        string
	TrelloToken      string
}

func NewConfig() *Config {
	return &Config{}
}

func (config *Config) ParseEnv() (int, error) {
	// map env variables to config pointers
	dataMap := map[string](ConfigEntry){
		"EMAIL_FROM":    ConfigEntry{"string", &(config.EmailFrom), nil},
		"EMAIL_TO":      ConfigEntry{"string", &(config.EmailTo), nil},
		"EMAIL_SUBJECT": ConfigEntry{"string", &(config.EmailSubject), nil},
		"TRELLO_URL":    ConfigEntry{"string", &(config.TrelloUrl), nil},
		"TRELLO_TOKEN":  ConfigEntry{"string", &(config.TrelloToken), nil},

		"SMTP_HOSTNAME": ConfigEntry{"string", &(config.SmtpHostname), nil},
		"SMTP_USERNAME": ConfigEntry{"string", &(config.SmtpUsername), nil},
		"SMTP_PASSWORD": ConfigEntry{"string", &(config.SmtpPassword), nil},
		"SMTP_PORT":     ConfigEntry{"uint16", &(config.SmtpPort), nil},

		"SMTP_AUTH_TYPE":     ConfigEntry{"string", &(config.SmtpAuthType), []string{"none", "plain", "login"}},
		"SMTP_SECURITY_TYPE": ConfigEntry{"string", &(config.SmtpSecurityType), []string{"none", "tls", "starttls"}},
	}

	for envVar, mapEntry := range dataMap {
		envValue := os.Getenv(envVar)
		if len(envValue) == 0 {
			return -1, errors.New(fmt.Sprintf(
				"Empty environment variable. Please set %s value", envVar))
		}

		if mapEntry.Values != nil {
			allowedValue := false
			for _, v := range mapEntry.Values {
				if v == envValue {
					allowedValue = true
				}
			}
			if !allowedValue {
				return -1, errors.New(fmt.Sprintf(
					"Wrong value for %s=%s. Value must be one of %v", envVar, envValue, mapEntry.Values))
			}
		}

		switch mapEntry.Type {
		case "string":
			*(mapEntry.Ptr.(*string)) = envValue

		case "uint16":
			u64, err := strconv.ParseUint(envValue, 10, 16)
			if err != nil {
				return -1, errors.New(fmt.Sprintf(
					"Unable to convert %s=%s to unsigned int", envVar, envValue))
			}
			*(mapEntry.Ptr.(*uint16)) = uint16(u64)

		case "bool":
			b, err := strconv.ParseBool(envValue)
			if err != nil {
				return -1, errors.New(fmt.Sprintf(
					"Unable to convert %s=%s to boolean", envVar, envValue))
			}
			*(mapEntry.Ptr.(*bool)) = b

		default:
			return -1, errors.New(fmt.Sprintf("Undefined parser for %s<%s>", envVar, mapEntry.Type))
		}
	}
	fmt.Printf("%#v\n", config)
	return 0, nil
}
