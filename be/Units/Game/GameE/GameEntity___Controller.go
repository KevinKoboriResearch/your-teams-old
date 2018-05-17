package GameEntity

import (
	"github.com/gorilla/mux"
	"net/http"
	"strings"
	"yt/be/Database"
	"yt/be/HyperText"
)

//-------------------------------- ADM --------------------------------//
func (c *GameEntityController) Insert(w http.ResponseWriter, r *http.Request) {
	var ge GameEntity
	if err := HyperText.BodyValidate(r, &ge); err != nil {
		HyperText.HttpErrorResponse(w, http.StatusBadRequest, err)
		return
	}
	ge.Game = strings.ToLower(ge.Game)
	ge.Server = strings.ToLower(ge.Server)
	if err := Database.InsertDB(&ge, DOCNAME); err != nil {
		HyperText.HttpErrorResponse(w, http.StatusBadRequest, err)
		return
	}
	HyperText.HttpResponse(w, http.StatusOK, ge)
	return
}

//-------------------------------- ADM --------------------------------//
func (c *GameEntityController) Update(w http.ResponseWriter, r *http.Request) {
	ueu := GameEntityUpdate{}
	ueu.Username = mux.Vars(r)["username"]
	if err := HyperText.BodyValidate(r, &ueu); err != nil {
		HyperText.HttpErrorResponse(w, http.StatusBadRequest, err)
		return
	}
	uev := GameEntityVerify{ueu.Username, ueu.Password}
	if _, boolean := c.GameEntityRepository.VerifyGameEntity(uev); boolean == false {
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
func (c *GameEntityController) GetUnit(w http.ResponseWriter, r *http.Request) {
	username := mux.Vars(r)["username"]
	username = strings.ToLower(username)
	result := GameEntity{}
	if 	err := Database.FindUnitDB(username, &result, DOCNAME); err != nil {
		HyperText.HttpErrorResponse(w, http.StatusBadRequest, HyperText.CustomResponses["not-found-entity"])
		return
	}
	HyperText.HttpResponse(w, http.StatusOK, result)
	return
}

//-------------------------------- ADM --------------------------------//
func (c *GameEntityController) GetAllWhile(w http.ResponseWriter, r *http.Request) {
	position := mux.Vars(r)["position"]
	value := mux.Vars(r)["value"]
	position = strings.ToLower(position)
	value = strings.ToLower(value)
	entities := GameEntities{}
	if 	err := Database.FindAllWhileDB(position, value, &entities, DOCNAME); err != nil {
		HyperText.HttpErrorResponse(w, http.StatusBadRequest, HyperText.CustomResponses["empty-database"])
		return
	}
	HyperText.HttpResponse(w, http.StatusOK, entities)
	return
}

//-------------------------------- ADM --------------------------------//
func (c *GameEntityController) GetAll(w http.ResponseWriter, r *http.Request) {
	result := UserEntities{}
	if 	err := Database.FindAllDB(&result, DOCNAME); err != nil {
		HyperText.HttpErrorResponse(w, http.StatusBadRequest, HyperText.CustomResponses["empty-database"])
		return
	}
	HyperText.HttpResponse(w, http.StatusOK, result)
	return
}

//-------------------------------- ADM --------------------------------//
func (c *GameEntityController) Delete(w http.ResponseWriter, r *http.Request) {
	game := strings.ToLower(mux.Vars(r)["game"])
	result := GameEntity{}
	if 	err := Database.DeleteDB(game, &result, DOCNAME); err != nil {
		HyperText.HttpErrorResponse(w, http.StatusBadRequest, HyperText.CustomResponses["not-found-entity"])
		return
	}
	HyperText.HttpResponse(w, http.StatusOK, result)
	return
}
