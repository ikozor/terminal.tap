package main

import (
	"context"
	"fmt"

	"github.com/terminaldotshop/terminal-sdk-go"
	"github.com/terminaldotshop/terminal-sdk-go/option"
)

func main() {
	client := terminal.NewClient(
		option.WithBaseURL("https://api.dev.terminal.shop"), // the Double Slash was causing panic 
	)
	product, err := client.Product.List(context.TODO())
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("%+v\n", product.Data)
}
