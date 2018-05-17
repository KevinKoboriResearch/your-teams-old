package GenreEntity

import (
	"yt/be/HyperText"
)

var controller = &GenreEntityController{GenreEntityRepository: GenreEntityRepository{}}
var routes = HyperText.Routes{{}}

func GenreRoutes() HyperText.Routes {
	routes := HyperText.Routes{
		HyperText.Route{
			"Register - Genre Entity",
			"POST",
			"/Admin/Genre/Insert",
			controller.Insert,
		},
		HyperText.Route{
			"Update - Genre Entity",
			"PUT",
			"/Admin/Genre/Update/genre={genre}",
			controller.Update,
		},
		HyperText.Route{
			"Get Unit - Genre Entity",
			"GET",
			"/Admin/Genre/Search/genre={genre}",
			controller.GetUnit,
		},
		HyperText.Route{
			"Get All Existing - Genre Entities",
			"GET",
			"/Admin/Genre/Search/genre=all",
			controller.GetAll,
		},
		HyperText.Route{
			"Delete - Genre Entity",
			"DELETE",
			"/Admin/Genre/Delete/genre={genre}",
			controller.Delete,
		},
	}
	return routes
}
