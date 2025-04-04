package commands

import (
	"context"

	"github.com/terminaldotshop/terminal-sdk-go"
)

func (c *CommandExecutor) GetProfile() (terminal.ProfileUser, error) {
	res, err := c.client.Profile.Me(context.TODO())
	if err != nil {
		return terminal.ProfileUser{}, getApiErrorMessage(err)
	}
	return res.Data.User, nil
}

func (c *CommandExecutor) SetProfileEmail(email string) error {
	profile, err := c.GetProfile()
	if err != nil {
		return err
	}
	_, err = c.client.Profile.Update(context.TODO(),
		terminal.ProfileUpdateParams{Email: terminal.String(email), Name: terminal.String(profile.Name)})
	if err != nil {
		return getApiErrorMessage(err)
	}
	return nil

}

func (c *CommandExecutor) SetProfileName(name string) error {
	profile, err := c.GetProfile()
	if err != nil {
		return err
	}

	_, err = c.client.Profile.Update(context.TODO(),
		terminal.ProfileUpdateParams{Name: terminal.String(name), Email: terminal.String(profile.Email)})
	if err != nil {
		return getApiErrorMessage(err)
	}
	return nil
}
