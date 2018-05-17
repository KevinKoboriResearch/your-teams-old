package Genre

import (
	"github.com/gorilla/mux"
	"net/http"
	"strings"
	"yt/be/Database"
	"yt/be/HyperText"
)

//-------------------------------- ADM --------------------------------//
func (c *GenreController) Insert(w http.ResponseWriter, r *http.Request) {
	var ge Genre
	if err := HyperText.BodyValidate(r, &ge); err != nil {
		HyperText.HttpErrorResponse(w, http.StatusBadRequest, err)
		return
	}
	ge.Genre = strings.ToLower(ge.Genre)
	if err := Database.InsertGenreEntity(&ge); err != nil {
		HyperText.HttpErrorResponse(w, http.StatusBadRequest, err)
		return
	}
	HyperText.HttpResponse(w, http.StatusOK, ge)
	return
}

//-------------------------------- ADM --------------------------------//
func (c *GenreController) Update(w http.ResponseWriter, r *http.Request) {
	ueu := GenreUpdate{}
	ueu.Username = mux.Vars(r)["username"]
	if err := HyperText.BodyValidate(r, &ueu); err != nil {
		HyperText.HttpErrorResponse(w, http.StatusBadRequest, err)
		return
	}
	uev := GenreVerify{ueu.Username, ueu.Password}
	if _, boolean := c.GenreRepository.VerifyGenre(uev); boolean == false {
		HyperText.HttpErrorResponse(w, http.StatusBadRequest, HyperText.CustomResponses["wrong-Login"])
		return
	}
	if result, err := Database.UpdateDB(ueu, DOCNAME, ueu.Username); err != nil {
		HyperText.HttpErrorResponse(w, http.StatusBadRequest, HyperText.CustomResponses["error-update"])
		return
	}
	HyperText.HttpResponse(w, http.StatusOK, result)
	return
}

//-------------------------------- ADM --------------------------------//
func (c *GenreController) GetUnit(w http.ResponseWriter, r *http.Request) {
	username := mux.Vars(r)["username"]
	username = strings.ToLower(username)
	result := Genre{}
	if 	err := Database.FindUnitDB(username, &result, DOCNAME); err != nil {
		HyperText.HttpErrorResponse(w, http.StatusBadRequest, HyperText.CustomResponses["not-found-"])
		return
	}
	HyperText.HttpResponse(w, http.StatusOK, result)
	return
}

//-------------------------------- ADM --------------------------------//
func (c *GenreController) GetAllWhile(w http.ResponseWriter, r *http.Request) {
	position := mux.Vars(r)["position"]
	value := mux.Vars(r)["value"]
	position = strings.ToLower(position)
	value = strings.ToLower(value)
	entities := GenreEntities{}
	if 	err := Database.FindAllWhileDB(position, value, &entities, DOCNAME); err != nil {
		HyperText.HttpErrorResponse(w, http.StatusBadRequest, HyperText.CustomResponses["empty-database"])
		return
	}
	HyperText.HttpResponse(w, http.StatusOK, entities)
	return
}

//-------------------------------- ADM --------------------------------//
func (c *GenreController) GetAll(w http.ResponseWriter, r *http.Request) {
	result := UserEntities{}
	if 	err := Database.FindAllDB(&result, DOCNAME); err != nil {
		HyperText.HttpErrorResponse(w, http.StatusBadRequest, HyperText.CustomResponses["empty-database"])
		return
	}
	HyperText.HttpResponse(w, http.StatusOK, result)
	return
}

//-------------------------------- ADM --------------------------------//
func (c *GenreController) Delete(w http.ResponseWriter, r *http.Request) {
	Genre := strings.ToLower(mux.Vars(r)["Genre"])
	result := Genre{}
	if 	err := Database.DeleteDB(Genre, &result, DOCNAME); err != nil {
		HyperText.HttpErrorResponse(w, http.StatusBadRequest, HyperText.CustomResponses["not-found-"])
		return
	}
	HyperText.HttpResponse(w, http.StatusOK, result)
	return
}
