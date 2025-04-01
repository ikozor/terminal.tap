package commands

import (
	"context"
)

func (c *commandExecutor) ListProductNames() ([]string, error) {
	products, err := c.client.Product.List(context.TODO())
	if err != nil {
		return nil, err
	}
	productNames := []string{}
	for _, product := range products.Data {
		productNames = append(productNames, product.Name)
	}
	return productNames, nil
}
