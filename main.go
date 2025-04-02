package main

import (
	"fmt"

	"github.com/ikozor/terminal.tap/repl"
)

func main() {

	r := repl.NewRepl()
	fmt.Print("\033[H\033[2J")
	for {
		if err := r.Read(); err != nil {
			fmt.Println(err)
			continue
		}

		if err := r.Evaluate(); err != nil {
			fmt.Println(err)
			continue
		}
		res, err := r.Process()
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println(res)

	}
}
