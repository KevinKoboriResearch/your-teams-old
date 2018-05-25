package UserEntity

import (
	"be/HyperText"
)

var controller = &UserEntityController{UserEntityRepository: UserEntityRepository{}}
var routes = HyperText.Routes{{}}

func UserEntityRoutes() HyperText.Routes {
	routes := HyperText.Routes{
		HyperText.Route{ //FUNCIONANDO-------------------------------------------------
			"SignUp - User Entity",
			"POST",
			"/SignUp",
			controller.SignUp,
		},
		HyperText.Route{ //FUNCIONANDO-------------------------------------------------
			"Login - User Entity",
			"POST",
			"/Login",
			controller.Login,
		},
		HyperText.Route{ //FUNCIONANDO-------------------------------------------------
			"Update Single - User Entity",
			"PUT",
			"/YourAccount/UpdateSingle/username={username}",
			controller.UpdateSingle,
		},
		HyperText.Route{ //FUNCIONANDO-------------------------------------------------
			"Update Partial - User Entity",
			"PUT",
			"/YourAccount/UpdatePartial/username={username}",
			controller.UpdatePartial,
		},
		HyperText.Route{ //FUNCIONANDO-------------------------------------------------
			"Update - User Entity",
			"PUT",
			"/YourAccount/Update/username={username}",
			controller.Update,
		},
		HyperText.Route{ //FUNCIONANDO-------------------------------------------------
			"Disable - User Entity",
			"PUT",
			"/YourAccount/Disable/username={username}",
			controller.Disable,
		},
		HyperText.Route{ //FUNCIONANDO-------------------------------------------------
			"Get Unit - User Entity",
			"GET",
			"/Search/User/username={username}",
			controller.GetUnit,
		},
		HyperText.Route{ //FUNCIONANDO-------------------------------------------------
			"Get All While - User Entities",
			"GET",
			"/Search/Users/position={position}&value={value}",
			controller.GetAllEnabledWhile,
		},
		HyperText.Route{ //FUNCIONANDO-------------------------------------------------
			"Get All Enabled - User Entities",
			"GET",
			"/Search/Users",
			controller.GetAllEnabled,
		},
		HyperText.Route{ //FUNCIONANDO-------------------------------------------------
			"Get All Existing - User Entities",
			"GET",
			"/Search/UsersAll",
			controller.GetAll,
		},
		HyperText.Route{ //FUNCIONANDO-------------------------------------------------
			"Delete - User Entity",
			"DELETE",
			"/YourAccount/Delete/username={username}",
			controller.Delete,
		},
	}
	return routes
}
