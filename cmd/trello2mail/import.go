// create taskell configuration
// run taskell and read content ?
// use https://github.com/adlio/trello ?

package main

import (
	"fmt"
	"os/exec"
	"strings"
)

func runcmd(command string) string {
	shell := "/bin/sh"
	flag := "-c"

	out, err := exec.Command(shell, flag, command).Output()
	if err != nil {
		panic(err)
	}

	return string(out)
}

func CreateTaskell() {
}

func RunTaskell(boardUrl string) {
	cmd := fmt.Sprintf("taskell -t %s -", boardUrl)
	markdown := strings.TrimSpace(runcmd(cmd))
}
