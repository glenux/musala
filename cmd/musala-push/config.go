package main

import (
	// "errors"
	"fmt"
	// "github.com/davecgh/go-spew/spew"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"strings"
)

const (
	programBinary string = "musala-push"
)

var (
	ALLOWED_AUTH_TYPES     = []string{"none", "plain", "login"}
	ALLOWED_SECURITY_TYPES = []string{"none", "tls", "starttls"}
)

// Config
type Config struct {
	EmailFrom    string   `mapstructure:"email-from"`
	EmailTo      []string `mapstructure:"email-to"`
	EmailSubject string   `mapstructure:"email-subject"`

	SMTPHostname     string `mapstructure:"smtp-hostname"`
	SMTPPort         uint16 `mapstructure:"smtp-port"`
	SMTPUsername     string `mapstructure:"smtp-username"`
	SMTPPassword     string `mapstructure:"smtp-password"`
	SMTPAuthType     string `mapstructure:"smtp-auth-type"`
	SMTPSecurityType string `mapstructure:"smtp-security-type"`

	TrelloURL    string `mapstructure:"trello-url"`
	TrelloAPIKey string `mapstructure:"trello-api-key"`
	TrelloToken  string `mapstructure:"trello-token"`

	Parser *cobra.Command `mapstructure:"-"`
}

// NewConfig : create configuration object
func NewConfig() *Config {
	config := &Config{}

	cmd := &cobra.Command{
		Use: programBinary,
		Run: func(cmd *cobra.Command, args []string) { /* placeholder */ },
	}

	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
	viper.AutomaticEnv()

	cmd.PersistentFlags().StringVarP(&config.EmailFrom, "email-from", "", "", "address of sender")
	cmd.PersistentFlags().StringArrayVarP(&config.EmailTo, "email-to", "", []string{}, "address(es) of recipient(s)")
	cmd.PersistentFlags().StringVarP(&config.EmailSubject, "email-subject", "", "", "email subject")

	cmd.PersistentFlags().StringVarP(&config.TrelloURL, "trello-url", "", "", "url of trello board")
	cmd.PersistentFlags().StringVarP(&config.TrelloURL, "trello-api-key", "", "", "API KEY for trello access")
	cmd.PersistentFlags().StringVarP(&config.TrelloToken, "trello-token", "", "", "TOKEN for trello access")

	cmd.PersistentFlags().StringVarP(&config.SMTPHostname, "smtp-hostname", "", "", "address of smtp server")
	cmd.PersistentFlags().StringVarP(&config.SMTPUsername, "smtp-username", "", "", "username for smtp server")
	cmd.PersistentFlags().StringVarP(&config.SMTPPassword, "smtp-password", "", "", "password for smtp server")
	cmd.PersistentFlags().Uint16VarP(&config.SMTPPort, "smtp-port", "", 25, "port for smtp server")
	cmd.PersistentFlags().StringVarP(&config.SMTPAuthType, "smtp-auth-type", "", "", "authentication type for smtp server")
	cmd.PersistentFlags().StringVarP(&config.SMTPSecurityType, "smtp-security-type", "", "", "security type for smtp server")

	viper.BindPFlag("email-from", cmd.PersistentFlags().Lookup("email-from"))
	viper.BindPFlag("email-to", cmd.PersistentFlags().Lookup("email-to"))
	viper.BindPFlag("email-subject", cmd.PersistentFlags().Lookup("email-subject"))
	viper.BindPFlag("trello-url", cmd.PersistentFlags().Lookup("trello-url"))
	viper.BindPFlag("trello-token", cmd.PersistentFlags().Lookup("trello-token"))
	viper.BindPFlag("trello-api-key", cmd.PersistentFlags().Lookup("trello-api-key"))
	viper.BindPFlag("smtp-hostname", cmd.PersistentFlags().Lookup("smtp-hostname"))
	viper.BindPFlag("smtp-username", cmd.PersistentFlags().Lookup("smtp-username"))
	viper.BindPFlag("smtp-password", cmd.PersistentFlags().Lookup("smtp-password"))
	viper.BindPFlag("smtp-port", cmd.PersistentFlags().Lookup("smtp-port"))
	viper.BindPFlag("smtp-auth-type", cmd.PersistentFlags().Lookup("smtp-auth-type"))
	viper.BindPFlag("smtp-security-type", cmd.PersistentFlags().Lookup("smtp-security-type"))

	config.Parser = cmd
	return config
}

// Parse : handle command line options
func (config *Config) Parse() error {
	// set config defaults
	// persistent flags
	// environment & config
	// viper.SetEnvPrefix("")

	if err := config.Parser.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	showHelp, _ := config.Parser.Flags().GetBool("help")
	if showHelp {
		os.Exit(0)
	}

	if err := viper.Unmarshal(&config); err != nil {
		panic("Unable to unmarshal config")
	}

	// spew.Dump(config)
	return nil
}
