package encoding

import (
	"encoding/json"

	"github.com/hive-bootcamp/final-project-encoding-go/models"

	"gopkg.in/yaml.v3"

	"os"
)

// JSONData тип для перекодирования из JSON в YAML
type JSONData struct {
	DockerCompose *models.DockerCompose
	FileInput     string
	FileOutput    string
}

// YAMLData тип для перекодирования из YAML в JSON
type YAMLData struct {
	DockerCompose *models.DockerCompose
	FileInput     string
	FileOutput    string
}

// MyEncoder интерфейс для структур YAMLData и JSONData
type MyEncoder interface {
	Encoding() error
}

// Encoding перекодирует файл из JSON в YAML
func (j *JSONData) Encoding() error {
	data, err := os.ReadFile(j.FileInput)
	if err != nil {
		return err
	}
	if err = json.Unmarshal(data, &j.DockerCompose); err != nil {
		return err
	}
	body, err := yaml.Marshal(j.DockerCompose)
	if err != nil {
		return err
	}
	return os.WriteFile(j.FileOutput, body, 0644)
}

// Encoding перекодирует файл из YAML в JSON
func (y *YAMLData) Encoding() error {
	data, err := os.ReadFile(y.FileInput)
	if err != nil {
		return err
	}
	if err = yaml.Unmarshal(data, &y.DockerCompose); err != nil {
		return err
	}
	body, err := json.Marshal(y.DockerCompose)
	if err != nil {
		return nil
	}
	return os.WriteFile(y.FileOutput, body, 0644)
}
