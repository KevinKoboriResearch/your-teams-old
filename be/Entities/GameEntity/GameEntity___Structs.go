package GameEntity

//import ("be/Entities/GenreEntity")

type GameEntityRepository struct{}

type GameEntityController struct {
	GameEntityRepository GameEntityRepository
}

type ByOne struct {
	Position string `json:"position" validate:"required"`
	Value    string `json:"value" validate:"required"`
}

type GameEntity struct {
	Name string `json:"name" validate:"required,name-exist"`
	Abbre string `json:"abbreviation"`
	Desc string `json:"desc"`
	Enable bool `json:"enable"`
//	Company CompanyEntity.CompanyEntity `json:"company"`
//	Genre GenreEntity.GenreEntity `json:"genre"`
//	Server     string `json:"server" validate:"required,password-length"`
}

type GameEntities []GameEntity

type GameExist struct {
	Name string `json:"name" validate:"required,name-exist"`
}

type AdminGameEntity struct {
	Username string `json:"username"`
	Password string `json:"password" validate:"required,password-length"`
	Name string `json:"name" validate:"required,name-exist"`
	Abbre string `json:"abbreviation"`
	Desc string `json:"desc"`
	Enable bool `json:"enable"`
//	Company CompanyEntity.CompanyEntity `json:"company"`
//	Genre GenreEntity.GenreEntity `json:"genre"`
//	Server     string `json:"server" validate:"required,password-length"`
}
