package helpers

import (
	"github.com/go-playground/locales/en"
	"net/mail"

	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
)

var (
	uni      *ut.UniversalTranslator
	validate *validator.Validate
)

func Validate() (*validator.Validate, ut.Translator) {
	translator := en.New()
	uni = ut.New(translator, translator)

	trans, _ := uni.GetTranslator("en")

	validate = validator.New()
	enTranslation := enTranslations.RegisterDefaultTranslations(validate, trans)

	if enTranslation != nil {
		panic(enTranslation)
	}

	return validate, trans
}

func EmailFormatValidation(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}
