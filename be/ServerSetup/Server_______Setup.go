package ServerSetup

import (
	//"be/Entities/GameEntity"
	"be/Entities/UserEntity"
	"be/HyperText"
	"be/Interface"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"log"
	"net/http"
)

const (
	SERVER_IP   = "localhost"
	SERVER_PORT = ":8080"
	SERVER_HOST = SERVER_IP + SERVER_PORT
)

func StartServer() {
	Interface.StartConectionDatabase()
	HyperText.StartValidator()
	StartValidatorUserEntity()
	StartValidatorInterfaceEntity()
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "DELETE", "PUT"},
		AllowCredentials: true,
	})
	router := CreateAllRoutes()
	handler := c.Handler(router)
	log.Print("\n		 					Starting myServer...\n\n")
	if err := http.ListenAndServe(SERVER_HOST, handler); err != nil {
		panic(err)
	}
}

func CreateAllRoutes() (routes *mux.Router) {
	userEntityRoutes := UserEntity.UserEntityRoutes()
	//	gameEntityRoutes := GameEntity.GameEntityRoutes()
	appRoutes := append(userEntityRoutes) //, gameEntityRoutes...)
	//	appRoutes := append(appRoutes, ...)
	routes = HyperText.NewRouter(appRoutes)
	return routes
}

func StartValidatorUserEntity() {
	HyperText.Validate.RegisterValidation("username-used", UserEntity.ValidateUsernameUsed)
	HyperText.Validate.RegisterValidation("username-exist", UserEntity.ValidateUsernameExist)
	HyperText.Validate.RegisterValidation("username-length", UserEntity.ValidateUsernameLength)
	HyperText.Validate.RegisterValidation("password-length", UserEntity.ValidatePasswordLength)
	HyperText.Validate.RegisterValidation("email-used", UserEntity.ValidateEmailUsed)
}

func StartValidatorInterfaceEntity() {
	HyperText.Validate.RegisterValidation("name-used", Interface.ValidateNameUsed)
	HyperText.Validate.RegisterValidation("name-exist", Interface.ValidateNameExist)
}
