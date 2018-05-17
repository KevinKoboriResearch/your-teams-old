package GameEntity

import (
	"yt/be/HyperText"
)

var controller = &GameEntityController{GameEntityRepository: GameEntityRepository{}}
var routes = HyperText.Routes{{}}

func GameEntityRoutes() HyperText.Routes {
	routes := HyperText.Routes{
		HyperText.Route{
			"Register - Game Entity",
			"POST",
			"/Admin/Game/Insert",
			controller.Insert,
		},
		HyperText.Route{
			"Update - Game Entity",
			"PUT",
			"/Admin/Game/Update/game={game}",
			controller.Update,
		},
		HyperText.Route{
			"Get Unit - Game Entity",
			"GET",
			"/Admin/Game/Search/game={game}",
			controller.GetUnit,
		},
		HyperText.Route{
			"Get All While - Game Entities",
			"GET",
			"/Admin/Game/Search/position={position}&value={value}",
			controller.GetAllWhile,
		},
		HyperText.Route{
			"Get All Existing - Game Entities",
			"GET",
			"/Admin/Game/Search/game=all",
			controller.GetAll,
		},
		HyperText.Route{
			"Delete - Game Entity",
			"DELETE",
			"/Admin/Game/Delete/game={game}",
			controller.Delete,
		},
	}
	return routes
}
