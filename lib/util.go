package utils

import (
	"github.com/levigross/grequests"
	"os"
)

func InitClient() {
	_, _ = grequests.Delete("www.baidu.com", nil)
}

func IsFileExist(filename string) bool {
	if _, err := os.Stat(filename); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}
