package main

import (
	"fmt"
	"greeter/cmd/root"
)

func main() {
	rootCmd := root.RootCommand()

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}

}
