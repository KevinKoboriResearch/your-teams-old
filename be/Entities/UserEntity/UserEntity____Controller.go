package UserEntity

import (
	"github.com/gorilla/mux"
	"net/http"
	"strings"
	"be/Interface"
	"be/HyperText"
	"be/auth"
)

//-------------------------------- USR --------------------------------//
func (c *UserEntityController) SignUp(w http.ResponseWriter, r *http.Request) {
	var ue UserEntity
	if err := HyperText.BodyValidate(r, &ue); err != nil {
		HyperText.HttpErrorResponse(w, http.StatusBadRequest, err)
		return
	}
	if err := c.UserEntityRepository.InsertUserEntity(&ue); err != nil {
		HyperText.HttpErrorResponse(w, http.StatusBadRequest, err)
		return
	}
	HyperText.HttpResponse(w, http.StatusOK, ue)
	return
}

//-------------------------------- USR --------------------------------//
func (c *UserEntityController) Login(w http.ResponseWriter, r *http.Request) {
	uev := UserEntityVerify{}
	if err := HyperText.BodyValidate(r, &uev); err != nil {
		HyperText.HttpErrorResponse(w, http.StatusBadRequest, err)
		return
	}
	ue, boolean := c.UserEntityRepository.VerifyUserEntity(uev)
	if boolean == false {
		HyperText.HttpErrorResponse(w, http.StatusUnauthorized, HyperText.CustomResponses["wrong-verify"])
		return
	}
	if err := c.UserEntityRepository.EnableUserEntity(ue); err != nil {
		HyperText.HttpErrorResponse(w, http.StatusBadRequest, HyperText.CustomResponses["error-enable"])
		return
	}
	token := auth.GenerateJWT(ue)
	w.Header().Add("Authorization", "Bearer " + token)
	HyperText.HttpResponse(w, http.StatusOK, HyperText.CustomResponses["success-Login"])
	return
}

//-------------------------------- USR --------------------------------//
func (c *UserEntityController) UpdateSingle(w http.ResponseWriter, r *http.Request) {
	ueus := UserEntityUpdateSingle{}
	ueus.Username = mux.Vars(r)["username"]
	if err := HyperText.BodyValidate(r, &ueus); err != nil {
		HyperText.HttpErrorResponse(w, http.StatusBadRequest, err)
		return
	}
	if err := ValidateUpdateSingle(ueus); err != nil {
		HyperText.HttpErrorResponse(w, http.StatusBadRequest, err)
		return
	}
	uev := UserEntityVerify{}
	uev.Username = ueus.Username
	uev.Password = ueus.Password
	if _, boolean := c.UserEntityRepository.VerifyUserEntity(uev); boolean == false {
		HyperText.HttpErrorResponse(w, http.StatusBadRequest, HyperText.CustomResponses["wrong-verify"])
		return
	}
	if _, err := Interface.UpdateSingleDB(ueus.Position, ueus.Value, ueus, DOCNAME, ueus.Username); err != nil {
		HyperText.HttpErrorResponse(w, http.StatusBadRequest, HyperText.CustomResponses["error-update"])
		return
	}
	HyperText.HttpResponse(w, http.StatusOK, HyperText.CustomResponses["success-update"])
	return
}

//-------------------------------- USR --------------------------------//
func (c *UserEntityController) UpdatePartial(w http.ResponseWriter, r *http.Request) {
	username := mux.Vars(r)["username"]
	ueu := UserEntityUpdate{}
	if err := HyperText.BodyValidate(r, &ueu); err != nil {
		HyperText.HttpErrorResponse(w, http.StatusBadRequest, err)
		return
	}
	uev := UserEntityVerify{}
	uev.Username = ueu.Username
	uev.Password = ueu.Password
	if _, boolean := c.UserEntityRepository.VerifyUserEntity(uev); boolean == false {
		HyperText.HttpErrorResponse(w, http.StatusBadRequest, HyperText.CustomResponses["wrong-Login"])
		return
	}
	ueu.Password = ``
	result, err := Interface.UpdatePartialDB(ueu, DOCNAME, username)
	if err != nil {
		HyperText.HttpErrorResponse(w, http.StatusBadRequest, HyperText.CustomResponses["error-update"])
		return
	}
	HyperText.HttpResponse(w, http.StatusOK, result)
	return
}

//-------------------------------- USR --------------------------------//
func (c *UserEntityController) Update(w http.ResponseWriter, r *http.Request) {
	ueu := UserEntityUpdate{}
	ueu.Username = mux.Vars(r)["username"]
	if err := HyperText.BodyValidate(r, &ueu); err != nil {
		HyperText.HttpErrorResponse(w, http.StatusBadRequest, err)
		return
	}
	uev := UserEntityVerify{}
	uev.Username = ueu.Username
	uev.Password = ueu.Password
	if _, boolean := c.UserEntityRepository.VerifyUserEntity(uev); boolean == false {
		HyperText.HttpErrorResponse(w, http.StatusBadRequest, HyperText.CustomResponses["wrong-Login"])
		return
	}
	result, err := Interface.UpdateDB(ueu, DOCNAME, ueu.Username)
	if err != nil {
		HyperText.HttpErrorResponse(w, http.StatusBadRequest, HyperText.CustomResponses["error-update"])
		return
	}
	HyperText.HttpResponse(w, http.StatusOK, result)
	return
}

//-------------------------------- USR --------------------------------//
func (c *UserEntityController) Disable(w http.ResponseWriter, r *http.Request) {
	username := mux.Vars(r)["username"]
	username = strings.ToLower(username)
	uev := UserEntityVerify{}
	if 	err := HyperText.BodyValidate(r, &uev); err != nil {
		HyperText.HttpErrorResponse(w, http.StatusBadRequest, err)
		return
	}
	uev.Username = username
	ueu, boolean := c.UserEntityRepository.VerifyUserEntity(uev)
	if boolean == false {
		HyperText.HttpErrorResponse(w, http.StatusBadRequest, HyperText.CustomResponses["wrong-Login"])
		return
	}
	if err := c.UserEntityRepository.DisableUserEntity(ueu); err != nil {
		HyperText.HttpErrorResponse(w, http.StatusBadRequest, HyperText.CustomResponses["error-disable"])
		return
	}
	HyperText.HttpResponse(w, http.StatusOK, HyperText.CustomResponses["success-disabled"])
	return
}

//-------------------------------- USR --------------------------------//
func (c *UserEntityController) GetUnit(w http.ResponseWriter, r *http.Request) {
	username := mux.Vars(r)["username"]
	username = strings.ToLower(username)
	result := UserEntityProtected{}
	if 	err := Interface.FindUnitDB(username, &result, DOCNAME); err != nil {
		HyperText.HttpErrorResponse(w, http.StatusBadRequest, HyperText.CustomResponses["not-found-entity"])
		return
	}
	HyperText.HttpResponse(w, http.StatusOK, result)
	return
}

//-------------------------------- USR --------------------------------//
func (c *UserEntityController) GetAllEnabledWhile(w http.ResponseWriter, r *http.Request) {
	position := mux.Vars(r)["position"]
	value := mux.Vars(r)["value"]
	position = strings.ToLower(position)
	value = strings.ToLower(value)
	entities := UserEntities{}
	if 	err := Interface.FindAllEnabledWhileDB(position, value, &entities, DOCNAME); err != nil {
		HyperText.HttpErrorResponse(w, http.StatusBadRequest, HyperText.CustomResponses["empty-Interface"])
		return
	}
	HyperText.HttpResponse(w, http.StatusOK, entities)
	return
}

//-------------------------------- ADM --------------------------------//
func (c *UserEntityController) GetAllEnabled(w http.ResponseWriter, r *http.Request) {
	result := UserEntities{}
	if 	err := Interface.FindAllEnabledDB(&result, DOCNAME); err != nil {
		HyperText.HttpErrorResponse(w, http.StatusBadRequest, HyperText.CustomResponses["empty-Interface"])
		return
	}
	HyperText.HttpResponse(w, http.StatusOK, result)
	return
}

//-------------------------------- ADM --------------------------------//
func (c *UserEntityController) GetAll(w http.ResponseWriter, r *http.Request) {
	result := UserEntities{}
	if 	err := Interface.FindAllDB(&result, DOCNAME); err != nil {
		HyperText.HttpErrorResponse(w, http.StatusBadRequest, HyperText.CustomResponses["empty-Interface"])
		return
	}
	HyperText.HttpResponse(w, http.StatusOK, result)
	return
}

//-------------------------------- ADM --------------------------------//
func (c *UserEntityController) Delete(w http.ResponseWriter, r *http.Request) {
	username := strings.ToLower(mux.Vars(r)["username"])
	result := UserEntity{}
	if 	err := Interface.DeleteDB(username, &result, DOCNAME); err != nil {
		HyperText.HttpErrorResponse(w, http.StatusBadRequest, HyperText.CustomResponses["not-found-entity"])
		return
	}
	HyperText.HttpResponse(w, http.StatusOK, result)
	return
}
