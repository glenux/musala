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

type TrelloConfig struct {
	Url   string
	Token string
}

type SmtpConfig struct {
	Hostname     string
	Port         uint16
	Username     string
	Password     string
	AuthType     string
	SecurityType string
}

type EmailConfig struct {
	From    string
	To      string
	Subject string
}
type Config struct {
	Email  EmailConfig
	Smtp   SmtpConfig
	Trello TrelloConfig
}

func NewConfig() *Config {
	return &Config{}
}

func (config *Config) ParseEnv() (int, error) {
	// map env variables to config pointers
	dataMap := map[string](ConfigEntry){
		"EMAIL_FROM":    ConfigEntry{"string", &(config.Email.From), nil},
		"EMAIL_TO":      ConfigEntry{"string", &(config.Email.To), nil},
		"EMAIL_SUBJECT": ConfigEntry{"string", &(config.Email.Subject), nil},
		"TRELLO_URL":    ConfigEntry{"string", &(config.Trello.Url), nil},
		"TRELLO_TOKEN":  ConfigEntry{"string", &(config.Trello.Token), nil},

		"SMTP_HOSTNAME": ConfigEntry{"string", &(config.Smtp.Hostname), nil},
		"SMTP_USERNAME": ConfigEntry{"string", &(config.Smtp.Username), nil},
		"SMTP_PASSWORD": ConfigEntry{"string", &(config.Smtp.Password), nil},
		"SMTP_PORT":     ConfigEntry{"uint16", &(config.Smtp.Port), nil},

		"SMTP_AUTH_TYPE":     ConfigEntry{"string", &(config.Smtp.AuthType), []string{"none", "plain", "login"}},
		"SMTP_SECURITY_TYPE": ConfigEntry{"string", &(config.Smtp.SecurityType), []string{"none", "tls", "starttls"}},
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
