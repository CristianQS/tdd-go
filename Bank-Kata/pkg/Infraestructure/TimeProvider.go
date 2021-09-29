package Infraestructure

type TimeProvider interface {
	GetTodayDate() string
}
