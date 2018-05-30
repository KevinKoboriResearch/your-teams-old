package UserGame

import (
	"be/HyperText"
)

var controller = &GameEntityController{GameEntityRepository: GameEntityRepository{}}
var routes = HyperText.Routes{{}}

func GameEntityRoutes() HyperText.Routes {
	routes := HyperText.Routes{
		HyperText.Route{ //FUNCIONANDO-------------------------------------------------
			"Register - Game Entity",
			"POST",
			"/Game/Insert",
			controller.Insert,
		},
		HyperText.Route{ //FUNCIONANDO-------------------------------------------------
			"Update Single - Game Entity",
			"PUT",
			"/UpdateSingle",
			controller.UpdateSingle,
		},
		HyperText.Route{ //FUNCIONANDO-------------------------------------------------
			"Update - Game Entity",
			"PUT",
			"/Game/Update/name={name}",
			controller.Update,
		},
		HyperText.Route{ //FUNCIONANDO-------------------------------------------------
			"Get Unit - Game Entity",
			"GET",
			"/Game/Search/name={name}",
			controller.GetUnit,
		},
		HyperText.Route{ //FUNCIONANDO-------------------------------------------------
			"Get All While - Game Entities",
			"GET",
			"/Game/Search/position={position}&value={value}",
			controller.GetAllWhile,
		},
		HyperText.Route{ //FUNCIONANDO-------------------------------------------------
			"Get All Existing - Game Entities",
			"GET",
			"/Game/Search/All",
			controller.GetAll,
		},
		HyperText.Route{ //FUNCIONANDO-------------------------------------------------
			"Delete - Game Entity",
			"DELETE",
			"/Game/Delete/name={name}",
			controller.Delete,
		},
	}
	return routes
}
