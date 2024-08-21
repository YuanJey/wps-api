/*
** description("").
** copyright('tuoyun,www.tuoyun.net').
** author("fg,Gordon@tuoyun.net").
** time(2021/4/8 15:09).
 */
package utils

import (
	"encoding/json"
	"github.com/jinzhu/copier"
	"math/rand/v2"
	"strconv"
	"time"
)

func OperationIDGenerator() string {
	return strconv.FormatInt(time.Now().UnixNano()+int64(rand.Uint32()), 10)
}
func IntToString(i int) string {
	return strconv.FormatInt(int64(i), 10)
}

func StringToInt(i string) int {
	j, _ := strconv.Atoi(i)
	return j
}
func StringToInt64(i string) int64 {
	j, _ := strconv.ParseInt(i, 10, 64)
	return j
}
func StringToInt32(i string) int32 {
	j, _ := strconv.ParseInt(i, 10, 64)
	return int32(j)
}
func Int32ToString(i int32) string {
	return strconv.FormatInt(int64(i), 10)
}

func Uint32ToString(i uint32) string {
	return strconv.FormatInt(int64(i), 10)
}

// judge a string whether in the  string list
func IsContain(target string, List []string) bool {
	for _, element := range List {

		if target == element {
			return true
		}
	}
	return false
}
func IsContainInt32(target int32, List []int32) bool {
	for _, element := range List {
		if target == element {
			return true
		}
	}
	return false
}
func IsContainInt(target int, List []int) bool {
	for _, element := range List {
		if target == element {
			return true
		}
	}
	return false
}

func StructToJsonBytes(param interface{}) []byte {
	dataType, _ := json.Marshal(param)
	return dataType
}

func int64ToString(i int64) string {
	return strconv.FormatInt(i, 10)
}
func Int64ToString(i int64) string {
	return strconv.FormatInt(i, 10)
}

func RemoveDuplicateElement(idList []string) []string {
	result := make([]string, 0, len(idList))
	temp := map[string]struct{}{}
	for _, item := range idList {
		if _, ok := temp[item]; !ok {
			temp[item] = struct{}{}
			result = append(result, item)
		}
	}
	return result
}

// copy a by b  b->a
func CopyStructFields(a interface{}, b interface{}, fields ...string) (err error) {
	return copier.Copy(a, b)
}
