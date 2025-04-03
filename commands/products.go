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
		return nil, getApiErrorMessage(err)
	}
	c.currentProducts = products.Data
	productNames := []string{}
	for _, product := range products.Data {
		productNames = append(productNames, product.Name)
	}
	return productNames, nil
}

func (c *CommandExecutor) GetProductInfo(productName string) (terminal.Product, error) {
	productId, err := c.GetProductId(productName)
	if err != nil {
		return terminal.Product{}, err
	}

	productGetResponse, err := c.client.Product.Get(context.TODO(), productId)
	if err != nil {
		return terminal.Product{}, getApiErrorMessage(err)
	}
	product := productGetResponse.Data
	return product, nil
}

func (c *CommandExecutor) GetProductId(name string) (string, error) {
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

func (c *CommandExecutor) GetProductByVariantId(variantId string) (terminal.Product, error) {
	if c.currentProducts == nil {
		if err := c.populateCurrentProducts(); err != nil {
			return terminal.Product{}, err
		}
	}
	if len(c.currentProducts) < 1 {
		if err := c.populateCurrentProducts(); err != nil {
			return terminal.Product{}, err
		}
	}
	for _, product := range c.currentProducts {
		for _, variant := range product.Variants {
			if variant.ID == variantId {
				return product, nil
			}
		}
	}
	return terminal.Product{}, fmt.Errorf("Product Variant Id not found : %s", variantId)

}

func (c *CommandExecutor) populateCurrentProducts() error {
	products, err := c.client.Product.List(context.TODO())
	if err != nil {
		return getApiErrorMessage(err)
	}
	c.currentProducts = products.Data
	return nil
}
