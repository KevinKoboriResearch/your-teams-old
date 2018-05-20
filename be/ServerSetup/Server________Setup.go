package ServerSetup

import (
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"log"
	"net/http"
	"be/Database"
	"be/HyperText"
	"be/Entities/UserEntity"
)

const (
	SERVER_IP   = "localhost"
	SERVER_PORT = ":8080"
	SERVER_HOST = SERVER_IP + SERVER_PORT
)

func StartServer() {
	Database.StartConectionDatabase()
	HyperText.StartValidator()
	StartValidatorUserEntity()
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
	userRoutes := UserEntity.UserEntityRoutes()
	appRoutes := append(userRoutes, userRoutes...)
	routes = HyperText.NewRouter(appRoutes)
	return routes
}

func StartValidatorUserEntity() {
	HyperText.Validate.RegisterValidation("username-used", UserEntity.ValidateUsernameUsed)
	HyperText.Validate.RegisterValidation("email-used", UserEntity.ValidateEmailUsed)
	HyperText.Validate.RegisterValidation("username-length", UserEntity.ValidateUsernameLength)
	HyperText.Validate.RegisterValidation("password-length", UserEntity.ValidatePasswordLength)
}
