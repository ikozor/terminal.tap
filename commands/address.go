package commands

import (
	"context"
	"fmt"
	"strings"

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

func (c *CommandExecutor) RemoveAddress(name string) error {
	addresses, err := c.ListAddresses()
	if err != nil {
		return err
	}

	for _, e := range addresses {
		if strings.ToUpper(e.Name) == name {
			_, err := c.client.Address.Delete(context.TODO(), e.ID)
			if err != nil {
				return getApiErrorMessage(err)
			}
			if c.currentAddress == e.ID {
				c.currentAddress = ""
			}
		}
	}
	return nil
}

func (c *CommandExecutor) SetAddress(name string) error {
	addresses, err := c.ListAddresses()
	if err != nil {
		return err
	}
	for _, e := range addresses {
		if strings.ToUpper(e.Name) == name {
			c.currentAddress = e.ID
			return nil
		}
	}
	return fmt.Errorf("Address not found with name: %s", name)
}
