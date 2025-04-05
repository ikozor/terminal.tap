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

type NewSubscription struct {
	ProductName      string
	ProductVariantId int
	Quantity         int
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

func (c *CommandExecutor) AddSubscription(newSub NewSubscription) error {
	if c.currentAddress == "" {
		return fmt.Errorf("No Address Set")
	}
	if c.currentCard == "" {
		return fmt.Errorf("no card set")
	}

	product, err := c.FindProductByName(newSub.ProductName)
	if err != nil {
		return err
	}
	if len(product.Variants) < newSub.ProductVariantId+1 {
		return fmt.Errorf("No Variant found with id: %d", newSub.ProductVariantId)
	}

	_, err = c.client.Subscription.New(context.TODO(), terminal.SubscriptionNewParams{
		Subscription: terminal.SubscriptionParam{
			ProductVariantID: terminal.String(product.Variants[newSub.ProductVariantId].ID),
			Quantity:         terminal.Int(int64(newSub.Quantity)),
			AddressID:        terminal.String(c.currentAddress),
			CardID:           terminal.String(c.currentCard),
		},
	})
	if err != nil {
		return getApiErrorMessage(err)
	}

	return nil
}

func (c *CommandExecutor) RemoveSubscription(id int) error {
	subs, err := c.ListSubscriptions()
	if err != nil {
		return err
	}
	if len(subs) < id+1 {
		return fmt.Errorf("Subscription id not found: %d", id)
	}
	_, err = c.client.Subscription.Delete(context.TODO(), subs[id].ID)
	if err != nil {
		return getApiErrorMessage(err)
	}
	return nil

}
