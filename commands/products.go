package commands

import (
	"context"
	"fmt"
	"strings"

	"github.com/terminaldotshop/terminal-sdk-go"
)

func (c *CommandExecutor) ListProductNames() ([]string, error) {
	products, err := c.client.Product.List(context.TODO())
	if err != nil {
		return nil, err
	}
	c.currentProducts = products.Data
	productNames := []string{}
	for _, product := range products.Data {
		productNames = append(productNames, product.Name)
	}
	return productNames, nil
}

func (c *CommandExecutor) GetProductInfo(productName string) (terminal.Product, error) {
	productId, err := c.getProductId(productName)
	if err != nil {
		return terminal.Product{}, err
	}

	productGetResponse, err := c.client.Product.Get(context.TODO(), productId)
	if err != nil {
		return terminal.Product{}, err
	}
	product := productGetResponse.Data
	return product, nil
}

func (c *CommandExecutor) getProductId(name string) (string, error) {
	if c.currentProducts == nil {
		if err := c.populateCurrentProducts(); err != nil {
			return "", err
		}
	}
	if len(c.currentProducts) < 1 {
		if err := c.populateCurrentProducts(); err != nil {
			return "", err
		}
	}
	for _, e := range c.currentProducts {
		if strings.ToUpper(e.Name) == name {
			return e.ID, nil
		}
	}
	return "", fmt.Errorf("Product Id not found for: %s", name)

}

func (c *CommandExecutor) populateCurrentProducts() error {
	products, err := c.client.Product.List(context.TODO())
	if err != nil {
		return err
	}
	c.currentProducts = products.Data
	return nil
}
