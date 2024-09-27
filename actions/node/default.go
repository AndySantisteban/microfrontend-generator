package node

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/fatih/color"
	node "unifiedsoftware.mf-generator/actions/env"
	"unifiedsoftware.mf-generator/models"
	"unifiedsoftware.mf-generator/templates"
	"unifiedsoftware.mf-generator/utils"
)

func CreateTsOrJsConfig(projectPath, language string) {
	os.WriteFile(filepath.Join(projectPath, utils.GetLanguageExtension(language)+"config.json"), []byte(templates.TSCONFIG_TEMPLATE), 0644)
	os.WriteFile(filepath.Join(projectPath, utils.GetLanguageExtension(language)+"config.node.json"), []byte(templates.TSCONFIG_NODE_TEMPLATE), 0644)
}

func ConfigViteAndSpaServer(projectPath string, port int, base string, spa bool, language string, projectName string, proxy []models.ProxyConfig, remotes models.ModuleFederationConfig, serveStatics bool) {
	data := buildProxyConfig(proxy)

	remoteEntries, shared := buildRemotesEntriesConfig(remotes, serveStatics)
	viteConfig := fmt.Sprintf(
		templates.ViteConfigTemplate,
		fmt.Sprintf("'%s'", strings.Title(strings.ToLower(projectName))),
		utils.GetFileExtension(language),
		remoteEntries,
		shared,
		port,
		data,
		base,
	)

	err := os.WriteFile(filepath.Join(projectPath, "vite.config."+utils.GetLanguageExtension(language)), []byte(viteConfig), 0644)
	if err != nil {
		color.Red("Error writing vite.config.ts: %v", err)
		return
	}
	createPackageJson(projectPath, language, spa, projectName)
	if spa {
		createServerSPA(projectPath, base, port, remotes, serveStatics)
	}

}

func buildStaticRoutesConfig(exposes models.ExposesModuleFederationConfig, isBasePath bool) string {
	if isBasePath {
		return `
			app.use('/', express.static(path.join(__dirname, './dist')));
			app.get('*', (_, res) => {
				res
				.status(200)
				.sendFile(path.join(__dirname, './dist' + "/index.html"));
			});`
	}
	return fmt.Sprintf(`
	app.use('%s', express.static(path.join(__dirname, '%s')));
	app.get('%s', (_, res) => {
		res
		.status(200)
		.sendFile(path.join(__dirname, '%s' + "/index.html"));
	});`, *exposes.Rename, *exposes.Dir, *exposes.Rename, *exposes.Dir)
}

func buildProxyRemoteConfig(remotes models.ExposesModuleFederationConfig) string {
	return fmt.Sprintf("\t\t{ base:  process.env.%s, url: process.env.%s_URL }", remotes.Name, remotes.Name)
}

func createServerSPA(projectPath, base string, port int, remoteEntries models.ModuleFederationConfig, serveStatics bool) {
	node.CreateEnviromentValues(projectPath, models.Config{
		Spa:     true,
		Remotes: remoteEntries,
	})
	var remote []string

	for _, valuesRemote := range remoteEntries.Exposes {
		values := buildProxyRemoteConfig(valuesRemote)
		remote = append(remote, values)
	}
	remotesValues := strings.Join(remote, ",\n")

	var static []string
	for _, valuesStatic := range remoteEntries.Exposes {
		values := buildStaticRoutesConfig(valuesStatic, false)
		static = append(static, values)
	}
	values := buildStaticRoutesConfig(models.ExposesModuleFederationConfig{}, true)
	static = append(static, values)
	staticValues := strings.Join(static, "\n")

	expressServer := `import dotenv from 'dotenv'
import express from 'express';
import path from 'path';
import { fileURLToPath } from 'url';
import cors from 'cors';
import { createProxyMiddleware } from 'http-proxy-middleware';

const app = express();
const __filename = fileURLToPath(import.meta.url);
const __dirname = path.dirname(__filename);
const isStaticFilesSPA =  %t
dotenv.config()


const PORT = process.env.PORT || 8888;

app.use(cors());


if(isStaticFilesSPA) {
	%s
} else {
	const createProxyOptions = (appBase, appUrl) => {
		return {
			target: appUrl + appBase,
			changeOrigin: true,
			pathRewrite: {
				['^' + appBase]: appBase,
			},
		};
	};
 	const proxyConfigs = [	
%s
	]

	proxyConfigs.forEach(({ base, url }) => {
		const proxyOptions = createProxyOptions(base, url);
		app.use(base, createProxyMiddleware(proxyOptions));
	});

	app.use('%s', createProxyMiddleware({
		target: 'http://localhost:' + %d + '%s',
		changeOrigin: true,
		pathRewrite: {
			['^' + '%s']: '%s',
		},
	},))
}

app.listen(PORT, () => {
    console.log("Server is running on port: " + PORT);
});

`
	expressServer = fmt.Sprintf(expressServer, serveStatics, staticValues, remotesValues, base, port, base, base, base)
	err := os.WriteFile(filepath.Join(projectPath, "server.js"), []byte(expressServer), 0644)

	if err != nil {
		color.Red("Error writing server.js: %v", err)
		return
	}
}

func createPackageJson(projectPath, language string, isMain bool, projectName string) {
	packageJson := templates.PACKAGE_JSON_DEFAULT

	if language == "typescript" {
		packageJson = templates.PACKAGEJSON_WITH_TYPESCRIPT
	}

	scriptServe := ""
	devDependenciesServe := ""
	dependencies := ""

	if isMain {
		if language == "typescript" {
			scriptServe = "\"serve\":\"tsc && vite build && node server.js\","
			devDependenciesServe = "\"@types/express\": \"^4.17.21\",\n\t\"@types/cors\": \"^2.8.5\","

		} else {
			scriptServe = "\"serve\":\"vite build && node server.js\","
		}
		dependencies = "\"cors\": \"^2.8.5\",\n\t\"dotenv\": \"^16.4.5\",\n\t\"express\": \"^4.19.2\",\n\t\"http-proxy-middleware\": \"^3.0.0\""
	}
	if language == "typescript" {
		packageJson = fmt.Sprintf(packageJson, fmt.Sprintf("%s", strings.Title(strings.ToLower(projectName))), scriptServe, dependencies, devDependenciesServe)
	} else {
		packageJson = fmt.Sprintf(packageJson, fmt.Sprintf("%s", strings.Title(strings.ToLower(projectName))), scriptServe)
	}

	err := os.WriteFile(filepath.Join(projectPath, "package.json"), []byte(packageJson), 0644)
	if err != nil {
		color.Red("Error writing package.json: %v", err)
		return
	}
}

func buildProxyConfig(proxies []models.ProxyConfig) string {
	var proxyEntries []string
	for _, proxy := range proxies {
		entry := fmt.Sprintf(`
		'%s': {
				target: '%s',
				ws: true,
				changeOrigin: true,
				secure: false,
				rewrite: (path) => path.replace(/^\%s/, '%s')
	  }`, proxy.Path, proxy.Target, proxy.Path, proxy.Path)
		proxyEntries = append(proxyEntries, entry)
	}
	return strings.Join(proxyEntries, ",\n")
}

func buildRemotesEntriesConfig(modules models.ModuleFederationConfig, serveStatics bool) (string, string) {
	var sharedEntries []string
	for _, values := range modules.Shared {
		entry := fmt.Sprintf(`'%s'`, values)
		sharedEntries = append(sharedEntries, entry)
	}
	sharedValues := fmt.Sprintf("[%s]", strings.Join(sharedEntries, ","))

	var remotesEntries []string
	for _, values := range modules.Exposes {
		if serveStatics {
			entry := fmt.Sprintf(`
		%s: '%s'`, values.Name, values.RemoteUrl)
			remotesEntries = append(remotesEntries, entry)
		} else {
			entry := fmt.Sprintf(`
		%s: '%s'`, values.Name, *values.Rename+"/assets/remoteEntry.js")
			remotesEntries = append(remotesEntries, entry)
		}

	}

	return strings.Join(remotesEntries, ",\n"), sharedValues
}
