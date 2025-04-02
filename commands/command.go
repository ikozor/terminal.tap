package commands

import (
	"github.com/terminaldotshop/terminal-sdk-go"
	"github.com/terminaldotshop/terminal-sdk-go/option"
)

/*
list product

get product
	* name

get cart

add to cart
	* product name
	* quantity

set address
	* address

set card
	* card

order

order
	loop:
		* product name
		* quantity
	* address
	* cardinfo

*/

type CommandExecutor struct {
	client          *terminal.Client
	currentProducts []terminal.Product
}

func CreateCommandExecutor() *CommandExecutor {

	client := terminal.NewClient(
		option.WithBaseURL("https://api.dev.terminal.shop"), // the Double Slash was causing panic
	)
	return &CommandExecutor{client: client}

}
