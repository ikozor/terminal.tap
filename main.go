package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	morsecode "github.com/ikozor/terminal.tap/morse-code"
	"github.com/ikozor/terminal.tap/repl"
	"gopkg.in/yaml.v3"
)

func main() {
	input := "stdin"
	output := "stdout"
	config := getInputOutput()

	if config.Input != "" {
		input = config.Input
	}
	if config.Output != "" {
		output = config.Output
	}

	confirm := bufio.NewScanner(os.Stdin)

	fmt.Printf("Using:\n\tInput: %s\n\tOutput: %s\nokay? [y/n]\n\n", input, output)
	for {
		fmt.Print("> ")
		confirm.Scan()
		confirmation := strings.ToLower(confirm.Text())
		if confirmation == "y" || confirmation == "yes" {
			break
		}
		if confirmation == "n" || confirmation == "no" {
			return
		}
		fmt.Printf("Unknown input: %s\n", confirmation)
	}

	r := repl.NewRepl(input)
	for {
		if err := r.Read(); err != nil {
			morse, morseErr := morsecode.ReadStringIntoMorse(err.Error())
			if morseErr != nil {
				// print invalid morse code character found
				fmt.Println(".. -. ...- .- .-.. .. -.. / -- --- .-. ... . / -.-. --- -.. . / -.-. .... .- .-. .- -.-. - . .-.")
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
				// print invalid morse code character found
				fmt.Println(".. -. ...- .- .-.. .. -.. / -- --- .-. ... . / -.-. --- -.. . / -.-. .... .- .-. .- -.-. - . .-.")
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
				// print invalid morse code character found
				fmt.Println(".. -. ...- .- .-.. .. -.. / -- --- .-. ... . / -.-. --- -.. . / -.-. .... .- .-. .- -.-. - . .-.")
			}
			fmt.Println(morse)

			text, err := morsecode.ReadMorseIntoString(morse)
			fmt.Println(text)
			fmt.Println(err)
			continue
		}
		morse, morseErr := morsecode.ReadStringIntoMorse(res)
		if morseErr != nil {
			// print invalid morse code character found
			fmt.Println(".. -. ...- .- .-.. .. -.. / -- --- .-. ... . / -.-. --- -.. . / -.-. .... .- .-. .- -.-. - . .-.")
		}
		fmt.Println(morse)
		text, err := morsecode.ReadMorseIntoString(morse)
		fmt.Println(text)
		fmt.Println(err)

	}
}

type inputOutput struct {
	Input  string `yaml:"input"`
	Output string `yaml:"output"`
}

func getInputOutput() inputOutput {
	file, err := os.ReadFile("config.yaml")
	config := inputOutput{}
	if err != nil {
		return inputOutput{}
	}

	err = yaml.Unmarshal(file, &config)
	if err != nil {
		return inputOutput{}
	}

	return config
}
