package repl

import (
	"fmt"
	"strings"

	"github.com/terminaldotshop/terminal-sdk-go"
)

func validAddress(addressString string) (terminal.AddressNewParams, error) {
	parts := strings.Split(addressString, ",")
	if len(parts) < 7 {
		return terminal.AddressNewParams{}, fmt.Errorf("Invalid Address: %s", addressString)
	}
	if len(parts) == 7 {
		return terminal.AddressNewParams{
			Name:     terminal.String(parts[0]),
			Street1:  terminal.String(parts[1]),
			City:     terminal.String(parts[2]),
			Province: terminal.String(strings.ReplaceAll(parts[3], " ", "")),
			Zip:      terminal.String(strings.ReplaceAll(parts[4], " ", "")),
			Country:  terminal.String(strings.ReplaceAll(parts[5], " ", "")),
			Phone:    terminal.String(strings.ReplaceAll(parts[6], " ", "")),
		}, nil
	}
	return terminal.AddressNewParams{
		Name:     terminal.String(parts[0]),
		Street1:  terminal.String(parts[1]),
		Street2:  terminal.String(parts[2]),
		City:     terminal.String(parts[3]),
		Province: terminal.String(strings.ReplaceAll(parts[4], " ", "")),
		Zip:      terminal.String(strings.ReplaceAll(parts[5], " ", "")),
		Country:  terminal.String(strings.ReplaceAll(parts[6], " ", "")),
		Phone:    terminal.String(strings.ReplaceAll(parts[7], " ", "")),
	}, nil
}

func (r *repl) AddAddressAction(address terminal.AddressNewParams) {
	r.args = address

	r.currentCommand = func(i interface{}) (string, error) {
		address, ok := i.(terminal.AddressNewParams)
		if !ok {
			return "", fmt.Errorf("Address type is not valid: %v", i)
		}
		if err := r.commandExecutor.AddAddress(address); err != nil {
			return "", err
		}
		return "Successfully added address", nil
	}
}
