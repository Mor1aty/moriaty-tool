package mowrap

import "github.com/Mor1aty/moriaty-tool/moweb/moconstant"

// 成功封装
func OK(code uint16, msg string, data interface{}) *Wrapper {
	return &Wrapper{
		Code: code,
		Msg:  msg,
		Data: data,
	}
}

// 执行成功封装
func OKExec(msg string, data ...interface{}) *Wrapper {
	if len(data) == 0 {
		return &Wrapper{
			Code: moconstant.CodeExecSuccess,
			Msg:  msg,
			Data: nil,
		}
	}
	return &Wrapper{
		Code: moconstant.CodeExecSuccess,
		Msg:  msg,
		Data: data[0],
	}
}

// 获取成功封装
func OKObtain(msg string, data interface{}) *Wrapper {
	return &Wrapper{
		Code: moconstant.CodeObtainSuccess,
		Msg:  msg,
		Data: data,
	}
}

// 错误封装
func Error(code uint16, msg string) *Wrapper {
	return &Wrapper{
		Code: code,
		Msg:  msg,
		Data: nil,
	}
}

// 执行错误封装
func ErrorExec(msg string) *Wrapper {
	return &Wrapper{
		Code: moconstant.CodeExecFailure,
		Msg:  msg,
		Data: nil,
	}
}

// 获取错误封装
func ErrorObtain(msg string) *Wrapper {
	return &Wrapper{
		Code: moconstant.CodeObtainFailure,
		Msg:  msg,
		Data: nil,
	}
}
