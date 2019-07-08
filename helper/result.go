package helper

type Result struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func ReturnResult(code int, msg string, data interface{}) *Result {
	return &Result{code, msg, data}
}

func Assert(b bool, res *Result) {
	if !b {
		panic(res)
	}
}

func CheckErr(err error, res *Result) {
	if err != nil {
		Assert(false, res)
	}
}
