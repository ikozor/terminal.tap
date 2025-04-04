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
