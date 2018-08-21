package main

// Examples
// - Sending emails with SSL : https://gist.github.com/chrisgillis/10888032
// - Project layout https://github.com/golang-standards/project-layout
// - Markdown rendering https://github.com/russross/blackfriday

import (
	"fmt"
	"log"
	"os"
	// "net"
	// "net/mail"
	// "gopkg.in/russross/blackfriday.v2"
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
	if _, err := config.ParseEnv(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err.Error())
		os.Exit(1)
	}

	// Get task list as markdown
	trelloCtx := NewTrello(config.Trello.Token)
	trelloBoard := trelloCtx.GetBoard(config.Trello.Url)
	trelloMarkdown := trelloBoard.ExportToMarkdown()

	// Create and send email
	email := NewEmail()
	email.MakeHeaders(config.Email)
	email.MakeBody(trelloMarkdown)
	email.Send()

	transport := NewTransport(config.Smtp) 
	if err := transport.Authenticate() {
		log.Panic(err)
	}
	transport.Send(email)

	// Connect & authenticate
	fmt.Println("Connecting...")
	client := NewSmtpClient(*config)

	// Build auth
	authConfig := NewAuth(*config)
	fmt.Printf("Authenticating...\n")

	if err := client.Auth(*authConfig); err != nil {
	fmt.Println("Disconnecting...")
	client.Quit()

	// Write email
}
