package xxlog

import (
	"fmt"
)
type XConsole struct {
	level int
}

func (c *XConsole) LogDebug(fm string,args... interface{})  {
	fmt.Printf("log debug of line")

}


func (c *XConsole) LogTrace(fm string,args... interface{})  {
	fmt.Printf("log trace od line")

}



func (c *XConsole) LogWarring(fm string,args... interface{})  {
	fmt.Printf("log warring od line")

}


func (c *XConsole) LogInfo(fm string,args... interface{})  {
	fmt.Printf("log info od line")

}

func (c *XConsole) LogError(fm string,args... interface{})  {
	fmt.Printf("log error od line")

}

func (c *XConsole) LogFatal(fm string,args... interface{})  {
	fmt.Printf("log fatal od line")

}

func (c *XConsole) SetLevel(level int)  {
	c.level =level

}
