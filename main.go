package main

import (
	"flag"
	"fmt"
	"os"
	"yml2docker/templates"
)

func main() {
	baseImage := flag.String("b", "", "Base image for the dockerfile. (Required)")
	inputPath := flag.String("i", "", "Input path for the yml file. Default is './ci.yml'.")
	outputPath := flag.String("o", "", "Output path of the folder including docker compose and services. Default is './export'.")
	flag.Parse()

	if *baseImage == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	if *inputPath == "" {
		*inputPath = "./ci.yml"
	}
	if *outputPath == "" {
		*outputPath = "./export"
	}

	fmt.Printf("base image: %s, input path: %s, output path: %s\n", *baseImage, *inputPath, *outputPath)

	// Get map from yml file
	ymlContent, err := templates.ReadYmlFile(*inputPath)
	if err != nil {
		fmt.Printf("error getting map from yml file: %s\n", err)
		os.Exit(1)
	}

	for serviceName, service := range ymlContent.Run {
		// Create Dockerfile
		fmt.Printf("creating dockerfile for service %s\n", serviceName)

		config := templates.DockerTemplateConfig{
			OutputPath:   *outputPath + "/" + serviceName,
			BaseImage:    *baseImage,
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
		OutputPath: *outputPath,
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
		OutputPath: *outputPath,
		Services:   ymlContent.Run,
	}
	err = templates.CreateDockerCompose(configDockerCompose)
	if err != nil {
		fmt.Printf("error creating docker compose file: %s\n", err)
		os.Exit(1)
	}
}
