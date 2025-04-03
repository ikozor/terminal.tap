package repl

import "fmt"

func validateAddress(line []string) error {
	if len(line) < 1 {
		return fmt.Errorf("Invalid Address Error: Missing Address line")
	}
	if len(line) < 2 {
		return fmt.Errorf("Invalid Address: %s Error: Missing Address line", line[0])
	}
	return nil
}
