package validate

import "github.com/asaskevich/govalidator"

func InitValidate() string {
	return govalidator.Email
}
