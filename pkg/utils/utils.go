package utils

import (
	"encoding/json"
	"github.com/pkg/errors"
	"reflect"
	"runtime"
	"strconv"
)

func StructToJsonString(param interface{}) string {
	dataType, _ := json.Marshal(param)
	dataString := string(dataType)
	return dataString
}

func StructToJsonStringDefault(param interface{}) string {
	if reflect.TypeOf(param).Kind() == reflect.Slice && reflect.ValueOf(param).Len() == 0 {
		return "[]"
	}
	return StructToJsonString(param)
}

// The incoming parameter must be a pointer
func JsonStringToStruct(s string, args interface{}) error {
	return Wrap(json.Unmarshal([]byte(s), args), "json Unmarshal failed")
}
func Wrap(err error, message string) error {
	return errors.Wrap(err, "==> "+printCallerNameAndLine()+message)
}
func printCallerNameAndLine() string {
	pc, _, line, _ := runtime.Caller(2)
	return runtime.FuncForPC(pc).Name() + "()@" + strconv.Itoa(line) + ": "
}
func InterfaceArrayToSIntArray(data []interface{}) (i []int) {
	for _, param := range data {
		i = append(i, int(param.(float64)))
	}
	return i
}
func InterfaceArrayToStringArray(data []interface{}) (i []string) {
	for _, param := range data {
		i = append(i, param.(string))
	}
	return i
}
