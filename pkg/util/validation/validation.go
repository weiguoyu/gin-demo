package validation

import (
	"errors"
	"fmt"
	"github.com/go-playground/locales/zh"
	"github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	ZhTranslations "github.com/go-playground/validator/v10/translations/zh"
	"strings"
)

var (
	Validate *validator.Validate
	Trans    ut.Translator
)

func init() {
	ZhCn := zh.New()
	Uni := ut.New(ZhCn)
	Trans, _ = Uni.GetTranslator("zh")
	//验证器
	Validate = validator.New()
	//验证器注册翻译器
	ZhTranslations.RegisterDefaultTranslations(Validate, Trans)
}

func ValidateStruct(input interface{}) error {
	if err := Validate.Struct(input); err != nil {
		if validatorErr, ok := err.(*validator.InvalidValidationError); ok {
			return fmt.Errorf("%v", validatorErr)
		}
		errs := make([]string, 0)
		for _, err := range err.(validator.ValidationErrors) {
			errs = append(errs, err.Translate(Trans))
		}
		return errors.New(strings.Join(errs, ";"))
	}
	return nil
}
