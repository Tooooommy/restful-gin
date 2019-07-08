package helper

import "github.com/tidwall/gjson"

func InitGson() gjson.Result {
	return gjson.Get("todo", ".")
}
