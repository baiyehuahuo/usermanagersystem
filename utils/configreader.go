package utils

import (
	"io/ioutil"
	"os"
	"usermanagersystem/model"

	"github.com/pkg/errors"

	"gopkg.in/yaml.v2"
)

var Config model.ConfigModel

func ConfigRead() (err error) {
	file, err := os.Open("configs/config.yaml")
	if err != nil {
		return errors.Wrap(err, RunFuncNameWithFail())
	}
	yamlFile, err := ioutil.ReadAll(file)
	if err != nil {
		return errors.Wrap(err, RunFuncNameWithFail())
	}

	err = yaml.Unmarshal(yamlFile, &Config)
	if err != nil {
		return errors.Wrap(err, RunFuncNameWithFail())
	}
	return nil
}
