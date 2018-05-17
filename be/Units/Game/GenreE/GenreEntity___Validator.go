package GameEntity

import (
	"gopkg.in/go-playground/validator.v9"
	"gopkg.in/mgo.v2/bson"
	"strings"
	"yt/be/Database"
	"yt/be/HyperText"
)

//-------------------------------- VAL --------------------------------//
func ValidateGenreExist(genre validator.FieldLevel) bool {
	c := Database.OpenSession(DOCNAME)
	genre := strings.ToLower(genre.Field().String())
	result := GenreEntity{}
	c.Find(bson.M{"genre": genre}).One(&result)
	return result.Genre == ""
}

//-------------------------------- VAL --------------------------------//
func ValidateSubGenreExist(subgenre validator.FieldLevel) bool {
	c := Database.OpenSession(DOCNAME)
	subgenre := strings.ToLower(subgenre.Field().String())
	result := GenreEntity{}
	c.Find(bson.M{"subgenre": subgenre}).One(&result)
	return result.SubGenre == ""
}
