package main

import (
	"bufio"
	"fmt"
	"log"
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

	log.SetFlags(0)
	if output == "stdout" {
		log.SetOutput(os.Stdout)
	} else {
		file, err := os.OpenFile(output, os.O_RDWR, 0666)
		if err != nil {
			panic(err)
		}
		defer file.Close()
		log.SetOutput(file)
	}

	r := repl.NewRepl(input)
	for {
		if err := r.Read(); err != nil {
			morse, morseErr := morsecode.ReadStringIntoMorse(err.Error())
			if morseErr != nil {
				// print invalid morse code character found
				log.Println(".. -. ...- .- .-.. .. -.. / -- --- .-. ... . / -.-. --- -.. . / -.-. .... .- .-. .- -.-. - . .-.")
			}
			log.Println(morse)
			continue
		}

		if err := r.Evaluate(); err != nil {
			morse, morseErr := morsecode.ReadStringIntoMorse(err.Error())
			if morseErr != nil {
				// print invalid morse code character found
				log.Println(".. -. ...- .- .-.. .. -.. / -- --- .-. ... . / -.-. --- -.. . / -.-. .... .- .-. .- -.-. - . .-.")
			}
			log.Println(morse)
			continue
		}

		res, err := r.Process()
		if err != nil {
			morse, morseErr := morsecode.ReadStringIntoMorse(err.Error())
			if morseErr != nil {
				// print invalid morse code character found
				log.Println(".. -. ...- .- .-.. .. -.. / -- --- .-. ... . / -.-. --- -.. . / -.-. .... .- .-. .- -.-. - . .-.")
			}
			log.Println(morse)
			continue
		}
		morse, morseErr := morsecode.ReadStringIntoMorse(res)
		if morseErr != nil {
			// print invalid morse code character found
			log.Println(".. -. ...- .- .-.. .. -.. / -- --- .-. ... . / -.-. --- -.. . / -.-. .... .- .-. .- -.-. - . .-.")
		}
		log.Println(morse)
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
