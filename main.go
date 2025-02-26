package main

import (
	"flag"
	"fmt"
	"os"
	"yml2docker/templates"

	"github.com/jessevdk/go-flags"
)

type CommandArguments struct {
	BaseImage  string   `short:"b" long:"base-image" description:"Base image for the dockerfile." required:"true"`
	InputPath  string   `short:"i" long:"input-path" description:"Input path for the yml file." default:"./ci.yml"`
	OutputPath string   `short:"o" long:"output-path" description:"Output path of the folder including docker compose and services." default:"./export"`
	EnvVars    []string `short:"e" long:"env-vars" description:"Env vars to put into docker compose services."`
}

func main() {
	opts := CommandArguments{}
	_, err := flags.ParseArgs(&opts, os.Args)
	if err != nil {
		fmt.Printf("error parsing command arguments: %s\n", err)
		os.Exit(1)
	}

	if opts.BaseImage == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	if opts.InputPath == "" {
		opts.InputPath = "./ci.yml"
	}
	if opts.OutputPath == "" {
		opts.OutputPath = "./export"
	}

	fmt.Printf("base image: %s, input path: %s, output path: %s\n", opts.BaseImage, opts.InputPath, opts.OutputPath)

	// Get map from yml file
	ymlContent, err := templates.ReadYmlFile(opts.InputPath)
	if err != nil {
		fmt.Printf("error getting map from yml file: %s\n", err)
		os.Exit(1)
	}

	for serviceName, service := range ymlContent.Run {
		// Create Dockerfile
		fmt.Printf("creating dockerfile for service %s\n", serviceName)

		config := templates.DockerTemplateConfig{
			OutputPath:   opts.OutputPath + "/" + serviceName,
			BaseImage:    opts.BaseImage,
			PrepareSteps: ymlContent.Prepare.Steps,
			RunSteps:     service.Steps,
		}
		err = templates.CreateDockerfile(config)
		if err != nil {
			fmt.Printf("error creating dockerfile: %s\n", err)
			os.Exit(1)
		}
	}

	// Create nginx config
	fmt.Printf("creating nginx config file\n")

	configNginx := templates.NginxConfigTemplateConfig{
		OutputPath: opts.OutputPath,
		Services:   ymlContent.Run,
	}
	err = templates.CreateNginxConfig(configNginx)
	if err != nil {
		fmt.Printf("error creating docker compose file: %s\n", err)
		os.Exit(1)
	}

	// Create Docker compose file
	fmt.Printf("creating docker compose file\n")

	configDockerCompose := templates.DockerComposeTemplateConfig{
		OutputPath: opts.OutputPath,
		Services:   ymlContent.Run,
		EnvVars:    opts.EnvVars,
	}
	err = templates.CreateDockerCompose(configDockerCompose)
	if err != nil {
		fmt.Printf("error creating docker compose file: %s\n", err)
		os.Exit(1)
	}
}
