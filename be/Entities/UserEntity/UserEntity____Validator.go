package UserEntity

import (
	"be/HyperText"
	"be/Interface"
	"gopkg.in/go-playground/validator.v9"
	"gopkg.in/mgo.v2/bson"
	"strings"
)

//__ USERNAME USED ___________________________________________________________//
func ValidateUsernameUsed(username validator.FieldLevel) bool {
	c := Interface.OpenSession(DOCNAME)
	u := strings.ToLower(username.Field().String())
	result := UserEntity{}
	if err := c.Find(bson.M{"username": u}).One(&result); err != nil {
	}
	return result.Username == ""
}

//__ USERNAME EXIST __________________________________________________________//
func ValidateUsernameExist(username validator.FieldLevel) bool {
	c := Interface.OpenSession(DOCNAME)
	u := strings.ToLower(username.Field().String())
	result := UserEntity{}
	if err := c.Find(bson.M{"username": u}).One(&result); err != nil {
	}
	return result.Username != ""
}

//__ USERNAME LENGTH__________________________________________________________//
func ValidateUsernameLength(username validator.FieldLevel) bool {
	length := len(username.Field().String())
	return length >= 5
}

//__ PASSWORD LENGTH _________________________________________________________//
func ValidatePasswordLength(password validator.FieldLevel) bool {
	length := len(password.Field().String())
	return length >= 8
}

func ValidatePasswordMatch(password validator.FieldLevel) bool {
	c := Interface.OpenSession(DOCNAME)
	u := password.Field().String()
	result := UserEntity{}
	if err := c.Find(bson.M{"username": u}).One(&result); err != nil {
	}
	return result.Username != ""
}

//__ EMAIL USED ______________________________________________________________//
func ValidateEmailUsed(email validator.FieldLevel) bool {
	c := Interface.OpenSession(DOCNAME)
	e := strings.ToLower(email.Field().String())
	result := UserEntity{}
	if err := c.Find(bson.M{"email": e}).One(&result); err != nil {
	}
	return result.Email == ""
}

//__ UPDATE SINGLE ___________________________________________________________//
func ValidateUpdateSingle(c *UserEntityController, ueus UserEntityUpdateSingle) (err interface{}) {
	if ueus.Position == "username" {
		var u UserUsername
		u.Username = ueus.Value
		err = HyperText.StructValidate(&u)
		return
	}
	if ueus.Position == "email" {
		var e UserEmail
		e.Email = ueus.Value
		err = HyperText.StructValidate(&e)
		return
	}
	if ueus.Position == "password" {
		var p UserPassword
		p.Password = ueus.Value
		err = HyperText.StructValidate(&p)
		if err == nil {
			var uev UserEntityVerify
			uev.Username = ueus.Username
			uev.Password = ueus.Password
			err = c.UserEntityRepository.VerifyUserEntity(c, uev)
			return
		}
		return
	}
	return
}

//__ UPDATE __________________________________________________________________//
func ValidateUpdate(ueu UserEntityUpdate) interface{} {
	if ueu.Email != "" {
		var e UserEmail
		e.Email = ueu.Email
		err := HyperText.StructValidate(e)
		if err != nil {
			return err
		}
	}
	if ueu.Image != "" {
		var i UserImage
		i.Image = ueu.Image
		err := HyperText.StructValidate(i)
		if err != nil {
			return err
		}
	}
	return nil
}
