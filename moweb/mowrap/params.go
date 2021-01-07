package mowrap

import (
	"encoding/json"
	"log"
	"mime/multipart"
	"reflect"
	"strconv"
)

type Params struct {
	param map[string]interface{}
}

func NewWrapParams(param map[string]interface{}) *Params {
	return &Params{param: param}
}

func (wp Params) String() string {
	paramJson, err := json.Marshal(wp.param)
	if err != nil {
		log.Println(err)
		return ""
	}
	return string(paramJson)
}

func (wp Params) Json() {

}

// 存入字段
func (wp Params) Put(key string, value interface{}) {
	wp.param[key] = value
}

// 获取参数数量
func (wp Params) Size() int {
	return len(wp.param)
}

// 判断参数是否为空
func (wp Params) IsEmpty() bool {
	return len(wp.param) == 0
}

// 判断参数是否包含 key
func (wp Params) ContainsKey(key string) bool {
	_, ok := wp.param[key]
	return ok
}

// 判断参数是否包含 key，且 val 必为 string 的情况下不为空串
func (wp Params) ContainsStringKey(key string) bool {
	val, ok := wp.param[key]
	if !ok {
		return false
	}
	if val.(string) == "" {
		return false
	}
	return true
}

// 获取 key 获取对应 value
func (wp Params) Get(key string) interface{} {
	return wp.param[key]
}

// 根据 key 获取对应 value，转化为 boolean
func (wp Params) GetBoolValue(key string) bool {
	return wp.param[key].(bool)
}

// 根据 key 获取对应 value，转化为 int
func (wp Params) GetIntValue(key string) int {
	var v int
	if reflect.TypeOf(wp.param[key]).Kind() == reflect.String {
		v, _ = strconv.Atoi(wp.param[key].(string))
	} else {
		v = int(wp.param[key].(float64))
	}
	return v
}

// 根据 key 获取对应 value，转化为 float64
func (wp Params) GetFloatValue(key string) float64 {
	var v float64
	if reflect.TypeOf(wp.param[key]).Kind() == reflect.String {
		v, _ = strconv.ParseFloat(wp.param[key].(string), 64)
	} else {
		v = wp.param[key].(float64)
	}
	return v
}

// 根据 key 获取对应 value，转化为 string
func (wp Params) GetString(key string) string {
	// TODO string 需要优化
	return wp.param[key].(string)
}

// 根据 key 获取对应 value，转化为 File
func (wp Params) GetFileValue(key string) *multipart.FileHeader {
	return wp.param[key].(*multipart.FileHeader)
}

// 清除所有参数
func (wp Params) Clear() {
	wp.param = make(map[string]interface{})
}

// 删除 key 对应的参数
func (wp Params) remove(key string) {
	delete(wp.param, key)
}

// 获取参数 key 集合
func (wp Params) KeySet() []string {
	keys := make([]string, 0, len(wp.param))
	for key := range wp.param {
		keys = append(keys, key)
	}
	return keys
}
