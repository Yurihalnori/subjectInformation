package config

import (
	"subjectInformation/service/validator"
)

func init() {
	initConfig()
	initLogger()
	validator.InitValidator()
}
