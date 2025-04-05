package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	morsecode "github.com/ikozor/terminal.tap/morse-code"
	"github.com/ikozor/terminal.tap/repl"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
	apiKey, ok := os.LookupEnv("TERMINAL_TOKEN")
	if !ok {
		panic("No api key provided")
	}
	apiUrl, ok := os.LookupEnv("TERMINAL_URL")
	if !ok {
		panic("No api url provided")
	}

	var input, output string
	input, ok = os.LookupEnv("INPUT")
	if !ok {
		input = "stdin"
	}
	output, ok = os.LookupEnv("OUTPUT")
	if !ok {
		output = "stdout"
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

	r := repl.NewRepl(apiUrl, apiKey, input)
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
