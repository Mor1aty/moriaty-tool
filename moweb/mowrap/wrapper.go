package mowrap

type Wrapper struct {
	Code uint16      `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}
