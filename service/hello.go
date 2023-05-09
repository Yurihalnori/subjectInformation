package service

import (
	"fmt"
	"time"
)

type HelloService struct {
}

func (h HelloService) Hello(msg string) string {
	return fmt.Sprintf("hello %v", msg)
}

func (h HelloService) HelloTime(date time.Time) string {
	return fmt.Sprintf("tomorrow is %v", date.Format("2006-01-02"))
}
