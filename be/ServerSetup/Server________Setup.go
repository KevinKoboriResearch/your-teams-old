package ServerSetup

import (
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"log"
	"net/http"
	"be/Interface"
	"be/HyperText"
	"be/Entities/UserEntity"
		"be/Entities/GameEntity"
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
	gameEntityRoutes := GameEntity.GameEntityRoutes()
	appRoutes := append(userEntityRoutes, gameEntityRoutes...)
//	appRoutes := append(appRoutes, ...)
	routes = HyperText.NewRouter(appRoutes)
	return routes
}

func StartValidatorUserEntity() {
	HyperText.Validate.RegisterValidation("username-used", UserEntity.ValidateUsernameUsed)
	HyperText.Validate.RegisterValidation("email-used", UserEntity.ValidateEmailUsed)
	HyperText.Validate.RegisterValidation("username-length", UserEntity.ValidateUsernameLength)
	HyperText.Validate.RegisterValidation("password-length", UserEntity.ValidatePasswordLength)
}

func StartValidatorInterfaceEntity() {
	HyperText.Validate.RegisterValidation("name-exist", Interface.ValidateNameExist)
}
