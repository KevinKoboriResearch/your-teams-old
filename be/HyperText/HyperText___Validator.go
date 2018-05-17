package HyperText

import (
	"gopkg.in/go-playground/validator.v9"
	"log"
	"net/http"
)

var Validate *validator.Validate

func StartValidator() {
	Validate = validator.New()
}

func BodyValidate(r *http.Request, entity interface{}) interface{} {
	if err := DecodeJson(r.Body, entity); err != nil {
		log.Println("[ERROR] Can't decode json: ", err)
		return CustomResponses["wrong-json"]
	}
	if err := StructValidate(entity); err != nil {
		log.Println("[ERROR] Can't validate struct: ", err)
		return err
	}
	return nil
}

func StructValidate(entity interface{}) interface{} {
	err := Validate.Struct(entity)
	if err != nil {
		log.Println("[ERROR] Can't validate struct: ", err)
		var fields []string
		for _, value := range err.(validator.ValidationErrors) {
			fields = append(fields, value.Tag()+": "+value.Field())
		}
		return fields[len(fields)-1]
	}
	return nil
}
