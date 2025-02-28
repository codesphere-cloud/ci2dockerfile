package templates

import (
	_ "embed"
	"fmt"
	"os"
	"text/template"
	"yml2docker/model"
)

//go:embed dockercompose.tmpl
var DockerComposeTemplateFile string

type DockerComposeTemplateConfig struct {
	OutputPath string
	// Docker compose configuration
	Services map[string]model.Service
	EnvVars  []string
}

func CreateDockerCompose(config DockerComposeTemplateConfig) error {
	t, err := template.New("dockercompose.tmpl").Parse(DockerComposeTemplateFile)
	if err != nil {
		return fmt.Errorf("error parsing docker compose template: %w\n", err)
	}

	err = CreateDirectory(config.OutputPath)
	if err != nil {
		return fmt.Errorf("error creating directory: %w\n", err)
	}

	f, err := os.Create(config.OutputPath + "/docker-compose.yml")
	if err != nil {
		return fmt.Errorf("error creating docker compose file: %w\n", err)
	}

	err = t.Execute(f, config)
	if err != nil {
		return fmt.Errorf("error executing docker compose template: %w\n", err)
	}

	return f.Close()
}
