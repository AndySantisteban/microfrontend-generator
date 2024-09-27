package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"unifiedsoftware.mf-generator/commands"
)

var rootCmd = &cobra.Command{
	Short: "Unified Software Generator of Microfrontend",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to Unified Software Generator of Microfrontend")
		fmt.Println("Execute ./us-generator-mf help")
	},
}

func Execute() {
	c := color.New(color.FgCyan)
	data, err := ioutil.ReadFile("logo.txt")
	if err != nil {
		c.Print(err)
	}

	fmt.Println(string(data))
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(commands.InitCommand)
	rootCmd.AddCommand(commands.VersionCmd)
	rootCmd.AddCommand(commands.GenerateCommand)
	rootCmd.AddCommand(commands.UpdateCommand)
}

func main() {
	Execute()
}
