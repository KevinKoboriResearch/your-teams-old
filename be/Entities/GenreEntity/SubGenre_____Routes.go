package GenreEntity

import (
	"be/HyperText"
)

var controller = &SubGenreEntityController{SubGenreEntityRepository: SubGenreEntityRepository{}}
var routes = HyperText.Routes{{}}

func SubGenreRoutes() HyperText.Routes {
	routes := HyperText.Routes{
		HyperText.Route{
			"Register - SubGenre Entity",
			"POST",
			"/Admin/SubGenre/Insert",
			controller.Insert,
		},
		HyperText.Route{
			"Update - SubGenre Entity",
			"PUT",
			"/Admin/SubGenre/Update/SubGenre={SubGenre}",
			controller.Update,
		},
		HyperText.Route{
			"Get Unit - SubGenre Entity",
			"GET",
			"/Admin/SubGenre/Search/SubGenre={SubGenre}",
			controller.GetUnit,
		},
		HyperText.Route{
			"Get All Existing - SubGenre Entities",
			"GET",
			"/Admin/SubGenre/Search/SubGenre=all",
			controller.GetAll,
		},
		HyperText.Route{
			"Delete - SubGenre Entity",
			"DELETE",
			"/Admin/SubGenre/Delete/SubGenre={SubGenre}",
			controller.Delete,
		},
	}
	return routes
}
