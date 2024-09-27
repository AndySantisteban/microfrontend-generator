package node

import (
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"time"

	"unifiedsoftware.mf-generator/models"
)

func getRandomPort(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min+1) + min
}

func CreateEnviromentValues(projectPath string, configuration models.Config) {

	if configuration.Spa {
		var envVars []string
		for _, entries := range configuration.Remotes.Exposes {

			value := fmt.Sprintf(`
%s="%s"
%s_URL="%s"
%s_DIR="%s"
`,
				entries.Name, *entries.Rename,
				entries.Name, entries.RemoteUrl,
				entries.Name, *entries.Dir,
			)
			envVars = append(envVars, value)
		}
		envVars = append(envVars, fmt.Sprintf("PORT=%d", getRandomPort(3000, 8000)))
		os.WriteFile(filepath.Join(projectPath, ".env"), []byte(strings.Join(envVars, "\n")), 0644)
		os.WriteFile(filepath.Join(projectPath, ".env.local"), []byte(strings.Join(envVars, "\n")), 0644)
		os.WriteFile(filepath.Join(projectPath, ".env.development"), []byte(strings.Join(envVars, "\n")), 0644)
		os.WriteFile(filepath.Join(projectPath, ".env.integration"), []byte(strings.Join(envVars, "\n")), 0644)
		os.WriteFile(filepath.Join(projectPath, ".env.production"), []byte(strings.Join(envVars, "\n")), 0644)

	}

}
