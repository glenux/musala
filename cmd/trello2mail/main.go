package main

import (
	"fmt"
)

func main() {
	// Setup config
	config := NewConfig()
	config.Parse()

	// Get task list as markdown
	trelloCtx := NewTrello(config.TrelloToken)
	trelloBoard := trelloCtx.GetBoard(config.TrelloUrl)
	trelloMarkdown := trelloBoard.ExportToMarkdown()
	trelloHtml := trelloBoard.ExportToHtml()
	config.EmailSubject = fmt.Sprintf("Daily mail for %s", trelloBoard.Name)

	// Create email enveloppe
	email := NewEmail()
	email.SetHeaders(EmailConfig{
		From:    config.EmailFrom,
		To:      config.EmailTo,
		Subject: config.EmailSubject,
	})
	email.SetBody(trelloHtml, trelloMarkdown)

	// Connect and send email
	transport := NewTransport(SmtpConfig{
		Hostname:     config.SmtpHostname,
		Port:         config.SmtpPort,
		Username:     config.SmtpUsername,
		Password:     config.SmtpPassword,
		AuthType:     config.SmtpAuthType,
		SecurityType: config.SmtpSecurityType,
	})

	transport.Dial()
	transport.Authenticate()
	transport.Send(email)
	transport.Quit()
}
