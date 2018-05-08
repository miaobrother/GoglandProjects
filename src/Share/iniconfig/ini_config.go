package iniconfig

import (
	//"fmt"
	"errors"
	"fmt"
	"reflect"
	"strings"
	"strconv"
	"io/ioutil"
	//"os"
)

func MarshalFile(data interface{},filename string)(err error)  {
	result,err := Marshal(data)
	if err != nil{
		return
	}
	return ioutil.WriteFile(filename,result,0755)
}
func Marshal(data interface{}) (result []byte, err error) {
	typeInfo := reflect.TypeOf(data)
	if typeInfo.Kind() != reflect.Struct{
		err = errors.New("please pass struct")
		return
	}
	var dataSlice []string
	valueInfo := reflect.ValueOf(data)
	for i := 0; i < typeInfo.NumField();i++{
		field := typeInfo.Field(i)
		sectionVal := valueInfo.Field(i)
		fieldType := field.Type
		if fieldType.Kind() != reflect.Struct{
			continue
		}
		tagVal := field.Tag.Get("ini")
		if len(tagVal) == 0{
			tagVal = field.Name //是节
		}
		section := fmt.Sprintf("[%s]\n",tagVal)

		dataSlice = append(dataSlice,section)

		for j := 0;j < fieldType.NumField();j++{
			keyField := fieldType.Field(j)
			fieldTagVal := keyField.Tag.Get("ini")
			if len(fieldTagVal) == 0{
				fieldTagVal = keyField.Name
			}
			valField := sectionVal.Field(j)
			item := fmt.Sprintf("%s=%v\n",fieldTagVal,valField.Interface())

			//fmt.Println("The elem is ",item)
			dataSlice = append(dataSlice,item)
		}


	}
	for _,v := range dataSlice{
		result = append(result,[]byte(v)...)
		fmt.Println(string(result))
	}
	return

}

func parseItem(lastFileName string, line string, result interface{}) (err error) { //解析k  v 选项
	index := strings.Index(line, "=") //如果有= 号 则是正常的 k v 合法性判断
	//var string keyFiledname
	if index == -1 {
		err = fmt.Errorf("syntax error,line:%s lineno :%d", line, index+1)
		return
	}

	key := strings.TrimSpace(line[0:index])
	val := strings.TrimSpace(line[index+1:])
	_ = val

	if len(key) == 0 {
		err = fmt.Errorf("syntax error,line:%s", line)
		return
	}

	resultValue := reflect.ValueOf(result)
	sectionValue := resultValue.Elem().FieldByName(lastFileName)
	sectionType := sectionValue.Type()
	if sectionType.Kind() != reflect.Struct{
		err = fmt.Errorf("filed: %s must be struct",lastFileName)
		return

	}
	keyFieldName := ""
	for i := 0;i < sectionType.NumField();i ++{
		field := sectionType.Field(i)
		tagVal := field.Tag.Get("ini")
		if tagVal == key{
			keyFieldName = field.Name
			break
		}

	}
	if len(keyFieldName) == 0{
		return
	}

	fieldValue := sectionValue.FieldByName(keyFieldName)
	if fieldValue == reflect.ValueOf(nil){
		return
	}
	switch (fieldValue.Type().Kind()) {
	case reflect.String:
		fieldValue.SetString(val)
	case reflect.Int,reflect.Int8,reflect.Int16,reflect.Int32,reflect.Int64:
		intVal,errRet := strconv.ParseInt(val,10,64)
		if errRet != nil{
			err = errRet
			return
		}
		fieldValue.SetInt(intVal)
	case reflect.Uint,reflect.Uint8,reflect.Uint16,reflect.Uint32,reflect.Uint64:
		uintVal,errRet := strconv.ParseUint(val,10,64)
		if errRet != nil{
			err = errRet
			return
		}
		fieldValue.SetUint(uintVal)
	case reflect.Float32,reflect.Float64:
		fintVal,errRet := strconv.ParseFloat(val,64)
		if errRet != nil{
			err = errRet
			return
		}
		fieldValue.SetFloat(fintVal)
		fmt.Println(fieldValue)
	default:
		err = fmt.Errorf("unsupport type:%d\n",fieldValue.Type().Kind())


		
	}
	//fmt.Println(keyFieldName)
	//fmt.Println(sectionType)
	//fmt.Println(ty)
	//if ty != reflect.Struct{
	//	//err = fmt.Errorf("filed: %s must be struct",lastName)
	//	err = fmt.Errorf("filed :%s must be struct",lastName)
	//}
	//switch ty {
	//case reflect.Float64:
	//	err = fmt.Errorf("filed :%s must be struct",lastName)
	//case reflect.Struct:
	//	err = fmt.Errorf("filed :%s must be struct",lastName)
	//case reflect.Ptr:
	//	err = fmt.Errorf("filed :%s this is a ptr",lastName)
	//default:
	//	err = fmt.Errorf("this is a %s",lastName)
	//
	//}

	return
}

func parseSection(line string, typeInfo reflect.Type) (fileName string, err error) {//获取ini文件中的 节
	if line[0] == '[' && len(line) <= 2 {
		err = fmt.Errorf("syntax error,invalid section:%s", line)
		return
	}
	if line[0] == '[' && line[len(line)-1] != ']' {
		err = fmt.Errorf("syntax error,invalid section:%s", line)
		return
	}

	if line[0] == '[' && line[len(line)-1] == ']' {
		sectionName := strings.TrimSpace(line[1 : len(line)-1])  //节的名字
		if len(sectionName) == 0 {
			err = fmt.Errorf("syntax error,invalid section:%s", line)
			return

		}

		for i := 0; i < typeInfo.NumField(); i++ { //获取tag信息
			filed := typeInfo.Field(i)
			//fmt.Println("The config.ini filed is ",filed)
			//fmt.Println(filed.Tag.Get("ini"))
			tagValue := filed.Tag.Get("ini")
			if tagValue == sectionName {
				fileName = filed.Name
				fmt.Println(fileName)
				break
			}
		}
	}
	return
}

func UnMarshalFile(filename string,result interface{})(err error)  {
	data,err := ioutil.ReadFile(filename)
	if err != nil{
		return
	}
	return UnMarshal(data,result)
}
func UnMarshal(data []byte, result interface{}) (err error) {
	lineArray := strings.Split(string(data), "\n")   //按照行 换行分割，得到每一行，并且是全部

	typeInfo := reflect.TypeOf(result) //利用反射 得到类型
	if typeInfo.Kind() != reflect.Ptr {  //如果结果值不等于 Ptr 抛出异常
		err = errors.New("please pass address")
		return
	}
	typeInfoTwo := typeInfo.Elem() //获取指针变量指向的类型
	if typeInfoTwo.Kind() != reflect.Struct {
		err = errors.New("please pass struct")
		return
	}

	var lastFieldName string
	for index, line := range lineArray {
		var lastName string
		_ = lastName
		line = strings.TrimSpace(line) // 去除两边的空格
		if len(line) == 0 {
			continue
		}
		if line[0] == ';' || line[0] == '#' {//判断是不是注释
			continue
		}
		if line[0] == '[' {

			lastFieldName, err = parseSection(line, typeInfoTwo)
			if err != nil {
				err = fmt.Errorf("%v lineno:%d", err, index+1)
				return
			}
			continue

		}
		err = parseItem(lastFieldName, line, result)
		if err != nil {
			err = fmt.Errorf("%v lineno:%d", err, index+1)//格式化日志
			return
		}

	}
	return
}
