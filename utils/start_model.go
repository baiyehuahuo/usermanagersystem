package utils

import (
	"log"
	"os/exec"
	"strconv"
	"usermanagersystem/consts"
)

func StartModel() {
	config := Config.RabbitMQConfig
	args := []string{
		config.Host, strconv.Itoa(config.Port), config.Account, config.Password, // ip port account password
		consts.PredictQueueName, consts.ExchangeName, strconv.FormatBool(consts.AutoDelete), strconv.Itoa(consts.MaxQueueLength), // predictQueueName exhcnageName
	}
	cmd := exec.Command("main.exe", args...)
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}
