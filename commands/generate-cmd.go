package commands

import (
	"bufio"
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"unifiedsoftware.mf-generator/actions/csharp"
	"unifiedsoftware.mf-generator/actions/node"
	"unifiedsoftware.mf-generator/models"
)

var GenerateCommand = &cobra.Command{
	Use:   "generate",
	Short: "us-generator-mf is a Generator of Microfrontend",
	Run: func(cmd *cobra.Command, args []string) {
		reader := bufio.NewReader(os.Stdin)
		c := color.New(color.FgCyan)

		configData, err := ioutil.ReadFile("us-mf.json")
		if err != nil {
			color.Red("Error reading config file: %v", err)
			return
		}

		color.Yellow(string(configData))
		var config models.Config
		err = json.Unmarshal(configData, &config)
		if err != nil {
			color.Red("Error unmarshalling config file: %v", err)
			return
		}

		projectName := config.ProjectName
		port := config.Port
		base := config.Base
		repoBase := config.RepoBase
		language := config.Language
		spa := config.Spa
		proxy := config.Proxy
		remotes := config.Remotes
		serveStatic := config.ServeStatics

		if projectName == "" {
			c.Print("Enter the project name: (preferably in lowercase all) ")
			projectName, _ = reader.ReadString('\n')
			projectName = strings.TrimSpace(projectName)
		}

		if port == 0 {
			c.Print("Enter the port for the preview (default 3000): ")
			portInput, _ := reader.ReadString('\n')
			portInput = strings.TrimSpace(portInput)
			if portInput != "" {
				port, _ = strconv.Atoi(portInput)
			} else {
				port = 3000
			}
		}

		if base == "" {
			c.Print("Enter the base URL of the repository (default /): ")
			base, _ = reader.ReadString('\n')
			base = strings.TrimSpace(base)
			if base == "" {
				base = "/"
			}
		}

		if repoBase == "" {
			c.Print("Enter the Ubication project (default .): ")
			repoBase, _ = reader.ReadString('\n')
			repoBase = strings.TrimSpace(repoBase)
		}

		if language == "" {
			c.Print("Enter the language (csharp, typescript, javascript): ")
			language, _ = reader.ReadString('\n')
			language = strings.TrimSpace(language)

			if language != "csharp" && language != "typescript" && language != "javascript" {
				color.Red("Invalid language. Please enter 'csharp', 'typescript', or 'javascript'.")
				return
			}

		}

		generateProject(projectName, port, base, spa, repoBase, language, proxy, remotes, serveStatic, configData)
	},
}

func generateProject(projectName string, port int, base string, spa bool, repoBase, language string, proxy []models.ProxyConfig, remotes models.ModuleFederationConfig, serveStatic bool, configData []byte) {

	projectPath := filepath.Join(repoBase, projectName)

	if _, err := os.Stat(projectPath); !os.IsNotExist(err) {
		reader := bufio.NewReader(os.Stdin)
		color.Cyan("Project directory already exists. Do you want to delete it and create a new one? (y/N): ")
		deleteInput, _ := reader.ReadString('\n')
		deleteInput = strings.TrimSpace(deleteInput)
		if deleteInput == "y" || deleteInput == "Y" {
			err := os.RemoveAll(projectPath)
			if err != nil {
				color.Red("Error deleting project directory: %v", err)
				return
			}
			color.Green("Project directory deleted.")
		} else {
			color.Red("Project creation aborted.")
			return
		}
	}
	err := os.MkdirAll(filepath.Join(projectPath, "src"), 0755)
	if err != nil {
		color.Red("Error creating project directories: %v", err)
		return
	}

	switch language {
	case "csharp":
		csharp.CreateCSharpProject(projectPath)
	case "typescript":
		node.CreateTypeScriptProject(projectPath)
		node.CreateTsOrJsConfig(projectPath, language)
		node.ConfigViteAndSpaServer(projectPath, port, base, spa, language, projectName, proxy, remotes, serveStatic)
	case "javascript":
		node.CreateJavaScriptProject(projectPath)
		node.ConfigViteAndSpaServer(projectPath, port, base, spa, language, projectName, proxy, remotes, serveStatic)
	}
	os.WriteFile(filepath.Join(projectPath, "us-mf.json"), configData, 0644)

	color.Green("\nProject generated successfully in folder! " + projectName)
}
