package commands

import (
	"context"
	"fmt"

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

func (c *CommandExecutor) RemoveCard(last4 string) error {
	cards, err := c.ListCards()
	if err != nil {
		return err
	}
	for _, e := range cards {
		if e.Last4 == last4 {
			_, err := c.client.Card.Delete(context.TODO(), e.ID)
			if err != nil {
				return getApiErrorMessage(err)
			}
			if c.currentCard == e.ID {
				c.currentCard = ""
			}
			return nil
		}
	}

	return fmt.Errorf("Card with last4 %s not found", last4)
}

func (c *CommandExecutor) SetCard(last4 string) error {
	cards, err := c.ListCards()
	if err != nil {
		return err
	}
	for _, e := range cards {
		if e.Last4 == last4 {
			c.currentCard = e.ID
			return nil
		}
	}
	return fmt.Errorf("Card with last4 %s not found", last4)

}

func (c *CommandExecutor) GetCardByTerminalId(cardId string) (terminal.Card, error) {
	res, err := c.client.Card.Get(context.TODO(), cardId)
	if err != nil {
		return terminal.Card{}, getApiErrorMessage(err)
	}
	return res.Data, nil

}
