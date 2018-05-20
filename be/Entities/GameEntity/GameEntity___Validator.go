package GameEntity

import (
	"gopkg.in/go-playground/validator.v9"
	"gopkg.in/mgo.v2/bson"
	"strings"
	"be/Database"
	"be/Units/Game/GenreEntity"
	"be/HyperText"
)

//-------------------------------- VAL --------------------------------//
func ValidateGameExist(game validator.FieldLevel) bool {
	c := Database.OpenSession(DOCNAME)
	game := strings.ToLower(game.Field().String())
	result := GameEntity{}
	c.Find(bson.M{"game": game}).One(&result)
	return result.Game == ""
}

//-------------------------------- VAL --------------------------------//
func ValidateGenreVerify(genre validator.FieldLevel) bool {
	c := Database.OpenSession(GenreEntity.DOCNAME)
	genre := strings.ToLower(genre.Field().String())
	result := GenreEntity.GenreEntity{}
	c.Find(bson.M{"genre": genre}).One(&result)
	return result.Genre == ""
}

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
