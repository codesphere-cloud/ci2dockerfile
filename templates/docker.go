package templates

import (
	"fmt"
	"os"
	"text/template"
	"yml2docker/model"
)

type DockerTemplateConfig struct {
	OutputPath string
	// Dockerfile congifuration
	BaseImage    string
	PrepareSteps []model.Step
	RunSteps     []model.Step
}

func CreateDockerfile(config DockerTemplateConfig) error {
	err := CreateDirectory(config.OutputPath)
	if err != nil {
		return fmt.Errorf("error creating directory: %w\n", err)
	}

	// Create Dockerfile
	f, err := os.Create(config.OutputPath + "/Dockerfile")
	if err != nil {
		return fmt.Errorf("error creating docker file: %w\n", err)
	}

	dockerTemplate, err := template.ParseFiles("./helper/docker.tmpl")
	if err != nil {
		return fmt.Errorf("error parsing docker template: %w\n", err)
	}

	err = dockerTemplate.Execute(f, config)
	if err != nil {
		return fmt.Errorf("error executing docker template: %w\n", err)
	}

	// Create shell script for entrypoint
	f, err = os.Create(config.OutputPath + "/entrypoint.sh")
	if err != nil {
		return fmt.Errorf("error creating entrypoint.sh: %w\n", err)
	}

	entrypointTemplate, err := template.ParseFiles("./helper/shell.tmpl")
	if err != nil {
		return fmt.Errorf("error parsing shell template: %w\n", err)
	}

	err = entrypointTemplate.Execute(f, config)
	if err != nil {
		return fmt.Errorf("error executing shell template: %w\n", err)
	}

	return f.Close()
}
