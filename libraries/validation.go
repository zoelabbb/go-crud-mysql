package libraries

import (
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

type Validation struct {
	validate *validator.Validate
	trans    ut.Translator
}

func NewValidation() *Validation {
	translator := en.New()

	// Universal translator
	uni := ut.New(translator, translator)
	trans, _ := uni.GetTranslator("en")

	validate := validator.New()
	en_translations.RegisterDefaultTranslations(validate, trans)

	return &Validation{
		validate: validate,
		trans:    trans,
	}
}

func (v *Validation) Struct(s interface{}) interface{} {
	errors := make(map[string]string)

	//
	err := v.validate.Struct(s)
	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			errors[e.StructField()] = e.Translate(v.trans)
		}
	}

	if len(errors) > 0 {
		return errors
	}

	return nil
}
