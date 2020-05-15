package iniconfig

import (
	"errors"
	"fmt"
	"io/ioutil"
	"reflect"
	"strconv"
	"strings"
)

// 声名一个函数将用户传入的结构体数据经过序列化后存入配置文件
func MarshalFile(filename string, data interface{}) (err error) {
	result, er := Marshal(data)
	if er != nil {
		return
	}
	// 将内容存入文件操作
	return ioutil.WriteFile(filename, result, 0755)
}

func Marshal(data interface{}) (result []byte, err error) {
	/*
		声名一个 Marshal 方法将用户传入的结构体序列化为 ini 格式的数据方法，
			该方法接收一个空接口数据，因为我们不知道用户会传什么样的结构体进来，
			该方法返回一个序列化之后的 byte 类型的数组数据，和一个序列化失败信息；
	*/
	// 通过 reflect 包获取用户传进来数据的类型信息
	typeInfo := reflect.TypeOf(data)
	valueInfo := reflect.ValueOf(data) // 获取用户传入数据中的值类型信息
	// 判断用户传过来的类型数据是否为结构体，不是直接抛出异常
	if typeInfo.Kind() != reflect.Struct {
		err = errors.New("please pass struct")
		return
	}

	var dataiInfo []string
	// 通过循环获取反射结构体中的数据
	for i := 0; i < typeInfo.NumField(); i++ {
		sectionField := typeInfo.Field(i) // 获取映射结构体中的每个字段
		sectionValue := valueInfo.Field(i)
		fieldType := sectionField.Type
		if fieldType.Kind() != reflect.Struct {
			continue
		}
		tagVal := sectionField.Tag.Get("ini")
		if len(tagVal) == 0 { // 如果没有定义 tag 信息则直接使用反射结构体中的字段名
			tagVal = sectionField.Name
		}
		section := fmt.Sprintf("\n[%s]\n", tagVal)
		dataiInfo = append(dataiInfo, section)

		// 获取结构体中的字段为结构体的字段名
		for j := 0; j < fieldType.NumField(); j++ {
			keyField := fieldType.Field(j)         // 获取反射结构体字段中的字段
			fieldTagVal := keyField.Tag.Get("ini") // 获取反射结构体中字段的 tag 信息
			if len(fieldTagVal) == 0 {
				fieldTagVal = keyField.Name
			}

			valField := sectionValue.Field(j) // 获取反射结构体中字段的名称
			// valField.Interface() 获取反射结构体中字段名中存储的值内容
			item := fmt.Sprintf("%s=%v\n", fieldTagVal, valField.Interface())
			dataiInfo = append(dataiInfo, item)
		}
	}

	for _, val := range dataiInfo {
		// 将字符串类型的数据转换为字节数组
		byteVal := []byte(val)
		result = append(result, byteVal...)
	}
	return
}

// 名一个函数将用户传入的配置文件经过序列化后转换为结构体类型数据
func UnMarshalFile(filename string, result interface{}) (err error) {
	// 读取文件, 返回一个 byte 数组类型的数据，和一个错误
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}
	// 调用 UnMarshal 方法将配置文件内容序列化为结构体
	return UnMarshal(data, result)
}

func UnMarshal(data []byte, result interface{}) (err error) {
	/*
			声名一个 UnMarshal 反序列化方法，用来将用户传入的配置文件数据反序列化为用户需要的数据类型，
				该方法接收两个参数：第一个参数为 byte 类型的数组数据，以及一个空接口用来确定用户需要将传进来的数据反序列化为什么样的数据类型，
				该方法仅返回一个序列化出错的错误信息给用户；
		    该方法由于需要对 result 中的结构体数据做修改，所以需要传入一个指针类型的数据
	*/
	// 将用户传进来的数据，通过 \n 进行分割即按行分割，得到一个字符串数组
	linesArr := strings.Split(string(data), "\n")

	// 获取用户传过来的数据结构类型信息
	typeInfo := reflect.TypeOf(result)

	// 通过 reflect 方法判断 result 空结构体中用户传入的内容是否为指针，不是则返回一个错误信息，
	// 因为如果不是指针是不能进行解析并返回给用户的；
	if typeInfo.Kind() != reflect.Ptr {
		// 创建一个错误信息
		err = errors.New("please pass address")
		return
	}

	// 通过 Elem 方法获取指针中用户传递的变量类型
	typeStruct := typeInfo.Elem()
	// 通过 Kind 判断用户传递的变量类型是否为结构体,不是则返回一个错误信息
	if typeStruct.Kind() != reflect.Struct {
		err = errors.New("please pass struct")
		return
	}

	// 用来保存当前节点岁对应反射结构体中结构体的名称，用作后面解析该节点下的内容
	var lastFieldName string
	// 循环遍历用户传入的每一个数据，开始序列化用户传入的数据
	for i, line := range linesArr {
		// 去掉每一行中前后的空格内容
		line = strings.TrimSpace(line)
		if len(line) == 0 { // 判断是否为空行
			continue
		}
		// 判断每一行的第一个字符是否为 ; 和 #  是则表示为注释行
		if line[0] == ';' || line[0] == '#' {
			continue
		}

		if line[0] == '[' {
			// 调用节点解析函数，传入每行的数据和结构体的变量类型
			lastFieldName, err = parseSection(line, typeStruct)
			if err != nil {
				err = fmt.Errorf("%v, lineNo:%d", err, i+1)
				return
			}
			continue
		}

		// 调用选项解析方法，传入最后保存的节点名称、行内容以及用户传入分结构体信息
		err = parseItem(lastFieldName, line, result)
		if err != nil {
			err = fmt.Errorf("%v, lineNo:%d", err, i+1)
			return
		}
	}
	return
}

// 声名一个 parseItem 函数用来解析配置文件中节点下的选项内容
func parseItem(lastFieldName, line string, result interface{}) (err error) {
	/* strings.Index(str, char) 用来查找一个字符转中 = 号所在的位置，返回一个位置信息
	接收两个参数：str、第一个为需要查找的字符串内容
				 char、第二个为需要查找的字符
	*/
	index := strings.Index(line, "=")
	if index == -1 { // 如果没有找到 = 号对应的位置信息,则说明配置文件内容错误
		err = fmt.Errorf("config file syntax error, line:%s", line)
		return
	}

	// 使用切片的方式获取每一行中 = 号左右的内容，左边的设置为 key 右边的设置为 value，并去除切出来字符前后的多余空格
	key := strings.TrimSpace(line[0:index])
	val := strings.TrimSpace(line[index+1:])

	//去掉空格之后如果 key 的长度为 0 表示为空内容抛出异常
	if len(key) == 0 {
		err = fmt.Errorf("config file syntax error, line:%s", line)
		return
	}
	// 获取用户传入指针类型结构体中的值类型
	resultValue := reflect.ValueOf(result)
	// 通过值类型对象获取反射结构体中的字段内容对象
	sectionValue := resultValue.Elem().FieldByName(lastFieldName)
	sectionType := sectionValue.Type() // 通过值类型对象获取结构体类型对象
	// 判断如果获取到的字段内容不属于结构体则抛出异常
	if sectionType.Kind() != reflect.Struct {
		err = fmt.Errorf("field %s must be struct", lastFieldName)
		return
	}

	keyFieldaName := ""
	// 遍历反射结构体中的所有字段，获取结构体中的字段名
	for i := 0; i < sectionType.NumField(); i++ {
		filed := sectionType.Field(i)
		tagVal := filed.Tag.Get("ini")
		if tagVal == key {
			keyFieldaName = filed.Name
			break
		}
	}

	// 如果从映射中没有找到字段内容则直接退出执行
	if len(keyFieldaName) == 0 {
		return
	}
	// 通过值类型对象获取反射结构体中字段的名称
	fieldValue := sectionValue.FieldByName(keyFieldaName)
	// 如果没有获取到反射结构体中的字段内容则直接结束循环
	if fieldValue == reflect.ValueOf(nil) {
		return
	}

	// 判断获取到的字段值类型进行断言并设置值
	switch fieldValue.Type().Kind() {
	case reflect.String:
		fieldValue.SetString(val)
	// 多行处理整数类型断言语法
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		intVal, errInt := strconv.ParseInt(val, 10, 64)
		if errInt != nil {
			err = errInt
			return
		}
		fieldValue.SetInt(intVal)
	// 多行处理无符号整数类型断言语法
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		uintVal, errUint := strconv.ParseUint(val, 10, 64)
		if errUint != nil {
			err = errUint
			return
		}
		fieldValue.SetUint(uintVal)
	// 多行处理浮点数数、类型断言语法
	case reflect.Float32, reflect.Float64:
		floatVal, errFloat := strconv.ParseFloat(val, 64)
		if errFloat != nil {
			err = errFloat
			return
		}
		fieldValue.SetFloat(floatVal)
	default:
		err = fmt.Errorf("unsuport type:%v", fieldValue.Type().Kind())
	}
	return
}

// 声名一个函数用来解析用户传入配置文件数据中的节点信息
func parseSection(line string, typeStruct reflect.Type) (fieldName string, err error) {
	// 判断用户传入的配置内容中有非法内容，如 [、[] 等不完整的节点
	if line[0] == '[' && len(line) <= 2 {
		err = fmt.Errorf("syntax error, invalid section:%s", line)
		return
	}
	// 判断用户传入的配置内容中有非法内容，如 [server 等不完整的节点
	if line[0] == '[' && line[len(line)-1] != ']' {
		err = fmt.Errorf("syntax error, invalid section:%s", line)
		return
	}

	// 确定数据为一个 [节点] 内容
	if line[0] == '[' && line[len(line)-1] == ']' {
		// 获取中括号节点中间的内容
		sectionName := strings.TrimSpace(line[1 : len(line)-1])
		// 判断用户传入的配置内容中有非法内容，如 [	] 等不完整的节点
		if len(sectionName) == 0 {
			err = fmt.Errorf("syntax error, invalid section:%s", line)
			return
		}

		// 通过前面获取到的反射结构体中变量类型调用 NumField() 方法可获得反射中共有多少字段，
		// 通过循环的方法结合 Field() 方法便可以获得每个字段的名称
		for i := 0; i < typeStruct.NumField(); i++ {
			// 获取每个字段的名字(即反射的结构体名)
			field := typeStruct.Field(i)
			// 通过 Tag 获取结构体中每个字段所对应的 tag 信息
			tagValue := field.Tag.Get("ini")
			// 判断从反射结构体获取到 tag 字段信息与用户传入的文件中的节点信息是否相等，是则说明找到了我们需要的信息
			if tagValue == sectionName {
				fieldName = field.Name
				break
			}
		}
	}
	return
}
