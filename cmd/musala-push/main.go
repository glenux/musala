package main

import (
	"crypto/tls"
	"fmt"
	"github.com/go-mail/mail"
	"os"
)

func main() {
	// Setup config
	fmt.Println("d: parsing config")
	config := NewConfig()
	config.Parse()
	fmt.Printf("%+v\n", config)

	// Get task list as markdown
	fmt.Println("d: configuring trello")
	trelloCtx := NewTrello(config.TrelloAPIKey, config.TrelloToken)
	if trelloCtx == nil {
		fmt.Println("ERROR: Unable to initialize trello context")
		os.Exit(1)
	}

	fmt.Println("d: getting trello boards")
	var trelloBoardsList []TrelloBoard
	if len(config.TrelloURL) > 0 {
		fmt.Printf("d: using given url %s\n", config.TrelloURL)
		trelloBoard := trelloCtx.GetBoard(config.TrelloURL)
		trelloBoardsList = append(trelloBoardsList, trelloBoard)
	} else {
		fmt.Println("d: fetching boards")
		trelloBoardsList = trelloCtx.GetBoards()
	}

	for _, trelloBoard := range trelloBoardsList {
		fmt.Printf("d: loading board %s\n", trelloBoard.Name)
		if !trelloBoard.Starred || trelloBoard.Closed {
			fmt.Printf("d: skipping %s\n", trelloBoard.Name)
			continue
		}
		fmt.Printf("d: exporting content of %s\n", trelloBoard.Name)

		trelloMarkdown := trelloBoard.ExportToMarkdown()
		trelloHTML := trelloBoard.ExportToHtml()
		config.EmailSubject = fmt.Sprintf("Daily mail for %s", trelloBoard.Name)

		// Create email enveloppe
		email := mail.NewMessage()
		email.SetHeader("To", config.EmailTo[0])
		if len(config.EmailTo) > 0 {
			email.SetHeader("Cc", config.EmailTo[1:]...)
		}
		email.SetHeader("From", config.EmailFrom)
		email.SetHeader("Subject", config.EmailSubject)
		email.SetBody("text/plain", trelloMarkdown)
		email.AddAlternative("text/html", trelloHTML)

		// Connect and send email
		var transport *mail.Dialer
		if len(config.SMTPUsername) > 0 {
			fmt.Println("d: transport w/ username")
			transport = mail.NewDialer(
				config.SMTPHostname,
				int(config.SMTPPort),
				config.SMTPUsername,
				config.SMTPPassword,
			)
			// disable cert verification
			transport.TLSConfig = &tls.Config{InsecureSkipVerify: true}
		} else {
			fmt.Println("d: transport w/out username")
			transport = &mail.Dialer{
				Host: config.SMTPHostname,
				Port: int(config.SMTPPort),
			}
		}

		if err := transport.DialAndSend(email); err != nil {
			panic(err)
		}
	}
}
