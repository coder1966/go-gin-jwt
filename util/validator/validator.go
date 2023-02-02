package validator

// 公共的验证工具
import (
	"fmt"
	"github.com/go-playground/locales/zh_Hans_CN"
	universalTranslator "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	translationsZh "github.com/go-playground/validator/v10/translations/zh"
	"go-gin-jwt/util/errmsg"
	"reflect"
)

// Validate 数据验证的方法
func Validate(data interface{}) (string, int) {
	validate := validator.New()                                     // 实例化验证
	universalTranslate := universalTranslator.New(zh_Hans_CN.New()) // 实例化翻译
	trans, _ := universalTranslate.GetTranslator("zh_Hans_CN")      // 翻译方法引进来

	err := translationsZh.RegisterDefaultTranslations(validate, trans) // 注册默认的翻译
	if err != nil {
		fmt.Println("err", err)
	}
	// 映射标签。通过反射，把User 结构体定义的标签反出来。
	validate.RegisterTagNameFunc(func(field reflect.StructField) string {
		return field.Tag.Get("label")
	})

	// 传进来的data进行验证
	err = validate.Struct(data) // 用验证结构体这个方法
	if err != nil {
		for _, v := range err.(validator.ValidationErrors) {
			return v.Translate(trans), errmsg.ERROR
		}
	}
	return "", errmsg.SUCCESS
}
