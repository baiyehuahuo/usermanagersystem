package utils

import (
	"io/ioutil"
	"os"
	"usermanagersystem/model"

	"gopkg.in/yaml.v2"
)

var Config model.ConfigModel

func ConfigRead() (err error) {
	file, err := os.Open("configs/config.yaml")
	if err != nil {
		return ErrWrapOrWithMessage(true, err)
	}
	yamlFile, err := ioutil.ReadAll(file)
	if err != nil {
		return ErrWrapOrWithMessage(true, err)
	}

	err = yaml.Unmarshal(yamlFile, &Config)
	if err != nil {
		return ErrWrapOrWithMessage(true, err)
	}
	return nil
}
