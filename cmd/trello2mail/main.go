package main

import (
	"crypto/tls"
	"fmt"
	"github.com/go-mail/mail"
)

func main() {
	// Setup config
	fmt.Println("d: parsing config")
	config := NewConfig()
	config.Parse()

	// Get task list as markdown
	fmt.Println("d: configuring trello")
	trelloCtx := NewTrello(config.TrelloToken)

	fmt.Println("d: getting trello boards")
	var trelloBoardsList []TrelloBoard
	if len(config.TrelloUrl) > 0 {
		fmt.Printf("d: using given url %s\n", config.TrelloUrl)
		trelloBoard := trelloCtx.GetBoard(config.TrelloUrl)
		trelloBoardsList = append(trelloBoardsList, trelloBoard)
	} else {
		fmt.Println("d: fetching boards")
		trelloBoardsList = trelloCtx.GetBoards()
	}

	for _, trelloBoard := range trelloBoardsList {
		fmt.Printf("d: loading board %s\n", trelloBoard.Name)
		if !trelloBoard.Starred || trelloBoard.Closed {
			fmt.Println("d: skipping")
			continue
		}
		fmt.Println("d: exporting content")

		trelloMarkdown := trelloBoard.ExportToMarkdown()
		trelloHtml := trelloBoard.ExportToHtml()
		config.EmailSubject = fmt.Sprintf("Daily mail for %s", trelloBoard.Name)

		// Create email enveloppe
		email := mail.NewMessage()
		email.SetHeader("To", config.EmailTo[0])
		if len(config.EmailTo) > 0 {
			email.SetHeader("Cc", config.EmailTo[1:]...)
		}
		email.SetHeader("From", config.EmailFrom)
		email.SetHeader("Subject", config.EmailSubject)
		email.SetBody(trelloHtml, trelloMarkdown)

		// Connect and send email
		var transport *mail.Dialer
		if len(config.SmtpUsername) > 0 {
			fmt.Println("transport w/ username")
			transport = mail.NewDialer(
				config.SmtpHostname,
				int(config.SmtpPort),
				config.SmtpUsername,
				config.SmtpPassword,
			)
			// disable cert verification
			transport.TLSConfig = &tls.Config{InsecureSkipVerify: true}
		} else {
			fmt.Println("transport w/ no username")
			transport = &mail.Dialer{
				Host: config.SmtpHostname,
				Port: int(config.SmtpPort),
			}
		}

		if err := transport.DialAndSend(email); err != nil {
			panic(err)
		}
	}
}
