package lib

import "github.com/levigross/grequests"

func InitClient() {
	_, _ = grequests.Delete("www.baidu.com", nil)
}
