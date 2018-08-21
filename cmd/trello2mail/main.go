package main

// Examples
// - Sending emails with SSL : https://gist.github.com/chrisgillis/10888032
// - Project layout https://github.com/golang-standards/project-layout
// - Markdown rendering https://github.com/russross/blackfriday

import (
// "gopkg.in/russross/blackfriday.v2"
// "github.com/davecgh/go-spew/spew"
)

func BuildContent(config Config) []string {
	// run taskell (download tasks from trello and export markdown)
	// read file as an array
	// insert trello board url
	// convert to HTML

	// output := blackfriday.Run(input, blackfriday.WithNoExtensions())
	return []string{}
}

func ImportFromTrello() {
}

func main() {
	// Setup config
	config := NewConfig()
	config.ParseEnv()

	// Get task list as markdown
	trelloCtx := NewTrello(config.Trello.Token)
	trelloBoard := trelloCtx.GetBoard(config.Trello.Url)
	trelloMarkdown := trelloBoard.ExportToMarkdown()
	panic("samere")

	// Create email enveloppe
	email := NewEmail()
	email.MakeHeaders(config.Email)
	email.MakeBody(trelloMarkdown)
	email.Send()

	// Connect and send email
	transport := NewTransport(config.Smtp)
	transport.Dial()
	transport.Authenticate()
	transport.Send(email)
	transport.Quit()
}
