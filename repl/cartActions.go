package repl

import (
	"fmt"

	"github.com/ikozor/terminal.tap/commands"
)

func (r *repl) getCart() {
	r.args = nil
	r.currentCommand = func(i interface{}) (string, error) {
		cart, err := r.commandExecutor.GetCart()
		if err != nil {
			return "", err
		}

		cartString := ""
		for _, e := range cart.Items {
			cartString += fmt.Sprintf("(Name: %s, Variant: %d, Price: %.2f USD, Quantity: %d), ",
				e.ProductName,
				e.VariantId,
				float32(e.Price)/100,
				e.Quantity,
			)
		}
		if len(cart.Items) > 1 {
			cartString = cartString[:len(cartString)-2] + " "
		}

		cartString += fmt.Sprintf("Total: %.2f USD", float32(cart.Total)/100)

		return cartString, nil
	}
	r.args = nil
}

func (r *repl) addtoCart(productName string, variantId, quantity int) {
	r.args = commands.CartItem{
		ProductName: productName,
		VariantId:   variantId,
		Quantity:    quantity,
	}

	r.currentCommand = func(i interface{}) (string, error) {
		item, ok := i.(commands.CartItem)
		if !ok {
			return "", fmt.Errorf("invalid Item to add to cart: %v", i)
		}

		if err := r.commandExecutor.ManageCart(item.ProductName, item.VariantId, item.Quantity); err != nil {
			return "", err
		}
		return "Successfully added to cart", nil
	}
}

func (r *repl) removeFromCart(productName string, variantId, quantity int) {
	r.args = commands.CartItem{
		ProductName: productName,
		VariantId:   variantId,
		Quantity:    quantity,
	}

	r.currentCommand = func(i interface{}) (string, error) {
		item, ok := i.(commands.CartItem)
		if !ok {
			return "", fmt.Errorf("invalid Item to add to cart: %v", i)
		}

		if err := r.commandExecutor.ManageCart(item.ProductName, item.VariantId, item.Quantity); err != nil {
			return "", err
		}
		return "Successfully removed item from cart", nil
	}
}

func (r *repl) convertToOrder() {
	r.args = nil
	r.currentCommand = func(i interface{}) (string, error) {
		tracking, err := r.commandExecutor.ConvertToOrder()
		if err != nil {
			return "", err
		}
		info := fmt.Sprintf("Tracking Info: Service: %s, Number: %s, url: %s",
			tracking.Service,
			tracking.Number,
			tracking.URL,
		)

		return info, nil

	}

}
