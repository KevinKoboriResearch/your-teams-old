package config

import (
	"gopkg.in/go-playground/validator.v9"
	"log"
	"net/http"
)

var Validate *validator.Validate

func StartValidator() {
	Validate = validator.New()
}

func BodyValidate(r *http.Request, entity interface{}) error {
	if err := DecodeJson(r.Body, entity); err != nil {
		log.Println("[ERROR] Can't decode json: ", err)
		return err
	}
	if err := Validate.Struct(entity); err != nil {
		log.Println("[ERROR] Can't validate struct: ", err)
		return err
	}
	return nil
}
