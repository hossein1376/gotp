package main

import (
	"fmt"
	"os"

	"github.com/hossein1376/gotp/cmd/gotp/command"
)

func main() {
	if err := command.Run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
}
