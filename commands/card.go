package commands

import (
	"context"

	"github.com/terminaldotshop/terminal-sdk-go"
)

func (c *CommandExecutor) ListCards() ([]terminal.Card, error) {
	res, err := c.client.Card.List(context.TODO())
	if err != nil {
		return nil, getApiErrorMessage(err)
	}
	return res.Data, nil
}

func (c *CommandExecutor) AddCard() (string, error) {
	res, err := c.client.Card.Collect(context.TODO())
	if err != nil {
		return "", getApiErrorMessage(err)
	}
	return res.Data.URL, nil
}
