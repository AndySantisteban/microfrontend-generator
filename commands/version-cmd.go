package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

var VersionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of us-generator-mf",
	Long:  `All software has versions. This is us-generator-mf's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("us-generator-mf Generator of Microfrontend v0.1 -- HEAD")
	},
}
