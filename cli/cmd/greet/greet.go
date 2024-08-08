package greet

import (
	"fmt"

	"github.com/spf13/cobra"
)

func GreetCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "greet",
		Short: "Greet the user",
		Long:  "Greet the user with a friendly message",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Hello!", args[0])
		},
	}
}
