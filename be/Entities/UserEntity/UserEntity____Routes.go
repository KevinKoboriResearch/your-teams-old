package UserEntity

import (
	"be/HyperText"
)

//__ ROUTES __________________________________________________________________//
func UserEntityRoutes() HyperText.Routes {
	routes := HyperText.Routes{
		HyperText.Route{ //___________________ FUNCIONANDO _______________________//
			"SignUp - User Entity",
			"POST",
			"/SignUp",
			controller.SignUp,
		},
		HyperText.Route{ //___________________ FUNCIONANDO _______________________//
			"Login - User Entity",
			"POST",
			"/Login",
			controller.Login,
		},
		HyperText.Route{ //___________________ FUNCIONANDO _______________________//
			"Update Single - User Entity",
			"PUT",
			"/YourAccount/UpdateSingle/username={username}",
			controller.UpdateSingle,
		},
		HyperText.Route{ //___________________ FUNCIONANDO _______________________//
			"Update - User Entity",
			"PUT",
			"/YourAccount/Update/username={username}",
			controller.Update,
		},
		HyperText.Route{ //___________________ FUNCIONANDO _______________________//
			"Disable - User Entity",
			"PUT",
			"/YourAccount/Disable/username={username}",
			controller.Disable,
		},
		HyperText.Route{ //___________________ FUNCIONANDO _______________________//
			"Get Entity - User Entity",
			"GET",
			"/Search/User/username={username}",
			controller.GetUnique,
		},
		HyperText.Route{ //___________________ FUNCIONANDO _______________________//
			"Get All While - User Entities",
			"GET",
			"/Search/Users/position={position}&value={value}",
			controller.GetAllEnabledWhile,
		},
		HyperText.Route{ //___________________ FUNCIONANDO _______________________//
			"Get All Enabled - User Entities",
			"GET",
			"/Search/Users",
			controller.GetAllEnabled,
		},
		HyperText.Route{ //___________________ FUNCIONANDO _______________________//
			"Get All - User Entities",
			"GET",
			"/Search/UsersAll",
			controller.GetAll,
		},
		HyperText.Route{ //___________________ FUNCIONANDO _______________________//
			"Delete - User Entity",
			"DELETE",
			"/YourAccount/Delete/username={username}",
			controller.Delete,
		},
	}
	return routes
}
