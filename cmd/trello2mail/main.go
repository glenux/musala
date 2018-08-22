package main

import (
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

	// Create email enveloppe
	email := NewEmail()
	email.MakeHeaders(config.Email)
	email.MakeBody(trelloMarkdown)

	// Connect and send email
	transport := NewTransport(config.Smtp)
	transport.Dial()
	transport.Authenticate()
	transport.Send(email)
	transport.Quit()
}
