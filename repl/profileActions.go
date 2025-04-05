package repl

import "fmt"

func (r *repl) getProfile() {
	r.currentCommand = func() (string, error) {
		profile, err := r.commandExecutor.GetProfile()
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("Name: %s, Email: %s", profile.Name, profile.Email), nil

	}
}

func (r *repl) setProfileEmail(email string) {
	r.currentCommand = func() (string, error) {
		if err := r.commandExecutor.SetProfileEmail(email); err != nil {
			return "", err
		}
		return "Successfully set email", nil
	}
}

func (r *repl) setProfileName(name string) {
	r.currentCommand = func() (string, error) {
		if err := r.commandExecutor.SetProfileName(name); err != nil {
			return "", err
		}
		return "Successfully set name", nil
	}
}
