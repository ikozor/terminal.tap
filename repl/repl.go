package repl

import (
	"bufio"
	"errors"
	"fmt"
	"os"

	morsecode "github.com/ikozor/terminal.tap/morse-code"
)

type repl struct {
	line           string
	currentCommand string
	args           string
	scanner        *bufio.Scanner
}

func NewRepl() *repl {
	return &repl{scanner: bufio.NewScanner(os.Stdin)}
}

func (r *repl) Read() error {
	fmt.Print("> ")
	r.scanner.Scan()
	text, err := morsecode.ReadStringIntoMorse(r.scanner.Text())

	if err != nil {
		errMessageInMorseCode, morseErr := morsecode.ReadStringIntoMorse(err.Error())
		if morseErr != nil {
			return morseErr
		}
		return errors.New(errMessageInMorseCode)
	}
	r.line = text
	fmt.Println(r.line)
	return nil
}

func (r *repl) Evaluate() error {
	return nil
}

func (r *repl) Execute(cmd string) (string, error) {
	switch cmd {
	case "LIST PRODUCT":
		return "Hello", nil
	default:
		return "", fmt.Errorf("Command %s not found", cmd)

	}
}
