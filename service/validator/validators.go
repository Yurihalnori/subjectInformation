package validator

import (
	"time"

	"github.com/go-playground/validator/v10"
)

func timing(fl validator.FieldLevel) bool {
	if date, ok := fl.Field().Interface().(time.Time); ok {
		today := time.Now()
		if today.After(date) {
			return true
		}
	}
	return false
}

func category(fl validator.FieldLevel) bool {
	if num, ok := fl.Field().Interface().(string); ok {
		if len(num) != 10 {
			return false
		}
		for i := 0; i < len(num); i++ {
			if num[i] != '1' && num[i] != '0' {
				return false
			}
		}
		return true
	}
	return false
}
