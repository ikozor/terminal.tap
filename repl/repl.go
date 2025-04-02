package repl

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/ikozor/terminal.tap/commands"
	morsecode "github.com/ikozor/terminal.tap/morse-code"
)

type repl struct {
	line            string
	currentCommand  func(interface{}) (string, error)
	args            interface{}
	commandExecutor *commands.CommandExecutor
	scanner         *bufio.Scanner
}

func NewRepl() *repl {
	return &repl{
		scanner:         bufio.NewScanner(os.Stdin),
		commandExecutor: commands.CreateCommandExecutor(),
	}
}

func (r *repl) Read() error {
	fmt.Print("> ")
	r.scanner.Scan()
	text, err := morsecode.ReadMorseIntoString(r.scanner.Text())
	if err != nil {
		fmt.Println("READ ERROR: ", err)
		errMessageInMorseCode, morseErr := morsecode.ReadStringIntoMorse(err.Error())
		if morseErr != nil {
			return morseErr
		}
		return errors.New(errMessageInMorseCode)
	}
	r.line = text
	return nil
}

func (r *repl) Evaluate() error {
	line := strings.Split(r.line, " ")
	if len(line) < 1 || line[0] == "" {
		// nothing happened
		return nil
	}

	command := line[0]
	switch command {
	case "LIST":
		if len(line) < 2 {
			return fmt.Errorf("What to list unknown")
		}
		if line[1] != "PRODUCTS" {
			return fmt.Errorf("Cannot list: %s", line[1])
		}
		r.currentCommand = func(i interface{}) (string, error) {
			products, err := r.commandExecutor.ListProductNames()
			if err != nil {
				return "", err
			}
			if len(products) < 1 {
				return "", fmt.Errorf("No products found")
			}

			productsString := ""
			for _, e := range products {
				productsString += e + ", "
			}
			return productsString[:len(productsString)-2], nil

		}
		r.args = nil
	case "GET":
		if len(line) < 2 {
			return fmt.Errorf("What to get unknown")
		}

		switch line[1] {
		case "PRODUCT":
			if len(line) < 3 {
				return fmt.Errorf("What product info to get unkown")
			}
			productName := ""
			for _, e := range line[2:] {
				productName += e + " "
			}
			r.args = productName[:len(productName)-1]

			r.currentCommand = func(i interface{}) (string, error) {
				s, ok := i.(string)
				if !ok {
					return "", fmt.Errorf("invalid product: %v", i)
				}
				res, err := r.commandExecutor.GetProductInfo(s)
				if err != nil {
					return "", err
				}
				if len(res.Variants) < 1 {
					return "", fmt.Errorf("No Variants for product")
				}
				product := fmt.Sprintf("Name: %s, Type: %s, Price: %d, Description: %s",
					res.Name,
					res.Variants[0].Name,
					res.Variants[0].Price,
					res.Description,
				)
				return product, nil
			}
		}
	case "ADD":
	case "SET":
	case "ORDER":
	default:
		return fmt.Errorf("Command not found: %s", command)

	}

	return nil
}

func (r *repl) Process() (string, error) {
	if r.currentCommand == nil {
		return "", fmt.Errorf("No current command to execute")
	}
	return r.currentCommand(r.args)
}
