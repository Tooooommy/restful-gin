package helpers

type Result struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func ReturnResult(code int, msg string, data interface{}) *Result {
	return &Result{code, msg, data}
}

func Assert(bo bool, res *Result) {
	if !bo {
		panic(res)
	}
}

func CheckErr(err error, res *Result) {
	if err != nil {
		Assert(false, res)
	}
}
