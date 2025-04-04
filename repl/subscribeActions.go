package repl

import (
	"fmt"
)

func (r *repl) listSubscriptions() {
	r.args = nil
	r.currentCommand = func(i interface{}) (string, error) {
		subscriptions, err := r.commandExecutor.ListSubscriptions()
		if err != nil {
			return "", err
		}
		if len(subscriptions) < 1 {
			return "No subscriptions found", nil
		}
		str := ""
		for i, s := range subscriptions {
			product, err := r.commandExecutor.FindProductByProductVariant(s.ProductVariantID)
			if err != nil {
				return "", err
			}

			str += fmt.Sprintf("(Id: %d, Schedule: [Type: %s, Interval: %d], Product: [Name: %s, Variant: %s, Quantity: %d], Next: %s), ",
				i,
				s.Schedule.Type,
				s.Schedule.Interval,
				product.Product.Name,
				product.Variant.Name,
				s.Quantity,
				s.Next,
			)
		}
		return str[:len(str)-2], nil
	}

}
