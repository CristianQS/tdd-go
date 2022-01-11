package pkg

import "fmt"

type ConsoleLogger struct {
}

func (c ConsoleLogger) Log(message string) {
	fmt.Println(message)
}
