package main

import (
	"bytes"
	"fmt"
	"github.com/adlio/trello"
	"gopkg.in/russross/blackfriday.v2"
	"log"
	"net/url"
	"os/exec"
	"strings"
)

const (
	// FIXME: declare app to trello and get a real token for this app
	APP_KEY string = "80dbcf6f88f62cc5639774e13342c20b"
)

type TrelloCtx struct {
	Token  string
	Client *trello.Client
}

type TrelloBoard struct {
	Ctx  *TrelloCtx
	Ptr  *trello.Board
	Name string
}

func runcmd(command string) string {
	shell := "/bin/sh"
	flag := "-c"

	out, err := exec.Command(shell, flag, command).Output()
	if err != nil {
		panic(err)
	}

	return string(out)
}

func GetTokenProcessMessage() string {
	url := strings.Join([]string{
		"https://trello.com/1/authorize?expiration=never",
		"name=taskell",
		"scope=read",
		"response_type=token",
		fmt.Sprintf("key=%s", APP_KEY),
	}, "&")

	text := strings.Join([]string{
		"Wrong TRELLO_TOKEN value. Please visit:",
		url,
		"When you have your access token, set TRELLO_TOKEN=<your-token>",
	}, "\n\n")

	return text
}

func NewTrello(token string) *TrelloCtx {
	client := trello.NewClient(APP_KEY, token)

	ctx := TrelloCtx{}
	ctx.Token = token
	ctx.Client = client

	return &ctx
}

func (ctx *TrelloCtx) GetBoard(boardUrl string) TrelloBoard {
	parsedUrl, err := url.Parse(boardUrl)
	if err != nil {
		panic(err)
	}
	boardId := strings.Split(parsedUrl.Path, "/")[2]

	board, err := ctx.Client.GetBoard(boardId, trello.Defaults())
	return TrelloBoard{Ctx: ctx, Ptr: board, Name: board.Name}
}

func (board *TrelloBoard) ExportToMarkdown() string {
	var markdown bytes.Buffer
	var text string

	lists, err := board.Ptr.GetLists(trello.Defaults())
	if err != nil {
		log.Panic(err)
	}

	text = fmt.Sprintf("# Board %s\n\n", board.Ptr.Name)
	markdown.WriteString(text)

	text = fmt.Sprintf("URL: %s\n", board.Ptr.ShortUrl)
	markdown.WriteString(text)

	for listIdx := len(lists) - 1; listIdx >= 0; listIdx -= 1 {
		list := lists[listIdx]
		text := fmt.Sprintf("\n## %s\n\n", list.Name)
		markdown.WriteString(text)

		cards, err := list.GetCards(trello.Defaults())
		if err != nil {
			log.Panic(err)
		}
		for _, card := range cards {
			text := fmt.Sprintf("* %s\n", card.Name)
			markdown.WriteString(text)
		}
	}
	return markdown.String()
}

func (board *TrelloBoard) ExportToHtml() string {
	markdown := board.ExportToMarkdown()
	html := blackfriday.Run([]byte(markdown))
	return string(html)
}

/*
func RunTaskell(boardUrl string) {
	cmd := fmt.Sprintf("taskell -t %s -", boardUrl)
	markdown := strings.TrimSpace(runcmd(cmd))
}*/
