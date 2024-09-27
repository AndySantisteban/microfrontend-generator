package node

import (
	"os"
	"path/filepath"

	"unifiedsoftware.mf-generator/templates"
)

func CreateTypeScriptProject(projectPath string) {
	os.WriteFile(filepath.Join(projectPath, "index.html"), []byte(templates.INDEX_HTML_TSX_TEMPLATE), 0644)
	os.WriteFile(filepath.Join(projectPath, "src", "main.tsx"), []byte(templates.MAIN_TSX_TEMPLATE), 0644)
	os.WriteFile(filepath.Join(projectPath, "src", "App.tsx"), []byte(templates.APP_TSX_TEMPLATE), 0644)
}
