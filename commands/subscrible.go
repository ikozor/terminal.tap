package commands

import (
	"context"
	"fmt"

	"github.com/terminaldotshop/terminal-sdk-go"
)

type sub struct {
	Subscription terminal.Subscription
	Address      terminal.Address
	Card         terminal.Card
}

func (c *CommandExecutor) ListSubscriptions() ([]terminal.Subscription, error) {
	res, err := c.client.Subscription.List(context.TODO())
	if err != nil {
		return nil, getApiErrorMessage(err)
	}
	return res.Data, nil
}

func (c *CommandExecutor) GetSubscription(id int) (sub, error) {
	subscriptions, err := c.ListSubscriptions()
	if err != nil {
		return sub{}, err
	}
	if len(subscriptions) < id+1 {
		return sub{}, fmt.Errorf("No subscription found with id: %d", id)
	}
	subscription := subscriptions[id]
	address, err := c.GetAddressByTerminalId(subscription.AddressID)
	if err != nil {
		return sub{}, err
	}

	card, err := c.GetCardByTerminalId(subscription.CardID)
	if err != nil {
		return sub{}, err
	}

	return sub{Subscription: subscription, Address: address, Card: card}, nil

}
