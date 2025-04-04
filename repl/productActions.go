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
			productsString += e.Name + ", "
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
		variants := "("
		for i, e := range res.Variants {
			variants += fmt.Sprintf("Id: %d, Name: %s, Price: %.2f USD), ", i, e.Name, float32(e.Price)/100)
		}

		product := fmt.Sprintf("Name: %s, Variants: [%s], Description: %s",
			res.Name,
			variants[:len(variants)-2],
			res.Description,
		)
		return product, nil
	}

}
