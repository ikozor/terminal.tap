package commands

import (
	"context"
	"fmt"

	"github.com/terminaldotshop/terminal-sdk-go"
)

func (c *CommandExecutor) ListOrders() ([]terminal.Order, error) {
	res, err := c.client.Order.List(context.TODO())
	if err != nil {
		return nil, getApiErrorMessage(err)
	}
	return res.Data, nil
}

func (c *CommandExecutor) GetOrder(id int) (terminal.Order, error) {
	orders, err := c.ListOrders()
	if err != nil {
		return terminal.Order{}, err
	}

	for _, e := range orders {
		if e.Index == int64(id) {
			return e, nil
		}
	}
	return terminal.Order{}, fmt.Errorf("Order not found: %d", id)

}
