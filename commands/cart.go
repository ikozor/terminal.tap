package commands

import (
	"context"
	"fmt"
	"strings"

	"github.com/terminaldotshop/terminal-sdk-go"
)

type CartItem struct {
	ProductName string
	VariantId   int
	Quantity    int
	Price       int
}

type Cart struct {
	Items   []CartItem
	Total   int
	Address *terminal.Address
	Card    *terminal.Card
}

func (c *CommandExecutor) ManageCart(productName string, variantId, quantityDiff int) error {
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
	if len(product.Variants) < variantId+1 {
		return fmt.Errorf("Product %s has no variants with variantId: %d", productName, variantId)
	}
	body := terminal.CartSetItemParams{
		ProductVariantID: terminal.String(product.Variants[variantId].ID),
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
		variantId := -1
		for i, v := range product.Variants {
			if v.ID == e.ProductVariantID {
				variantId = i
			}
		}
		if variantId == -1 {
			return Cart{}, fmt.Errorf("Variant not found")
		}
		cart.Items = append(cart.Items, CartItem{
			ProductName: product.Name,
			VariantId:   variantId,
			Quantity:    int(e.Quantity),
			Price:       int(e.Subtotal)},
		)
	}
	cart.Total = int(res.Data.Subtotal)
	return cart, nil
}

func (c *CommandExecutor) ConvertToOrder() (terminal.OrderTracking, error) {
	if c.currentAddress == "" {
		return terminal.OrderTracking{}, fmt.Errorf("Address is not set, use ADDRESS SET")
	}
	if c.currentCard == "" {
		return terminal.OrderTracking{}, fmt.Errorf("Card is not set, use CARD SET")
	}
	_, err := c.client.Cart.SetCard(context.TODO(),
		terminal.CartSetCardParams{CardID: terminal.String(c.currentCard)})
	if err != nil {
		return terminal.OrderTracking{}, getApiErrorMessage(err)
	}
	_, err = c.client.Cart.SetAddress(context.TODO(),
		terminal.CartSetAddressParams{AddressID: terminal.String(c.currentAddress)})
	if err != nil {
		return terminal.OrderTracking{}, getApiErrorMessage(err)
	}

	res, err := c.client.Cart.Convert(context.TODO())
	if err != nil {
		return terminal.OrderTracking{}, getApiErrorMessage(err)
	}
	return res.Data.Tracking, nil
}
