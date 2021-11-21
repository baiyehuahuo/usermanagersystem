package configread

import (
	"io/ioutil"
	"os"
	"usermanagersystem/model"

	"gopkg.in/yaml.v2"
)

var Config model.ConfigModel

func ConfigRead() error {
	file, err := os.Open("configs/config.yaml")
	if err != nil {
		return err
	}
	yamlFile, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(yamlFile, &Config)
	return err
}
