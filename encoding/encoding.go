package encoding

import (
	"github.com/hive-bootcamp/final-project-encoding-go/models"
	"github.com/hive-bootcamp/final-project-encoding-go/utils"
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
	err := utils.ConvertJsonToYaml("artist.json", "artist.yaml")
	if err != nil {
		return err
	}

	return nil
}

// Encoding перекодирует файл из YAML в JSON
func (y *YAMLData) Encoding() error {
	err := utils.ConvertYamlToJson("artist.yaml", "artist.json")
	if err != nil {
		return err
	}

	return nil
}
