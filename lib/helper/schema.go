package helper

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type Schema interface {
}

// use a single instance of Validate, it caches struct info
var validate *validator.Validate

func validateSchema(sc Schema) {

	validate = validator.New()

	validateStruct(sc)
}

func validateStruct(sc Schema) {

	err := validate.Struct(sc)
	if err != nil {

		// this check is only needed when your code could produce
		// an invalid value for validation such as interface with nil
		// value most including myself do not usually have code like this.
		if _, ok := err.(*validator.InvalidValidationError); ok {
			fmt.Println(err)
			return
		}

		for _, err := range err.(validator.ValidationErrors) {

			fmt.Println(err.Namespace())
			fmt.Println(err.Field())
			fmt.Println(err.StructNamespace())
			fmt.Println(err.StructField())
			fmt.Println(err.Tag())
			fmt.Println(err.ActualTag())
			fmt.Println(err.Kind())
			fmt.Println(err.Type())
			fmt.Println(err.Value())
			fmt.Println(err.Param())
			fmt.Println()
		}

		// from here you can create your own error messages in whatever language you wish
		return
	}
}
