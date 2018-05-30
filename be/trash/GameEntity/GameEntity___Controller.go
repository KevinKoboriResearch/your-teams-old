package GameEntity

import (
	"be/Entities/UserEntity"
	"be/HyperText"
	"be/Interface"
	"github.com/gorilla/mux"
	"net/http"
	"strings"
)

//-------------------------------- ADM --------------------------------//
func (c *GameEntityController) Insert(w http.ResponseWriter, r *http.Request) {
	var age AdminGameEntity
	Interface.PutDOCNAME(DOCNAME)
	if err := HyperText.BodyValidate(r, &age); err != nil {
		HyperText.HttpErrorResponse(w, http.StatusBadRequest, err)
		return
	}
	uev := UserEntity.UserEntityVerify{}
	uev.Username = age.Username
	uev.Password = age.Password
	ue, boolean := UserEntity.AdminVerify(uev)
	if boolean == false {
		HyperText.HttpErrorResponse(w, http.StatusBadRequest, HyperText.CustomResponses["wrong-verify"])
		return
	} else if ue.Admin == false {
		HyperText.HttpErrorResponse(w, http.StatusBadRequest, HyperText.CustomResponses["not-admin"])
		return
	}
	if err := c.GameEntityRepository.InsertGameEntity(&age); err != nil {
		HyperText.HttpErrorResponse(w, http.StatusBadRequest, err)
		return
	}
	HyperText.HttpResponse(w, http.StatusOK, age)
	return
}

//-------------------------------- USR --------------------------------//
func (c *GameEntityController) UpdateSingle(w http.ResponseWriter, r *http.Request) {
	ageus := AdminGameEntityUpdateSingle{}
	Interface.PutDOCNAME(DOCNAME)
	if err := HyperText.BodyValidate(r, &ageus); err != nil {
		HyperText.HttpErrorResponse(w, http.StatusBadRequest, err)
		return
	}
	if err := ValidateUpdateSingle(ageus); err != nil {
		HyperText.HttpErrorResponse(w, http.StatusBadRequest, err)
		return
	}
	uev := UserEntity.UserEntityVerify{}
	uev.Username = ageus.Username
	uev.Password = ageus.Password
	if _, boolean := UserEntity.AdminVerify(uev); boolean == false {
		HyperText.HttpErrorResponse(w, http.StatusBadRequest, HyperText.CustomResponses["wrong-verify"])
		return
	}
	if err := Interface.UpdateSingleDB(DOCNAME, ID_NAME, ageus.Name, ageus.Position, ageus.Value, ageus); err != nil {
		HyperText.HttpErrorResponse(w, http.StatusBadRequest, HyperText.CustomResponses["error-update"])
		return
	}
	HyperText.HttpResponse(w, http.StatusOK, HyperText.CustomResponses["success-update"])
	return
}

//-------------------------------- ADM --------------------------------//
func (c *GameEntityController) Update(w http.ResponseWriter, r *http.Request) {
	var age AdminGameEntity
	Interface.PutDOCNAME(DOCNAME)
	age.Name = mux.Vars(r)["name"]
	if err := HyperText.BodyValidate(r, &age); err != nil {
		HyperText.HttpErrorResponse(w, http.StatusBadRequest, err)
		return
	}
	uev := UserEntity.UserEntityVerify{}
	uev.Username = age.Username
	uev.Password = age.Password
	ue, boolean := UserEntity.AdminVerify(uev)
	if boolean == false {
		HyperText.HttpErrorResponse(w, http.StatusBadRequest, HyperText.CustomResponses["wrong-verify"])
		return
	} else if ue.Admin == false {
		HyperText.HttpErrorResponse(w, http.StatusBadRequest, HyperText.CustomResponses["not-admin"])
		return
	}
	age.Name = strings.ToLower(age.Name)
	err := Interface.UpdateDB(DOCNAME, ID_NAME, age.Name, &age)
	if err != nil {
		HyperText.HttpErrorResponse(w, http.StatusBadRequest, HyperText.CustomResponses["error-update"])
		return
	}
	HyperText.HttpResponse(w, http.StatusOK, age)
	return
}

//-------------------------------- ADM --------------------------------//
func (c *GameEntityController) Get(w http.ResponseWriter, r *http.Request) {
	name := mux.Vars(r)["name"]
	name = strings.ToLower(name)
	ge := GameEntity{}
	if err := Interface.FindUnitDB(DOCNAME, ID_NAME, name, &ge); err != nil {
		HyperText.HttpErrorResponse(w, http.StatusBadRequest, HyperText.CustomResponses["not-found-entity"])
		return
	}
	HyperText.HttpResponse(w, http.StatusOK, ge)
	return
}

//-------------------------------- ADM --------------------------------//
func (c *GameEntityController) GetAllWhile(w http.ResponseWriter, r *http.Request) {
	position := mux.Vars(r)["position"]
	value := mux.Vars(r)["value"]
	position = strings.ToLower(position)
	value = strings.ToLower(value)
	entities := GameEntities{}
	if err := Interface.FindAllWhileDB(DOCNAME, position, value, &entities); err != nil {
		HyperText.HttpErrorResponse(w, http.StatusBadRequest, HyperText.CustomResponses["empty-Interface"])
		return
	}
	HyperText.HttpResponse(w, http.StatusOK, entities)
	return
}

//-------------------------------- ADM --------------------------------//
func (c *GameEntityController) GetAll(w http.ResponseWriter, r *http.Request) {
	ges := GameEntities{}
	if err := Interface.FindAllDB(DOCNAME, &ges); err != nil {
		HyperText.HttpErrorResponse(w, http.StatusBadRequest, HyperText.CustomResponses["empty-Interface"])
		return
	}
	HyperText.HttpResponse(w, http.StatusOK, ges)
	return
}

//-------------------------------- ADM --------------------------------//
func (c *GameEntityController) Delete(w http.ResponseWriter, r *http.Request) {
	name := strings.ToLower(mux.Vars(r)["name"])
	ge := GameEntity{}
	if err := Interface.DeleteDB(DOCNAME, ID_NAME, name, &ge); err != nil {
		HyperText.HttpErrorResponse(w, http.StatusBadRequest, HyperText.CustomResponses["not-found-entity"])
		return
	}
	HyperText.HttpResponse(w, http.StatusOK, ge)
	return
}
