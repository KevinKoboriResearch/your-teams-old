package Interface

import ("gopkg.in/go-playground/validator.v9"
	"gopkg.in/mgo.v2/bson"
	"strings")

type NameExist struct {
  Name string `json:"name"`
}

var DOCNAME string

func PutDocValidateNameExist(doc string) {
    DOCNAME = doc
}

//-------------------------------- VAL --------------------------------//
func ValidateNameUsed(name validator.FieldLevel) bool {
	c := OpenSession(DOCNAME)
	n := strings.ToLower(name.Field().String())
	var result NameExist
	if err := c.Find(bson.M{"name": n}).One(&result); err != nil {
	}
	return result.Name == ""
}

//-------------------------------- VAL --------------------------------//
func ValidateNameExist(name validator.FieldLevel) bool {
	c := OpenSession(DOCNAME)
	n := strings.ToLower(name.Field().String())
	var result NameExist
	if err := c.Find(bson.M{"name": n}).One(&result); err != nil {
	}
	return result.Name != ""
}
