package GameEntity

import (
	"be/HyperText"
)

const (
	DOCNAME = "game_entity"
	ID_NAME = "name"
)

var (
	controller = &GameEntityController{GameEntityRepository: GameEntityRepository{}}
	routes     = HyperText.Routes{{}}
)
