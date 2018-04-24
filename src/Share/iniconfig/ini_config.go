package iniconfig

import (
	//"fmt"
	"reflect"
	"errors"
	"strings"
	"fmt"
)
func Marshal(data interface{})(result []byte,err error)  {
	return
	
}

func UnMarshal(data []byte,result interface{})(err error)  {
	lineArray := strings.Split(string(data),"\n")

	typeInfo := reflect.TypeOf(result)
	if typeInfo.Kind() != reflect.Ptr{
		err = errors.New("please pass address")
	}
	typeInfoTwo := typeInfo.Elem()
	if typeInfoTwo.Kind() != reflect.Struct{
		err = errors.New("please pass struct")
		return
	}
	for index,line := range lineArray{
		line = strings.TrimSpace(line)
		if len(line) == 0{
			continue
		}
		if line[0] == ';'|| line[0] == '#'{
			continue
		}
		if line[0] == '[' && len(line)<= 2{
			err = fmt.Errorf("syntax error,invalid section:%s,lineNo:%d",line,index+1)
			return
		}
		if line[0] == '[' && line[len(line)-1] != ']'{
			err = fmt.Errorf("syntax error,invalid section:%s,lineNo:%d",line,index+1)
			return
		}
		sectionName := strings.TrimSpace(line[1:len(line)-1])
		if len(sectionName) == 0{
			err = fmt.Errorf("syntax error,invalid section:%s,lineNo:%d",line,index+1)
			return

		}

	}
	return



}