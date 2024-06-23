package utils

import (
	"fmt"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslation "github.com/go-playground/validator/v10/translations/en"
)

type GenericResponse struct {
	Status           string                 `json:"status,omitempty"`
	Message          string                 `json:"message"`
	DeveloperMessage string                 `json:"developerMessage,omitempty"`
	Data             []string               `json:"data,omitempty"`
	UpdatedAt        string                 `json:"updatedAt,omitempty"`
	Error            string                 `json:"error,omitempty"`
	Code             int                    `json:"code,omitempty"`
	Request          map[string]interface{} `json:"request,omitempty"`
	Results          interface{}            `json:"results,omitempty"`
}

type ErrorHandler struct {
	DevMessage string                 `json:"developerMessage,omitempty"`
	Request    map[string]interface{} `json:"request,omitempty"`
	Response   map[string]interface{} `json:"response,omitempty"`
	UserId     string                 `json:"userId,omitempty"`
	Message    string                 `json:"message,omitempty"`
	Method     string                 `json:"method,omitempty"`
	Code       int                    `json:"code,omitempty"`
	Status     string                 `json:"status,omitempty"`
}

const (

	//Server Response
	INTERNAL_SERVER_ERROR = "Internal Server Error"
	SOMETHING_WENT_WRONG  = "Unable to process your request, please try again later"
	DATA_NOT_FOUND        = "Data Not Found"
	INVALID_REQUEST       = "Invalid Request"
)

func Validator(request interface{}) *ErrorHandler {
	validate := validator.New()
	english := en.New()
	uni := ut.New(english, english)
	trans, _ := uni.GetTranslator("en")
	_ = enTranslation.RegisterDefaultTranslations(validate, trans)
	var message string
	if err := validate.Struct(request); err != nil {
		validatorErrs := err.(validator.ValidationErrors)
		for index, e := range validatorErrs {
			len := len(validatorErrs)
			translatedErr := fmt.Errorf(e.Translate(trans))
			if index != len-1 {
				message += fmt.Sprint(translatedErr, ", ")
				continue
			}
			message += fmt.Sprint(translatedErr, ".")
		}
		return &ErrorHandler{Message: message}
	}
	return nil
}
