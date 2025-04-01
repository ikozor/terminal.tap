package commands

import (
	"github.com/terminaldotshop/terminal-sdk-go"
	"github.com/terminaldotshop/terminal-sdk-go/option"
)

type commandExecutor struct {
	client *terminal.Client
}

func CreateCommandExecutor() *commandExecutor {

	client := terminal.NewClient(
		option.WithBaseURL("https://api.dev.terminal.shop"), // the Double Slash was causing panic
	)
	return &commandExecutor{client: client}

}
