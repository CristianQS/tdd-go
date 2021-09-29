package Infraestructure

import "time"

type CoreTimeProvider struct{}

func (c CoreTimeProvider) GetTodayDate() string {
	return time.Now().Format("02/01/2006")
}
