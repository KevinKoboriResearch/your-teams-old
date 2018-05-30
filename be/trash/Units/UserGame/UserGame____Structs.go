package UserGame

//import ("be/Entities/GenreEntity")

type UserGameRepository struct{}

type UserGameController struct {
	UserGameRepository UserGameRepository
}

type ByOne struct {
	Position string `json:"position" validate:"required"`
	Value    string `json:"value" validate:"required"`
}

type UserGame struct {
	Username string `json:"username" validate:"required"`
	Name     string `json:"name" validate:"required"`
	Abbre    string `json:"abbre"`
	UserDesc string `json:"userdesc"`
	Enable   bool   `json:"enable"`
	//	StarsHistory CompanyEntity.CompanyEntity `json:"company"`
}

type UserGames []UserGame

type GameUsed struct {
	Name string `json:"name" validate:"required"`
}

type UserGameUpdateSingle struct {
	Username string `json:"username"`
	Password string `json:"password" validate:"required,password-length"`
	Name     string `json:"name" validate:"required,name-exist`
	Position string `json:"position" validate:"required"`
	Value    string `json:"value" validate:"required"`
}
