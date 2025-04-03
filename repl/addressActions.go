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

func (r *repl) listAddresses() {
	r.args = nil
	r.currentCommand = func(i interface{}) (string, error) {
		addresses, err := r.commandExecutor.ListAddresses()
		if err != nil {
			return "", err
		}
		if len(addresses) < 1 {
			return "No Addresses saved", err
		}

		addressList := ""
		for _, address := range addresses {
			addressList += fmt.Sprintf("(Name: %s, Street 1: %s, ", address.Name, address.Street1)
			if address.Street2 != "" {
				addressList += fmt.Sprintf("Street 2: %s, ", address.Street2)
			}
			addressList += fmt.Sprintf("City: %s, State: %s, Zipcode: %s, Country: %s, Phone Number: %s), ",
				address.City,
				address.Province,
				address.Zip,
				address.Country,
				address.Phone,
			)

		}
		return addressList[:len(addressList)-2], nil
	}
}

func (r *repl) removeAddress(name string) {
	r.args = name

	r.currentCommand = func(i interface{}) (string, error) {
		str, ok := i.(string)
		if !ok {
			return "", fmt.Errorf("invalid address name: %v", i)
		}

		if err := r.commandExecutor.RemoveAddress(str); err != nil {
			return "", err
		}

		return fmt.Sprintf("Successfully Removed address: %s", name), nil
	}
}

func (r *repl) setAddress(name string) {
	r.args = name

	r.currentCommand = func(i interface{}) (string, error) {
		str, ok := i.(string)
		if !ok {
			return "", fmt.Errorf("invalid address name: %v", i)
		}

		if err := r.commandExecutor.SetAddress(str); err != nil {
			return "", err
		}
		return fmt.Sprintf("Addess %s set successfully", name), nil
	}
}
