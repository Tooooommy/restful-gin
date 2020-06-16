package types

import (
	"restful-gin/helpers/define"
)

// 存放一些类型的数据
// 例如参数和返回值
// 以_type.go结尾类似

type Result struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func ReturnResult(code int, msg string, data interface{}) *Result {
	return &Result{code, msg, data}
}

func Assert(bo bool, code int, message string, ds ...interface{}) {
	if !bo {
		var data interface{}
		if len(ds) > 0 {
			data = ds[0]
		}
		panic(ReturnResult(code, message, data))
	}
}

func CheckErr(err error, cs ...int) {
	if err != nil {
		var code = define.Failed
		if len(cs) > 0 {
			code = cs[0]
		}
		Assert(false, code, err.Error())
	}
}
