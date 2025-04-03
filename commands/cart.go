package commands

import (
	"context"
	"fmt"
	"strings"

	"github.com/terminaldotshop/terminal-sdk-go"
)

type CartItem struct {
	ProductName string
	Quantity    int
	Price       int
}

type Cart struct {
	Items   []CartItem
	Total   int
	Address *terminal.Address
	Card    *terminal.Card
}

func (c *CommandExecutor) ManageCart(productName string, quantityDiff int) error {
	product, err := c.GetProductInfo(productName)
	if err != nil {
		return err
	}
	curCart, err := c.GetCart()
	if err != nil {
		return err
	}
	curQuantity := 0
	for _, e := range curCart.Items {
		if strings.ToUpper(e.ProductName) == productName {
			curQuantity = e.Quantity
		}
	}
	if len(product.Variants) < 1 {
		return fmt.Errorf("Product has no variants: %s", productName)
	}
	body := terminal.CartSetItemParams{
		ProductVariantID: terminal.String(product.Variants[0].ID),
		Quantity:         terminal.Int(int64(curQuantity + quantityDiff)),
	}

	if curQuantity+quantityDiff < 0 {
		body.Quantity = terminal.Int(0)
	}
	if quantityDiff == 0 {
		body.Quantity = terminal.Int(0)
	}
	_, err = c.client.Cart.SetItem(context.TODO(), body)
	if err != nil {
		return getApiErrorMessage(err)
	}
	return nil
}

func (c *CommandExecutor) GetCart() (Cart, error) {
	res, err := c.client.Cart.Get(context.TODO())
	if err != nil {
		return Cart{}, getApiErrorMessage(err)
	}
	cart := Cart{}
	for _, e := range res.Data.Items {
		product, err := c.GetProductByVariantId(e.ProductVariantID)
		if err != nil {
			return Cart{}, err
		}
		cart.Items = append(cart.Items, CartItem{ProductName: product.Name, Quantity: int(e.Quantity), Price: int(e.Subtotal)})
	}
	cart.Total = int(res.Data.Subtotal)
	return cart, nil
}
