package main

import (
	"bytes"
	"fmt"
	// "github.com/adlio/trello"
	trello "github.com/glenux/contrib-trello"
	"github.com/russross/blackfriday/v2"
	"log"
	"net/url"
	"os"
	"os/exec"
	"path"
	"strings"
	"text/template"
)

const (
	// See https://trello.com/app-key
	APP_KEY string = "58117ebf843d49b05bca074c5fd520ee"
)

type TrelloConfig struct {
	Url   string
	Token string
}

type TrelloCtx struct {
	Token  string
	Client *trello.Client
}

type TrelloBoard struct {
	Ctx     *TrelloCtx
	Ptr     *trello.Board
	Starred bool
	Closed  bool
	Name    string
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

func (ctx *TrelloCtx) GetBoards() []TrelloBoard {
	var result []TrelloBoard

	token, err := ctx.Client.GetToken(ctx.Token, trello.Defaults())
	if err != nil {
		log.Panic(err)
	}

	member, err := ctx.Client.GetMember(token.IDMember, trello.Defaults())
	if err != nil {
		log.Panic(err)
	}

	boards, err := member.GetBoards(trello.Defaults())
	if err != nil {
		log.Panic(err)
	}
	for _, board := range boards {
		result = append(result, TrelloBoard{
			Ctx:     ctx,
			Starred: board.Starred,
			Closed:  board.Closed,
			Ptr:     board,
			Name:    board.Name,
		})
	}
	return result
}

func (ctx *TrelloCtx) GetBoard(boardUrl string) TrelloBoard {
	parsedUrl, err := url.Parse(boardUrl)
	if err != nil {
		panic(err)
	}
	boardId := strings.Split(parsedUrl.Path, "/")[2]

	board, err := ctx.Client.GetBoard(boardId, trello.Defaults())
	if err != nil {
		log.Panic(err)
	}
	return TrelloBoard{
		Ctx:     ctx,
		Starred: board.Starred,
		Closed:  board.Closed,
		Ptr:     board,
		Name:    board.Name,
	}
}

type CardData struct {
	Name string
	Url  string
}

type ListData struct {
	Name  string
	Cards []CardData
}

type BoardData struct {
	Name  string
	Url   string
	Lists []ListData
}

func (board *TrelloBoard) ExportData() BoardData {
	var boardData = BoardData{
		Name:  board.Ptr.Name,
		Url:   board.Ptr.ShortUrl,
		Lists: make([]ListData, 0),
	}

	// get lists
	lists, err := board.Ptr.GetLists(trello.Defaults())
	if err != nil {
		log.Panic(err)
	}

	// fill in reverse order
	for listIdx := len(lists) - 1; listIdx >= 0; listIdx -= 1 {
		list := lists[listIdx]
		listData := ListData{
			Name:  list.Name,
			Cards: make([]CardData, 0),
		}

		cards, err := list.GetCards(trello.Defaults())
		if err != nil {
			log.Panic(err)
		}
		for _, card := range cards {
			cardData := CardData{
				Name: card.Name,
				Url:  card.Url,
			}
			listData.Cards = append(listData.Cards, cardData)
		}

		boardData.Lists = append(boardData.Lists, listData)
	}

	return boardData
}

func (board *TrelloBoard) ExportToMarkdown() string {
	var buf bytes.Buffer

	data := board.ExportData()

	wd, err := os.Getwd()
	t, err := template.ParseFiles(path.Join(wd, "templates/markdown.tmpl"))
	if err != nil {
		log.Panic("Unable to parse template files")
	}

	if err := t.Execute(&buf, data); err != nil {
		log.Panicf("Template execution: %s", err)
	}

	return buf.String()
}

func (board *TrelloBoard) ExportToHtml() string {
	markdown := board.ExportToMarkdown()
	html := blackfriday.Run([]byte(markdown))
	return string(html)
}
