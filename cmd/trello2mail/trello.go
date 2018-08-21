package main

import (
	//	"errors"
	//	"fmt"
	"github.com/adlio/trello"
	//	"github.com/davecgh/go-spew/spew"
	//	"log"
	"os/exec"
	//	"strings"
)

const (
	// FIXME: declare app to trello and get a real token for this app
	APP_KEY string = "80dbcf6f88f62cc5639774e13342c20b"
)

type TrelloCtx struct {
	Token  string
	Client *trello.Client
}

type TrelloItem struct {
	Title string
}

type TrelloList struct {
	Items []TrelloItem
}

type TrelloBoard struct {
	Ctx   TrelloCtx
	Lists []TrelloList
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
		spew.Dump(client)
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
	return TrelloBoard{Ctx: *ctx}
}

func (*TrelloBoard) ExportToMarkdown() []string {
	return []string{}
}

/*
func RunTaskell(boardUrl string) {
	cmd := fmt.Sprintf("taskell -t %s -", boardUrl)
	markdown := strings.TrimSpace(runcmd(cmd))
}*/
