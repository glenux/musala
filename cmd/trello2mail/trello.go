// create taskell configuration
// run taskell and read content ?
// use https://github.com/adlio/trello ?

package main

import (
	"fmt"
	//"github.com/VojtechVitek/go-trello"
	"os/exec"
	"strings"
)

type TrelloCtx struct {
	Token string
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
	return &TrelloCtx{Token: token}
}

func (ctx *TrelloCtx) GetBoard(boardUrl string) TrelloBoard {
	return TrelloBoard{Ctx: *ctx}
}

func (*TrelloBoard) ExportToMarkdown() []string {
}

func RunTaskell(boardUrl string) {
	cmd := fmt.Sprintf("taskell -t %s -", boardUrl)
	markdown := strings.TrimSpace(runcmd(cmd))
}
