package GameEntity

import (
	"gopkg.in/go-playground/validator.v9"
	"gopkg.in/mgo.v2/bson"
	"strings"
	"be/Interface"
//	"be/Entities/GenreEntity"
	//"be/HyperText"
)

//-------------------------------- VAL --------------------------------//
func ValidateGameExist(name validator.FieldLevel) bool {
	c := Interface.OpenSession(DOCNAME)
	n := strings.ToLower(name.Field().String())
	result := GameEntity{}
	c.Find(bson.M{"name": n}).One(&result)
	return result.Name == ""
}
/*
//-------------------------------- VAL --------------------------------//
func ValidateUpdateSingle(geus GameEntityUpdateSingle) interface{} {
	var err interface{}
	if geus.Position == "game" {
		game := GameExist{}
		game.Game = geus.Value
		err = HyperText.StructValidate(&game)
	} else if geus.Position == "genres" {
		genre := GenreEntity.GenreEntities{}
		genre.Genre = geus.Value
		err = HyperText.StructValidate(&genre)
	}
	return err
}
*/
