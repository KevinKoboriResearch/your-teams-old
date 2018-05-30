package GameEntity

//import ("be/Entities/GenreEntity")

type GameEntityRepository struct{}

type GameEntityController struct {
	GameEntityRepository GameEntityRepository
}

type GameEntity struct {
	Name   string `json:"name" validate:"required,name-used"`
	Abbre  string `json:"abbre"`
	Desc   string `json:"desc"`
	Enable bool   `json:"enable"`
	//	Company CompanyEntity.CompanyEntity `json:"company"`
	//	Genre GenreEntity.GenreEntity `json:"genre"`
	//	Server    ServerEntity.ServerEntity `json:"server" validate:"required,password-length"`
}

type GameEntities []GameEntity

type GameUsed struct {
	Name string `json:"name" validate:"required,name-used"`
}

type AdminGameEntity struct {
	Username string `json:"username"`
	Password string `json:"password" validate:"required,password-length"`
	Name     string `json:"name" validate:"required,name-used"`
	Abbre    string `json:"abbre"`
	Desc     string `json:"desc"`
	Enable   bool   `json:"enable"`
	//	Company CompanyEntity.CompanyEntity `json:"company"`
	//	Genre GenreEntity.GenreEntity `json:"genre"`
	//	Server     string `json:"server" validate:"required,password-length"`
}

type AdminGameEntityUpdateSingle struct {
	Username string `json:"username"`
	Password string `json:"password" validate:"required,password-length"`
	Name     string `json:"name" validate:"required,name-exist`
	Position string `json:"position" validate:"required"`
	Value    string `json:"value" validate:"required"`
}
