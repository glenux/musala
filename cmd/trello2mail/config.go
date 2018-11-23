package main

import (
	// "errors"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"strings"
)

const (
	programBinary string = "trello2mail"
)

var (
	ALLOWED_AUTH_TYPES     = []string{"none", "plain", "login"}
	ALLOWED_SECURITY_TYPES = []string{"none", "tls", "starttls"}
)

type Config struct {
	EmailFrom    string   `mapstructure:"email-from"`
	EmailTo      []string `mapstructure:"email-to"`
	EmailSubject string   `mapstructure:"email-subject"`

	SmtpHostname     string `mapstructure:"smtp-hostname"`
	SmtpPort         uint16 `mapstructure:"smtp-port"`
	SmtpUsername     string `mapstructure:"smtp-username"`
	SmtpPassword     string `mapstructure:"smtp-password"`
	SmtpAuthType     string `mapstructure:"smtp-auth-type"`
	SmtpSecurityType string `mapstructure:"smtp-security-type"`

	TrelloUrl   string `mapstructure:"trello-url"`
	TrelloToken string `mapstructure:"trello-token"`

	Parser *cobra.Command `mapstructure:"-"`
}

func NewConfig() *Config {
	self := &Config{}

	cmd := &cobra.Command{
		Use: programBinary,
		Run: func(cmd *cobra.Command, args []string) { /* placeholder */ },
	}

	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
	viper.AutomaticEnv()

	cmd.PersistentFlags().StringVarP(&self.EmailFrom, "email-from", "", "", "address of sender")
	cmd.PersistentFlags().StringArrayVarP(&self.EmailTo, "email-to", "", []string{}, "address(es) of recipient(s)")
	cmd.PersistentFlags().StringVarP(&self.EmailSubject, "email-subject", "", "", "email subject")
	viper.BindPFlag("email-from", cmd.PersistentFlags().Lookup("email-from"))
	viper.BindPFlag("email-to", cmd.PersistentFlags().Lookup("email-to"))
	viper.BindPFlag("email-subject", cmd.PersistentFlags().Lookup("email-subject"))

	cmd.PersistentFlags().StringVarP(&self.TrelloUrl, "trello-url", "", "", "url of trello board")
	cmd.PersistentFlags().StringVarP(&self.TrelloToken, "trello-token", "", "", "url of trello token")
	viper.BindPFlag("trello-url", cmd.PersistentFlags().Lookup("trello-url"))
	viper.BindPFlag("trello-token", cmd.PersistentFlags().Lookup("trello-token"))

	cmd.PersistentFlags().StringVarP(&self.SmtpHostname, "smtp-hostname", "", "", "address of smtp server")
	cmd.PersistentFlags().StringVarP(&self.SmtpUsername, "smtp-username", "", "", "username for smtp server")
	cmd.PersistentFlags().StringVarP(&self.SmtpPassword, "smtp-password", "", "", "password for smtp server")
	cmd.PersistentFlags().Uint16VarP(&self.SmtpPort, "smtp-port", "", 25, "port for smtp server")
	cmd.PersistentFlags().StringVarP(&self.SmtpAuthType, "smtp-auth-type", "", "", "authentication type for smtp server")
	cmd.PersistentFlags().StringVarP(&self.SmtpSecurityType, "smtp-security-type", "", "", "security type for smtp server")
	viper.BindPFlag("smtp-hostname", cmd.PersistentFlags().Lookup("smtp-hostname"))
	viper.BindPFlag("smtp-username", cmd.PersistentFlags().Lookup("smtp-username"))
	viper.BindPFlag("smtp-password", cmd.PersistentFlags().Lookup("smtp-password"))
	viper.BindPFlag("smtp-port", cmd.PersistentFlags().Lookup("smtp-port"))
	viper.BindPFlag("smtp-auth-type", cmd.PersistentFlags().Lookup("smtp-auth-type"))
	viper.BindPFlag("smtp-security-type", cmd.PersistentFlags().Lookup("smtp-security-type"))

	self.Parser = cmd
	return self
}

func (self *Config) Parse() error {
	// set config defaults
	// persistent flags
	// environment & config
	// viper.SetEnvPrefix("")

	if err := self.Parser.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err := viper.Unmarshal(&self); err != nil {
		panic("Unable to unmarshal config")
	}

	// spew.Dump(config)
	return nil
}
