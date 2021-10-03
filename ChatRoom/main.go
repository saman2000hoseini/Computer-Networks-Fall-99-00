package main

import (
	"os"

	"github.com/saman2000hoseini/Computer-Networks-Fall-99-00/ChatRoom/cmd"
)

const (
	exitFailure = 1
)

func main() {
	root := cmd.NewRootCommand()

	if root != nil {
		if err := root.Execute(); err != nil {
			os.Exit(exitFailure)
		}
	}
}
