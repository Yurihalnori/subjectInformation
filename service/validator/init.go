package validator

import (
	"fmt"
	"reflect"
	"subjectInformation/global"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
)

// var validatorRouter = map[string]validator.Func{
// 	"timing":   timing,
// 	"category": category,
// }

// func InitValidator() {
// 	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
// 		for name, function := range validatorRouter {
// 			err := v.RegisterValidation(name, function)
// 			if err != nil {
// 				panic(err)
// 			}
// 		}
// 	}
// }

// InitTrans 初始化翻译器
func InitValidator(locale string) (err error) {
	// 修改gin框架中的Validator引擎属性，实现自定制
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {

		zhT := zh.New() // 中文翻译器
		enT := en.New() // 英文翻译器

		// 第一个参数是备用（fallback）的语言环境
		// 后面的参数是应该支持的语言环境（支持多个）
		// uni := ut.New(zhT, zhT) 也是可以的
		uni := ut.New(enT, zhT, enT)

		// locale 通常取决于 http 请求头的 'Accept-Language'
		var ok bool
		// 也可以使用 uni.FindTranslator(...) 传入多个locale进行查找
		global.Trans, ok = uni.GetTranslator(locale)
		if !ok {
			return fmt.Errorf("uni.GetTranslator(%s) failed", locale)
		}

		// 注册翻译器
		switch locale {
		case "zh":
			err = zhTranslations.RegisterDefaultTranslations(v, global.Trans)
			v.RegisterTagNameFunc(func(fld reflect.StructField) string {
				return fld.Tag.Get("comment")
			})

			//自定义验证方法
			//category
			v.RegisterValidation("category", category)
			v.RegisterTranslation("category", global.Trans, func(ut ut.Translator) error {
				return ut.Add("category", "{0}输入格式或长度不符(十位二进制)", true)
			},
				func(ut ut.Translator, fe validator.FieldError) string {
					t, _ := ut.T("category", fe.Field())
					return t
				})

			//timing
			v.RegisterValidation("timing", timing)
			v.RegisterTranslation("timing", global.Trans, func(ut ut.Translator) error {
				return ut.Add("timing", "{0}输入时间不合理", true)
			},
				func(ut ut.Translator, fe validator.FieldError) string {
					t, _ := ut.T("timing", fe.Field())
					return t
				})

		case "en":
			err = enTranslations.RegisterDefaultTranslations(v, global.Trans)
		default:
			err = enTranslations.RegisterDefaultTranslations(v, global.Trans)
		}
		return
	}
	return
}
