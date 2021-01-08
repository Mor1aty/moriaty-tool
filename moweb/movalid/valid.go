package movalid

import (
	"encoding/json"
	"fmt"
	"github.com/Mor1aty/moriaty-tool/moweb/moconstant"
	"github.com/Mor1aty/moriaty-tool/moweb/mowrap"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)

// 参数校验
func MiddleWare(params []Param) gin.HandlerFunc {
	return func(c *gin.Context) {
		paramMap := make(map[string]interface{})
		if c.Request.Method == "GET" {
			for _, param := range params {
				value, ok := c.GetQuery(param.Name)
				if !ok {
					c.Abort()
					c.JSON(http.StatusOK, mowrap.Error(moconstant.CodeParamFailure, param.Name+moconstant.MsgParamNull))
					return
				}
				res := handleParam(value, param.Methods)
				if res != "" {
					c.Abort()
					c.JSON(http.StatusOK, mowrap.Error(moconstant.CodeParamFailure, param.Name+res))
					return
				}
				paramMap[param.Name] = value
			}

		} else {
			if strings.Contains(c.Request.Header.Get("Content-Type"), "application/json") {
				body, err := ioutil.ReadAll(c.Request.Body)
				if err != nil {
					c.Abort()
					c.JSON(http.StatusOK, mowrap.Error(moconstant.CodeServerException, "服务器异常"))
					return
				}
				err = json.Unmarshal(body, &paramMap)
				if err != nil {
					c.Abort()
					c.JSON(http.StatusOK, mowrap.Error(moconstant.CodeServerException, "服务器异常"))
					return
				}
				for _, param := range params {
					_, ok := paramMap[param.Name]
					if !ok {
						c.Abort()
						c.JSON(http.StatusOK, mowrap.Error(moconstant.CodeParamFailure, param.Name+moconstant.MsgParamNull))
						return
					}
					res := handleParam(paramMap[param.Name], param.Methods)
					if res != "" {
						c.Abort()
						c.JSON(http.StatusOK, mowrap.Error(moconstant.CodeParamFailure, param.Name+res))
						return
					}
				}
			} else {
				for _, param := range params {
					isFile := handleFile(param.Methods)
					if isFile {
						file, err := c.FormFile(param.Name)
						if err != nil {
							if "http: no such file" == err.Error() {
								c.Abort()
								c.JSON(http.StatusOK, mowrap.Error(moconstant.CodeParamFailure, param.Name+moconstant.MsgParamNull))
								return
							}
							c.Abort()
							c.JSON(http.StatusOK, mowrap.Error(moconstant.CodeServerException, "服务器异常"))
							return
						}
						paramMap[param.Name] = file
					} else {
						value, ok := c.GetPostForm(param.Name)
						if !ok {
							c.Abort()
							c.JSON(http.StatusOK, mowrap.Error(moconstant.CodeParamFailure, param.Name+moconstant.MsgParamNull))
							return
						}
						res := handleParam(value, param.Methods)
						if res != "" {
							c.Abort()
							c.JSON(http.StatusOK, mowrap.Error(moconstant.CodeParamFailure, param.Name+res))
							return
						}
						paramMap[param.Name] = value
					}
				}
			}
		}
		if c.Keys == nil {
			c.Keys = make(map[string]interface{})
		}
		c.Keys["params"] = mowrap.NewWrapParams(paramMap)
		c.Next()
	}
}

// 获取参数
func GetWrapParams(c *gin.Context) *mowrap.Params {
	return c.Keys["params"].(*mowrap.Params)

}

// 处理参数
func handleParam(value interface{}, methods []Method) string {
	for _, method := range methods {
		switch method {
		case EMAIL:
			if compile := regexp.MustCompile(moconstant.RegularMail); !compile.MatchString(fmt.Sprintf("%v", value)) {
				return moconstant.MsgParamEmail
			}
		case NUMBER:
			if compile := regexp.MustCompile(moconstant.RegularNumber); !compile.MatchString(fmt.Sprintf("%v", value)) {
				return moconstant.MsgParamNumber
			}
		case PHONE:
			if compile := regexp.MustCompile(moconstant.RegularPhone); !compile.MatchString(fmt.Sprintf("%v", value)) {
				return moconstant.MsgParamPhone
			}
		}
	}
	return ""
}

// 判断参数是否为文件
func handleFile(methods []Method) bool {
	if len(methods) == 1 && methods[0] == FILE {
		return true
	}
	return false
}
