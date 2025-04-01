package main

import (
	"fmt"

	morsecode "github.com/ikozor/terminal.tap/morse-code"
	"github.com/ikozor/terminal.tap/repl"
)

func main() {

	r := repl.NewRepl()
	fmt.Print("\033[H\033[2J")
	for {
		if err := r.Read(); err != nil {
			fmt.Println(morsecode.ReadStringIntoMorse(err.Error()))
			continue
		}

	}
}
