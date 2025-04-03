package commands

import (
	"context"

	"github.com/terminaldotshop/terminal-sdk-go"
)

func (c *CommandExecutor) AddAddress(address terminal.AddressNewParams) error {
	_, err := c.client.Address.New(context.TODO(), address)
	if err != nil {
		return getApiErrorMessage(err)
	}
	return nil
}
