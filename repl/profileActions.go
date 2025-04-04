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

func (r *repl) setProfileEmail(email string) {
	r.args = email
	r.currentCommand = func(i interface{}) (string, error) {
		str, ok := i.(string)
		if !ok {
			return "", fmt.Errorf("email not string: %v", i)
		}

		if err := r.commandExecutor.SetProfileEmail(str); err != nil {
			return "", err
		}
		return "Successfully set email", nil
	}
}

func (r *repl) setProfileName(name string) {
	r.args = name
	r.currentCommand = func(i interface{}) (string, error) {
		str, ok := i.(string)
		if !ok {
			return "", fmt.Errorf("name not string: %v", i)
		}

		if err := r.commandExecutor.SetProfileName(str); err != nil {
			return "", err
		}
		return "Successfully set name", nil
	}
}
