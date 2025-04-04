package repl

import (
	"fmt"
)

func (r *repl) listOrders() {
	r.args = nil
	r.currentCommand = func(i interface{}) (string, error) {
		orders, err := r.commandExecutor.ListOrders()
		if err != nil {
			return "", err
		}

		if len(orders) < 1 {
			return "No orders", nil
		}

		str := "("
		for _, e := range orders {
			str += fmt.Sprintf("Id: %d, Amount: %d), ", e.Index, e.Amount.Subtotal)
		}
		return str[:len(str)-2], nil
	}
}

func (r *repl) getOrder(id int) {
	r.args = id
	r.currentCommand = func(i interface{}) (string, error) {
		order, err := r.commandExecutor.GetOrder(id)
		if err != nil {
			return "", err
		}
		str := fmt.Sprintf("Shipping: %s, amount: [subtotal: %.2f USD, shipping: %.2f USD], tracking: [Service: %s, number: %s], ",
			order.Shipping.Name,
			float32(order.Amount.Subtotal)/100,
			float32(order.Amount.Shipping)/100,
			order.Tracking.Service,
			order.Tracking.Number,
		)

		items := "items: ["
		if len(order.Items) < 1 {
			items += ")"
		} else {
			for _, item := range order.Items {
				product, err := r.commandExecutor.FindProductByProductVariant(item.ProductVariantID)
				if err != nil {
					return "", err
				}

				items += fmt.Sprintf("(Name: %s, Quantity: %d, Price: %.2f USD, Variant: %s), ",
					product.Product.Name,
					item.Quantity,
					float32(item.Amount)/100,
					product.Variant.Name,
				)
			}
			items = items[:len(items)-2] + "]"
		}
		str += items + "]"

		return str, nil

	}
}
