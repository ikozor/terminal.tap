package repl

import "fmt"

func (r *repl) listCards() {
	r.args = nil

	r.currentCommand = func(i interface{}) (string, error) {
		cards, err := r.commandExecutor.ListCards()
		if err != nil {
			return "", err
		}
		if len(cards) < 1 {
			return "No cards saved", nil
		}

		cardsList := "("
		for _, e := range cards {
			cardsList += fmt.Sprintf("%s, Last 4: %s, Exp: %d/%d), ",
				e.Brand,
				e.Last4,
				e.Expiration.Month,
				e.Expiration.Year,
			)
		}

		return cardsList[:len(cardsList)-2], nil
	}
}

func (r *repl) addCard() {
	r.args = nil
	r.currentCommand = func(i interface{}) (string, error) {
		url, err := r.commandExecutor.AddCard()
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("Go to %s to add the card", url), nil
	}
}

func (r *repl) removeCard(last4 string) {
	r.args = last4
	r.currentCommand = func(i interface{}) (string, error) {
		str, ok := i.(string)
		if !ok {
			return "", fmt.Errorf("Invalid arg for last4: %v", i)
		}
		if err := r.commandExecutor.RemoveCard(str); err != nil {
			return "", err
		}
		return fmt.Sprintf("Card %s successfully removed", last4), nil
	}
}

func (r *repl) setCard(last4 string) {
	r.args = last4
	r.currentCommand = func(i interface{}) (string, error) {
		str, ok := i.(string)
		if !ok {
			return "", fmt.Errorf("Invalid arg for last4: %v", i)
		}
		if err := r.commandExecutor.SetCard(str); err != nil {
			return "", err
		}
		return fmt.Sprintf("Card %s successfully set", last4), nil
	}
}
