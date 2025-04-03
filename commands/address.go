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

func (c *CommandExecutor) ListAddresses() ([]terminal.Address, error) {
	res, err := c.client.Address.List(context.TODO())
	if err != nil {
		return nil, getApiErrorMessage(err)
	}
	return res.Data, nil
}
