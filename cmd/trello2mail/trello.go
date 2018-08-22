package main

import (
	//	"errors"
	"fmt"
	"github.com/adlio/trello"
	// "github.com/davecgh/go-spew/spew"
	//	"log"
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
	Ctx *TrelloCtx
	Ptr *trello.Board
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

func NewTrello(token string) *TrelloCtx {
	client := trello.NewClient(APP_KEY, token)
	/*
		if client == nil {
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

			log.Panic(errors.New(text))
		}
	*/
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

	// spew.Dump(boardId)
	board, err := ctx.Client.GetBoard(boardId, trello.Defaults())
	// spew.Dump(board)
	return TrelloBoard{Ctx: ctx, Ptr: board}
}

func (board *TrelloBoard) ExportToMarkdown() []string {
	// var s []string

	lists, _ := board.Ptr.GetLists(trello.Defaults())
	// spew.Dump(lists)
	// s = append(s, "# Trello board")
	for _, v := range lists {
		fmt.Println(v.Name)
	}
	return []string{}
}

/*
func RunTaskell(boardUrl string) {
	cmd := fmt.Sprintf("taskell -t %s -", boardUrl)
	markdown := strings.TrimSpace(runcmd(cmd))
}*/
