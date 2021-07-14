package validate

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	"github.com/go-playground/locales/zh_Hant_TW"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"reflect"
	"regexp"
	"strings"
	"time"
)

var (
	Uni      *ut.UniversalTranslator
	Validate *validator.Validate
)

//InitValidate 初始化
func InitValidate() {
	en := en.New()
	zh := zh.New()
	zh_tw := zh_Hant_TW.New()
	Uni = ut.New(en, zh, zh_tw)
	Validate = validator.New()
	Validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		return fld.Tag.Get("comment")
	})
	//自定义格式
	_ = Validate.RegisterValidation("date_format", dateFormat)
	_ = Validate.RegisterValidation("hour_format", hourFormat)
	_ = Validate.RegisterValidation("id_format", IDFormat)
}

//DefaultGetValidParams 检测函数
func DefaultGetValidParams(c *gin.Context, params interface{}) error {
	if err := c.ShouldBind(params); err != nil {
		return err
	}
	v := c.Value("trans")
	trans, ok := v.(ut.Translator)
	if !ok {
		trans, _ = Uni.GetTranslator("zh")
	}
	err := Validate.Struct(params)
	if err != nil {
		errs := err.(validator.ValidationErrors)
		sliceErrs := []string{}
		for _, e := range errs {
			sliceErrs = append(sliceErrs, e.Translate(trans))
		}
		return errors.New(strings.Join(sliceErrs, ","))
	}
	return nil
}
func ValidParams(params interface{}) error {
	trans, _ := Uni.GetTranslator("zh")
	err := Validate.Struct(params)
	if err != nil {
		errs := err.(validator.ValidationErrors)
		sliceErrs := []string{}
		for _, e := range errs {
			sliceErrs = append(sliceErrs, e.Translate(trans))
		}
		return errors.New(strings.Join(sliceErrs, ","))
	}
	return nil
}

//dateFormat 检查日期格式
func dateFormat(fl validator.FieldLevel) bool {
	var errStr string
	defer func() {
		if errStr != "" {
			fmt.Println("dateFormat err ",errStr)
		}
	}()
	switch fl.Field().Kind() {
	case reflect.String:
		str := fl.Field().String()
		_,err := time.Parse("2006-01-02", str)
		if err != nil {
			errStr = str
			return false
		}
	default:
		for i := 0; i < fl.Field().Len();i++ {
			str := fl.Field().Index(i).String()
			_,err := time.Parse("2006-01-02", str)
			if err != nil {
				errStr = str
				return false
			}
		}
	}
	return true
}
//hourFormat 检查时间格式
func hourFormat(fl validator.FieldLevel) bool {
	var errStr string
	defer func() {
		if errStr != "" {
			fmt.Println("hourFormat err ",errStr)
		}
	}()
	switch fl.Field().Kind() {
	case reflect.String:
		str := fl.Field().String()
		_,err := time.Parse("15:04", str)
		if err != nil {
			errStr = str
			return false
		}
	default:
		for i := 0; i < fl.Field().Len();i++ {
			str := fl.Field().Index(i).String()
			_,err := time.Parse("15:04", str)
			if err != nil {
				errStr = str
				return false
			}
		}
	}
	return true
}
//IDFormat id 检查是否含有" "
func IDFormat(fl validator.FieldLevel) bool {
	var errStr string
	defer func() {
		if errStr != "" {
			fmt.Println("id err ",errStr)
		}
	}()
	//对" "进行检查
	str := "[\\s]+"
	reg := regexp.MustCompile(str)
	switch fl.Field().Kind() {
	case reflect.String:
		str := fl.Field().String()
		arr := reg.FindAllString(str ,-1)
		if len(arr) != 0 {
			return false
		}
	default:
		for i := 0; i < fl.Field().Len(); i++ {
			str := fl.Field().Index(i).String()
			arr := reg.FindAllString(str ,-1)
			if len(arr) != 0 {
				return false
			}
		}
	}
	return true
}


