package log

import (
	"fmt"
)
type ConsoleLog struct {
	
}

func LogConsoleDebug(file string) *ConsoleLog  {
	return &ConsoleLog{}
}

func (f * ConsoleLog) LogConsoleDebug(msg string)  {
	fmt.Println("console ",msg)
}

func (f *ConsoleLog) LogConsoleWran(msg string)  {
	fmt.Println("console",msg)
}
