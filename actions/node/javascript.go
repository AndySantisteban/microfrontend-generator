package node

import (
	"os"
	"path/filepath"

	"unifiedsoftware.mf-generator/templates"
)

func CreateJavaScriptProject(projectPath string) {
	os.WriteFile(filepath.Join(projectPath, "index.html"), []byte(templates.INDEX_HTML_JSX_TEMPLATE), 0644)
	os.WriteFile(filepath.Join(projectPath, "src", "main.jsx"), []byte(templates.MAIN_JSX_TEMPLATE), 0644)
	os.WriteFile(filepath.Join(projectPath, "src", "App.jsx"), []byte(templates.APP_JSX_TEMPLATE), 0644)
}
