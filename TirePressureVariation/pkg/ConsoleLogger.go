package pkg

import "fmt"

type ConsoleLogger struct {
}

func (c ConsoleLogger) log(message string) {
	fmt.Println(message)
}
