package GameEntity

import ("be/Units/Game/GenreEntity")

type GameEntityRepository struct{}

type GameEntityController struct {
	GameEntityRepository GameEntityRepository
}

type ByOne struct {
	Position string `json:"position" validate:"required"`
	Value    string `json:"value" validate:"required"`
}

type GameEntity struct {
	Game string `json:"game" validate:"required,game-exist"`
//	Company CompanyEntity.CompanyEntity `json:"company"`
//	Genre GenreEntity.GenreEntity `json:"genre"`
	Server     string `json:"server" validate:"required,password-length"`
	Enable   bool   `json:"enable"`
}

type GameEntities []GameEntity

type GameExist struct {
	Game string `json:"game" validate:"required,game-exist"`
}
