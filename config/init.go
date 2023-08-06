package config

import (
	"fmt"
	"subjectInformation/service/validator"
)

func init() {
	initConfig()
	initLogger()
	if err := validator.InitValidator("zh"); err != nil {
		fmt.Println("Init translator failed.")
		return
	}
}
