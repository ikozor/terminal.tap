package repl

import "fmt"

func (r *repl) listProducts() {
	r.currentCommand = func(i interface{}) (string, error) {
		products, err := r.commandExecutor.ListProductNames()
		if err != nil {
			return "", err
		}
		if len(products) < 1 {
			return "", fmt.Errorf("No products found")
		}

		productsString := ""
		for _, e := range products {
			productsString += e + ", "
		}
		return productsString[:len(productsString)-2], nil

	}
	r.args = nil
}

func (r *repl) getProduct() {
	r.currentCommand = func(i interface{}) (string, error) {
		s, ok := i.(string)
		if !ok {
			return "", fmt.Errorf("invalid product: %v", i)
		}
		res, err := r.commandExecutor.GetProductInfo(s)
		if err != nil {
			return "", err
		}
		if len(res.Variants) < 1 {
			return "", fmt.Errorf("No Variants for product")
		}
		product := fmt.Sprintf("Name: %s, Type: %s, Price: %d, Description: %s",
			res.Name,
			res.Variants[0].Name,
			res.Variants[0].Price,
			res.Description,
		)
		return product, nil
	}

}
