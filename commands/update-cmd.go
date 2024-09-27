package commands

import "github.com/spf13/cobra"

var UpdateCommand = &cobra.Command{
	Use:   "update",
	Short: "Update a Project with config file 'us-mf.json'",
}
