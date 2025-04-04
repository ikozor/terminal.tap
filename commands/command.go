package commands

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/terminaldotshop/terminal-sdk-go"
	"github.com/terminaldotshop/terminal-sdk-go/option"
)

type CommandExecutor struct {
	client          *terminal.Client
	currentProducts []terminal.Product
	currentAddress  string
	currentCard     string
}

func CreateCommandExecutor() *CommandExecutor {

	client := terminal.NewClient(
		option.WithBaseURL("https://api.dev.terminal.shop"), // the Double Slash was causing panic
	)
	return &CommandExecutor{client: client}

}

func getApiErrorMessage(err error) error {
	bodyIndex := strings.Index(err.Error(), "{")
	bodyString := err.Error()[bodyIndex:]

	body := map[string]string{}
	json.Unmarshal([]byte(bodyString), &body)
	message, ok := body["message"]
	if !ok {
		return fmt.Errorf("api error, cannot find message in api")
	}
	return fmt.Errorf("%s", message)

}
