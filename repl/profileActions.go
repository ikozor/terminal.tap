package repl

import "fmt"

func (r *repl) getProfile() {
	r.args = nil
	r.currentCommand = func(i interface{}) (string, error) {
		profile, err := r.commandExecutor.GetProfile()
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("Name: %s, Email: %s", profile.Name, profile.Email), nil

	}
}
