package commands

import (
	"os"
	"path/filepath"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"unifiedsoftware.mf-generator/templates"
)

var InitCommand = &cobra.Command{
	Use:   "init",
	Short: "Initialize a configuration for Microfrontend with module federation",
	Run: func(cmd *cobra.Command, args []string) {
		err := os.WriteFile(filepath.Join(".", "us-mf.json"), []byte(templates.INIT_JSON), 0644)
		if err != nil {
			color.Red("Error writing us-mf.json: %v", err)
			return
		}
	},
}
