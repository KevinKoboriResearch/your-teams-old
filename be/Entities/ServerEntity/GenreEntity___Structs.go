package GenreEntity

type GenreEntityRepository struct {
}

type GenreEntityController struct {
	GenreEntityRepository GenreEntityRepository
}

type Genre struct {
	Genre string `json:"genre" validate:"required,genre-exist"`
}

type SubGenre struct {
	SubGenre string `json:"subgenre" validate:"required,subgenre-exist"`
}

type GenreEntity struct {
	Genre string `json:"genre" validate:"required,genre-verify"`
	SubGenre string `json:"subgenre"`
}

type GenreEntities []GenreEntity
