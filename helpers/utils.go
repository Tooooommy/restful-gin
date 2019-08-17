package helpers

import (
	"CrownDaisy_GOGIN/config"
	"crypto/md5"
	"github.com/satori/go.uuid"
	"io"
)

func UUID() string {
	return uuid.NewV4().String()
}

func Md5(part string) string {
	hash := md5.New()
	return string(hash.Sum([]byte(part)))
}

func GenPwd(password string) string {
	hash := md5.New()
	_, _ = io.WriteString(hash, config.Get().App.Secret)
	_, _ = io.WriteString(hash, password)
	return string(hash.Sum(nil))
}
