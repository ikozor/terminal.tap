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
	r.args = nil
	r.currentCommand = func(i interface{}) (string, error) { return "", nil }
	line := strings.Split(r.line, " ")
	if len(line) < 1 || line[0] == "" {
		// nothing happened
		return nil
	}

	command := line[0]
	switch command {
	case "PRODUCT":
		if len(line) < 2 {
			return fmt.Errorf("product action not specified")
		}
		switch line[1] {
		case "LIST":
			r.listProducts()
		case "GET":
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
		case "GET":
			r.getCart()
		case "ADD":
			if len(line) < 3 {
				return fmt.Errorf("product to add not specified")
			}
			if len(line) < 4 {
				return fmt.Errorf("variant to add not specified")
			}
			variant, err := strconv.Atoi(line[3])
			if err != nil {
				return fmt.Errorf("Variant id is not int: %s", line[3])
			}

			var quantity int
			if len(line) < 5 {
				quantity = 1
			} else {
				i, err := strconv.Atoi(line[4])
				if err != nil {
					return fmt.Errorf("cannot convert quantity to string: %s", line[4])
				}
				if i < 1 {
					return fmt.Errorf("cannot add by less than 1, got: %d", i)
				}
				quantity = i
			}
			r.addtoCart(line[2], variant, quantity)

		case "REMOVE":
			if len(line) < 3 {
				return fmt.Errorf("product to remove not specified")
			}
			if len(line) < 4 {
				return fmt.Errorf("variant to remove not specified")
			}
			variant, err := strconv.Atoi(line[3])
			if err != nil {
				return fmt.Errorf("Variant id is not int: %s", line[3])
			}

			var quantity int
			if len(line) < 5 {
				quantity = 0
			} else {
				i, err := strconv.Atoi(line[4])
				if err != nil {
					return fmt.Errorf("cannot convert quantity to string: %s", line[4])
				}
				if i < 1 {
					return fmt.Errorf("cannot remove by less than 1, got: %d", i)
				}
				quantity = i * -1
			}
			r.removeFromCart(line[2], variant, quantity)
		case "ORDER":
			r.convertToOrder()
		default:
			return fmt.Errorf("Cart action not found: %s", line[1])

		}

	case "ADDRESS":
		if len(line) < 2 {
			return fmt.Errorf("No address action specified")
		}
		switch line[1] {
		case "LIST":
			r.listAddresses()
		case "ADD":
			if len(line) < 3 {
				return fmt.Errorf("No address specified to add")
			}
			addressString := strings.Join(line[2:], " ")
			address, err := validAddress(addressString)
			if err != nil {
				return err
			}
			r.AddAddressAction(address)
		case "REMOVE":
			if len(line) < 3 {
				return fmt.Errorf("Address to remove not specified")
			}
			r.removeAddress(line[2])

		case "SET":
			if len(line) < 3 {
				return fmt.Errorf("No address to set specified")
			}
			r.setAddress(line[2])

		default:
			return fmt.Errorf("Address action not found: %s", line[1])
		}
	case "CARD":
		if len(line) < 2 {
			return fmt.Errorf("No card action specified")
		}
		switch line[1] {
		case "LIST":
			r.listCards()
		case "ADD":
			r.addCard()
		case "REMOVE":
			if len(line) < 3 {
				return fmt.Errorf("No card specified to remove")
			}
			r.removeCard(line[2])
		case "SET":
			if len(line) < 3 {
				return fmt.Errorf("No card specified to set")
			}
			r.setCard(line[2])
		default:
			return fmt.Errorf("Card action not found: %s", line[1])
		}
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
