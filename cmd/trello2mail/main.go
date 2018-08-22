package main

import (
	"fmt"
	// "gopkg.in/russross/blackfriday.v2"
	// "github.com/davecgh/go-spew/spew"
)

func main() {
	// Setup config
	config := NewConfig()
	config.ParseEnv()

	// Get task list as markdown
	trelloCtx := NewTrello(config.Trello.Token)
	trelloBoard := trelloCtx.GetBoard(config.Trello.Url)
	trelloMarkdown := trelloBoard.ExportToMarkdown()
	trelloHtml := trelloBoard.ExportToHtml()
	config.Email.Subject = fmt.Sprintf("Daily mail for %s", trelloBoard.Name)

	// Create email enveloppe
	email := NewEmail()
	email.SetHeaders(config.Email)
	email.SetBody(trelloHtml, trelloMarkdown)

	// Connect and send email
	transport := NewTransport(config.Smtp)
	transport.Dial()
	transport.Authenticate()
	transport.Send(email)
	transport.Quit()
}
