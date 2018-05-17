package GenreEntity

import (
	"strings"
	"yt/be/Database"
)

const DOCNAME = "genre_entity"

//-------------------------------- USR --------------------------------//
func (r GenreEntityRepository) InsertGenreEntity(ge *GenreEntity) (err error) {
	ge.Genre = strings.ToLower(ge.Genre)
	err = Database.InsertDB(&ge, DOCNAME)
	return
}

//-------------------------------- USR --------------------------------//
func (r GenreEntityRepository) VerifyGenreEntity(ge GenreEntityVerify) (GenreEntityUpdate, bool) {
	ge := GenreEntity{}
	err := Database.FindOneDB(ge.Genre, &ge, DOCNAME)
	if err != nil {
		return ge, false
	} else if boolean := CheckPasswordHash(ge.Password, ge.Password); boolean == true {
		return ge, true
	}
	return ge, false
}
