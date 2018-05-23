package UserEntity

import (
	"gopkg.in/go-playground/validator.v9"
	"gopkg.in/mgo.v2/bson"
	"strings"
	"be/Interface"
	"be/HyperText"
)

//-------------------------------- VAL --------------------------------//
func ValidateUsernameUsed(username validator.FieldLevel) bool {
	c := Interface.OpenSession(DOCNAME)
	u := strings.ToLower(username.Field().String())
	result := UserEntity{}
	if err := c.Find(bson.M{"username": u}).One(&result); err != nil {
	}
	return result.Username == ""
}

//-------------------------------- VAL --------------------------------//
func ValidateEmailUsed(email validator.FieldLevel) bool {
	c := Interface.OpenSession(DOCNAME)
	e := strings.ToLower(email.Field().String())
	result := UserEntity{}
	if err := c.Find(bson.M{"email": e}).One(&result); err != nil {
	}
	return result.Email == ""
}

//-------------------------------- VAL --------------------------------//
func ValidateUsernameLength(username validator.FieldLevel) bool {
	length := len(username.Field().String())
	return length >= 5
}

//-------------------------------- VAL --------------------------------//
func ValidatePasswordLength(password validator.FieldLevel) bool {
	length := len(password.Field().String())
	return length >= 6
}

//-------------------------------- VAL --------------------------------//
func ValidateUpdateSingle(ueus UserEntityUpdateSingle) interface{} {
	var err interface{}
	if ueus.Position == "email" {
		e := UserEmail{}
		e.Email = ueus.Value
		err = HyperText.StructValidate(&e)
	} else if ueus.Position == "username" {
		u := UserUsername{}
		u.Username = ueus.Value
		err = HyperText.StructValidate(&u)
	}
	return err
}
