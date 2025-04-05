package repl

import (
	"fmt"

	"github.com/ikozor/terminal.tap/commands"
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
			product, err := r.commandExecutor.FindProductByTerminalProductVariantId(s.ProductVariantID)
			if err != nil {
				return "", err
			}

			if len(s.Next) < 10 {
				s.Next = "0000-00-00"
			}
			str += fmt.Sprintf("(Id: %d, Schedule: [Type: %s, Interval: %d], Product: [Name: %s, Variant: %s, Quantity: %d], Next: %s), ",
				i,
				s.Schedule.Type,
				s.Schedule.Interval,
				product.Product.Name,
				product.Variant.Name,
				s.Quantity,
				s.Next[:10],
			)
		}
		return str[:len(str)-2], nil
	}

}

func (r *repl) getSubscription(id int) {
	r.args = id
	r.currentCommand = func(i interface{}) (string, error) {
		id, ok := i.(int)
		if !ok {
			return "", fmt.Errorf("Id invalid format: %v", i)
		}
		subscription, err := r.commandExecutor.GetSubscription(id)
		if err != nil {
			return "", err
		}
		product, err := r.commandExecutor.FindProductByTerminalProductVariantId(subscription.Subscription.ProductVariantID)
		if err != nil {
			return "", err
		}

		if len(subscription.Subscription.Next) < 10 {
			subscription.Subscription.Next = "0000-00-00"
		}

		str := fmt.Sprintf("Schedule: [Type: %s, Interval: %d], Product: [Name: %s, Variant: %s, Quantity: %d], Next: %s, ",
			subscription.Subscription.Schedule.Type,
			subscription.Subscription.Schedule.Interval,
			product.Product.Name,
			product.Variant.Name,
			subscription.Subscription.Quantity,
			subscription.Subscription.Next[:10],
		)

		str += fmt.Sprintf("Address : [Name: %s, Street 1: %s, ", subscription.Address.Name, subscription.Address.Street1)
		if subscription.Address.Street2 != "" {
			str += fmt.Sprintf("Street 2: %s, ", subscription.Address.Street2)
		}
		str += fmt.Sprintf("City: %s, State: %s, Zipcode: %s, Country: %s, Phone Number: %s], ",
			subscription.Address.City,
			subscription.Address.Province,
			subscription.Address.Zip,
			subscription.Address.Country,
			subscription.Address.Phone,
		)

		str += fmt.Sprintf("Card: [%s, Last 4: %s, Exp: %d/%d]",
			subscription.Card.Brand,
			subscription.Card.Last4,
			subscription.Card.Expiration.Month,
			subscription.Card.Expiration.Year,
		)

		return str, nil

	}
}

func (r *repl) addSubscription(newSub commands.NewSubscription) {
	r.args = newSub

	r.currentCommand = func(i interface{}) (string, error) {
		newSubscription, ok := i.(commands.NewSubscription)
		if !ok {
			return "", fmt.Errorf("Subscripton type invalid: %v", i)
		}
		if err := r.commandExecutor.AddSubscription(newSubscription); err != nil {
			return "", err
		}
		return "Subscripton successfully added", nil
	}

}

func (r *repl) removeSubscription(id int) {
	r.args = id
	r.currentCommand = func(i interface{}) (string, error) {
		id, ok := i.(int)
		if !ok {
			return "", fmt.Errorf("id must be int: %v", i)
		}
		if err := r.commandExecutor.RemoveSubscription(id); err != nil {
			return "", err
		}
		return fmt.Sprintf("Subscription %d successfully removed", id), nil
	}
}
