package templates

import (
	"fmt"
	"os"
	"yml2docker/model"

	"gopkg.in/yaml.v2"
)

func ReadYmlFile(inputPath string) (*model.CiYml, error) {
	ymlFile, err := os.ReadFile(inputPath)
	if err != nil {
		return nil, fmt.Errorf("error reading yml file: %w\n", err)
	}

	ymlContent := &model.CiYml{}
	err = yaml.Unmarshal(ymlFile, &ymlContent)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling yml file: %w\n", err)
	}

	return ymlContent, nil
}
