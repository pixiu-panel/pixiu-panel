package validator

import (
	"errors"
	"gitee.ltd/lxh/logger/log"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
	"github.com/kataras/iris/v12"
	"strings"
)

var (
	uni      *ut.UniversalTranslator
	validate *validator.Validate
	trans    ut.Translator
)

// Init
// @description: 初始化验证器
func Init(app *iris.Application) {
	//注册翻译器
	zhTranslator := zh.New()
	uni = ut.New(zhTranslator, zhTranslator)

	trans, _ = uni.GetTranslator("zh")

	//获取gin的校验器
	validate = validator.New()
	//注册翻译器
	err := zhTranslations.RegisterDefaultTranslations(validate, trans)
	if err != nil {
		log.Panicf("注册翻译器失败：%v", err)
	}
	app.Validator = validate
}

// Translate
// @description: 翻译错误信息
// @param err
// @return error
func Translate(err error) error {
	errorMsg := make([]string, 0)

	ves, ok := err.(validator.ValidationErrors)
	if !ok {
		return err
	}

	for _, e := range ves {
		errorMsg = append(errorMsg, e.Translate(trans))
	}

	return errors.New(strings.Join(errorMsg, "；"))
}
