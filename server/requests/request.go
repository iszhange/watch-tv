package requests

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

// 自定义正则表达式验证函数
func regexpValidation(fl validator.FieldLevel) bool {
	fieldValue := fl.Field().String()
	regex := fl.Param()

	matched, err := regexp.MatchString(regex, fieldValue)
	if err != nil {
		return false
	}

	return matched
}
