package infraestructure

import "fmt"

type ConsoleLogger struct{}

func (c *ConsoleLogger) PrintLine(message string){
	fmt.Println(message)
}

