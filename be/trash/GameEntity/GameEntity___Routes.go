package GameEntity

import (
	"be/HyperText"
)

func GameEntityRoutes() HyperText.Routes {
	routes := HyperText.Routes{
		HyperText.Route{ //FUNCIONANDO-------------------------------------------------
			"Register - Game Entity",
			"POST",
			"/Admin/Game/Insert",
			controller.Insert,
		},
		HyperText.Route{ //FUNCIONANDO-------------------------------------------------
			"Update Single - Game Entity",
			"PUT",
			"/Admin/UpdateSingle",
			controller.UpdateSingle,
		},
		HyperText.Route{ //FUNCIONANDO-------------------------------------------------
			"Update - Game Entity",
			"PUT",
			"/Admin/Game/Update/name={name}",
			controller.Update,
		},
		HyperText.Route{ //FUNCIONANDO-------------------------------------------------
			"Get Unit - Game Entity",
			"GET",
			"/Admin/Game/Search/name={name}",
			controller.Get,
		},
		HyperText.Route{ //FUNCIONANDO-------------------------------------------------
			"Get All While - Game Entities",
			"GET",
			"/Admin/Game/Search/position={position}&value={value}",
			controller.GetAllWhile,
		},
		HyperText.Route{ //FUNCIONANDO-------------------------------------------------
			"Get All Existing - Game Entities",
			"GET",
			"/Admin/Game/Search/All",
			controller.GetAll,
		},
		HyperText.Route{ //FUNCIONANDO-------------------------------------------------
			"Delete - Game Entity",
			"DELETE",
			"/Admin/Game/Delete/name={name}",
			controller.Delete,
		},
	}
	return routes
}
