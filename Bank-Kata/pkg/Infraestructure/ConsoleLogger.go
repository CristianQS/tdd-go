package Infraestructure

import "fmt"

type ConsoleLogger struct{}

func (ConsoleLogger) Print(message string) {
	fmt.Println(message)
}

