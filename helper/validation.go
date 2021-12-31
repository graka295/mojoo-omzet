package helper

import (
	"reflect"
	"strings"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"gopkg.in/go-playground/validator.v9"
	en_translations "gopkg.in/go-playground/validator.v9/translations/en"
)

// Validation services for validation
type Validation interface {
	Validate(data interface{}) ([]map[string]string, bool)
}

// ValidationImpl struct for validation services
type ValidationImpl struct {
	Vd    *validator.Validate
	Trans ut.Translator
}

// NewValidate init services
func NewValidate(vd *validator.Validate) Validation {
	en := en.New()
	uni := ut.New(en, en)
	trans, _ := uni.GetTranslator("en")
	en_translations.RegisterDefaultTranslations(vd, trans)
	vd.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})
	return &ValidationImpl{
		Vd:    vd,
		Trans: trans,
	}
}

// Validate func for validate struct
func (c *ValidationImpl) Validate(data interface{}) ([]map[string]string, bool) {
	err := c.Vd.Struct(data)
	if err != nil {
		errorMap := []map[string]string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMapSingale := map[string]string{}
			rightOfDelimiter := strings.Join(strings.Split(e.Namespace(), ".")[1:], ".")
			errorMapSingale["message"] = e.Translate(c.Trans)
			errorMapSingale["address"] = rightOfDelimiter
			errorMap = append(errorMap, errorMapSingale)
		}
		return errorMap, false
	}
	return []map[string]string{}, true
}
