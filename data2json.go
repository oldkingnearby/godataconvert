package godataconvert

// 数据转换
import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

// struct 转为 map 并把字段小写
func Struct2Map(obj interface{}) map[string]interface{} {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	var data = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		data[strings.ToLower(t.Field(i).Name)] = v.Field(i).Interface()
	}
	return data
}

// map 转为json
func Map2Json(m map[string]interface{}) []byte {

	jsonStr, err := json.Marshal(m)

	if err != nil {
		log.Fatal(err)
	}
	return jsonStr
}

// json 转为 map
func Json2Map(jsonStr []byte) map[string]interface{} {

	var mapResult map[string]interface{}
	//使用 json.Unmarshal(data []byte, v interface{})进行转换,返回 error 信息
	if err := json.Unmarshal([]byte(jsonStr), &mapResult); err != nil {
		log.Fatal(err)
	}
	return mapResult
}

// struct 转 json
func Struct2Json(obj interface{}) []byte {
	jsonBytes, err := json.Marshal(obj)
	if err != nil {
		log.Fatal(err)
	}
	return jsonBytes
}

// map 转 bson
func Map2Bson(m map[string]interface{}) bson.M {
	retBson := bson.M{}
	for k, v := range m {
		retBson[k] = v
	}
	return retBson
}
