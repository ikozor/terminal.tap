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
			morse, morseErr := morsecode.ReadStringIntoMorse(err.Error())
			if morseErr != nil {
				panic(morseErr)
			}
			fmt.Println(morse)

			text, err := morsecode.ReadMorseIntoString(morse)
			fmt.Println(text)
			fmt.Println(err)
			continue
		}

		if err := r.Evaluate(); err != nil {
			morse, morseErr := morsecode.ReadStringIntoMorse(err.Error())
			if morseErr != nil {
				panic(morseErr)
			}
			fmt.Println(morse)

			text, err := morsecode.ReadMorseIntoString(morse)
			fmt.Println(text)
			fmt.Println(err)
			continue
		}
		res, err := r.Process()
		if err != nil {
			morse, morseErr := morsecode.ReadStringIntoMorse(err.Error())
			if morseErr != nil {
				panic(morseErr)
			}
			fmt.Println(morse)

			text, err := morsecode.ReadMorseIntoString(morse)
			fmt.Println(text)
			fmt.Println(err)
			continue
		}
		morse, morseErr := morsecode.ReadStringIntoMorse(res)
		if morseErr != nil {
			panic(morseErr)
		}
		fmt.Println(morse)
		text, err := morsecode.ReadMorseIntoString(morse)
		fmt.Println(text)
		fmt.Println(err)

	}
}
