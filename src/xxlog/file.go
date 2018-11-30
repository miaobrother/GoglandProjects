package xxlog

import (
	"fmt"
)
type XFile struct {
	level int
	filename string
}

func (f *XFile) LogDebug(fm string,args... interface{})  {
	fmt.Printf("log debug od line")
	
}


func (f *XFile) LogTrace(fm string,args... interface{})  {
	fmt.Printf("log trace od line")

}



func (f *XFile) LogWarring(fm string,args... interface{})  {
	fmt.Printf("log warring od line")

}


func (f *XFile) LogInfo(fm string,args... interface{})  {
	fmt.Printf("log info od line")

}

func (f *XFile) LogError(fm string,args... interface{})  {
	fmt.Printf("log eror od line")

}

func (f *XFile) LogFatal(fm string,args... interface{})  {
	fmt.Printf("log fatal od line")

}

func (f *XFile) SetLevel(level int)  {
	f.level=level

}
