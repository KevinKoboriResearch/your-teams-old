package GameEntity

import (
//	"be/Entities/GenreEntity"
	"be/HyperText"
)

//-------------------------------- VAL --------------------------------//
func ValidateUpdateSingle(ageus AdminGameEntityUpdateSingle) interface{} {
	var err interface{}
	if ageus.Position == "name" {
		game := GameUsed{}
		game.Name = ageus.Value
		err = HyperText.StructValidate(&game)
	} /*else if geus.Position == "genres" {
		genre := GenreEntity.GenreEntities{}
		genre.Genre = geus.Value
		err = HyperText.StructValidate(&genre)
	}*/
	return err
}
