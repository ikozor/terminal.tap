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
