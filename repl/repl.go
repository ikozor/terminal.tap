package repl

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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
		return err
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
		r.listProducts()

	case "GET":
		if len(line) < 2 {
			return fmt.Errorf("what to get not passed")
		}
		switch line[1] {
		case "CART":
			r.getCart()

		case "PRODUCT":
			if len(line) < 3 {
				return fmt.Errorf("Product to get not passed")
			}
			productName := ""
			for _, e := range line[2:] {
				productName += e + " "
			}
			r.args = productName[:len(productName)-1]
			r.getProduct()

		default:
			return fmt.Errorf("Cannot Get: %s", line[1])

		}
	case "CART":
		if len(line) < 2 {
			return fmt.Errorf("No cart action provided")
		}
		switch line[1] {
		case "ADD":
			if len(line) < 3 {
				return fmt.Errorf("Nothing to add to cart")
			}

			var quantity int
			if len(line) < 4 {
				quantity = 1
			} else {
				i, err := strconv.Atoi(line[3])
				if err != nil {
					return fmt.Errorf("cannot convert quantity to string: %s", line[3])
				}
				if i < 1 {
					return fmt.Errorf("cannot add by less than 1, got: %d", i)
				}
				quantity = i
			}
			r.addtoCart(line[2], quantity)

		case "REMOVE":
			if len(line) < 3 {
				return fmt.Errorf("Nothing to add to cart")
			}

			var quantity int
			if len(line) < 4 {
				quantity = 0
			} else {
				i, err := strconv.Atoi(line[3])
				if err != nil {
					return fmt.Errorf("cannot convert quantity to string: %s", line[3])
				}
				if i < 1 {
					return fmt.Errorf("cannot remove by less than 1, got: %d", i)
				}
				quantity = i * -1
			}
			r.removeFromCart(line[2], quantity)

		case "SET":
			if len(line) < 3 {
				return fmt.Errorf("What to set unspecified")
			}
			switch line[2] {
			case "ADDRESS":

			case "CARD":
			default:
				return fmt.Errorf("Cannot card set: %s", line[2])
			}
		default:
			return fmt.Errorf("No cart action: %s", line[1])
		}

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
