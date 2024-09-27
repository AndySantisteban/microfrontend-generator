package models

type ProxyConfig struct {
	Path   string `json:"path"`
	Target string `json:"target"`
}
type ExposesModuleFederationConfig struct {
	Name      string  `json:"name"`
	RemoteUrl string  `json:"remoteUrl"`
	Rename    *string `json:"rename,omitempty"`
	Dir       *string `json:"dir,omitempty"`
}
type ModuleFederationConfig struct {
	Shared  []string                        `json:"shared"`
	Exposes []ExposesModuleFederationConfig `json:"exposes"`
}

type Config struct {
	ProjectName  string                 `json:"projectName"`
	Port         int                    `json:"port"`
	Base         string                 `json:"base"`
	RepoBase     string                 `json:"repoBase"`
	Language     string                 `json:"language"`
	Spa          bool                   `json:"spa"`
	ServeStatics bool                   `json:"serveStatics"`
	Proxy        []ProxyConfig          `json:"proxy"`
	Remotes      ModuleFederationConfig `json:"remotes"`
}
