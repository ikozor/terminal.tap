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
			return fmt.Errorf("what to get not passed")
		}

		switch line[1] {
		case "PRODUCT":
			if len(line) < 3 {
				return fmt.Errorf("Product to get not passed")
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
			r.args = commands.CartItem{
				ProductName: line[2],
				Quantity:    quantity,
			}

			r.currentCommand = func(i interface{}) (string, error) {
				item, ok := i.(commands.CartItem)
				if !ok {
					return "", fmt.Errorf("invalid Item to add to cart: %v", i)
				}

				if err := r.commandExecutor.ManageCart(item.ProductName, item.Quantity); err != nil {
					return "", err
				}
				return "Successfully added to cart", nil
			}
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
			r.args = commands.CartItem{
				ProductName: line[2],
				Quantity:    quantity,
			}

			r.currentCommand = func(i interface{}) (string, error) {
				item, ok := i.(commands.CartItem)
				if !ok {
					return "", fmt.Errorf("invalid Item to add to cart: %v", i)
				}

				if err := r.commandExecutor.ManageCart(item.ProductName, item.Quantity); err != nil {
					return "", err
				}
				return "Successfully removed item from cart", nil
			}

		case "GET":
			r.currentCommand = func(i interface{}) (string, error) {
				cart, err := r.commandExecutor.GetCart()
				if err != nil {
					return "", err
				}

				cartString := ""
				for _, e := range cart.Items {
					cartString += fmt.Sprintf("(Name: %s, Price: %d USD, Quantity: %d), ",
						e.ProductName,
						e.Price,
						e.Quantity,
					)
				}
				cartString = cartString[:len(cartString)-2]

				cartString += fmt.Sprintf(" Total: %d USD", cart.Total)

				return cartString, nil
			}
			r.args = nil

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
