package commands

import (
	"context"

	"github.com/terminaldotshop/terminal-sdk-go"
)

func (c *CommandExecutor) ListSubscriptions() ([]terminal.Subscription, error) {
	res, err := c.client.Subscription.List(context.TODO())
	if err != nil {
		return nil, getApiErrorMessage(err)
	}
	return res.Data, nil
}
