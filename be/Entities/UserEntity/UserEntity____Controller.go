package UserEntity

import (
	"be/HyperText"
	"be/Interface"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strings"
)

//____________________________ INSERT ________________________________________//
func (c *UserEntityController) SignUp(w http.ResponseWriter, r *http.Request) {

	err := HyperText.BodyValidate(r, &ue)

	if err != nil {
		HyperText.HttpSpecificErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	err = c.UserEntityRepository.InsertUserEntity(ue)
	if err != nil {
		HyperText.HttpErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	HyperText.HttpResponse(w, http.StatusOK, HyperText.CustomResponses["success-insert"])
	return
}

//____________________________ VERIFY ________________________________________//
func (c *UserEntityController) Login(w http.ResponseWriter, r *http.Request) {

	err := HyperText.BodyValidate(r, &uev)

	if err != nil {
		HyperText.HttpErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	err = c.UserEntityRepository.VerifyUserEntity(c, uev)

	if err != "" {
		HyperText.HttpErrorResponse(w, http.StatusUnauthorized, err)
		return
	}

	token := Interface.GenerateJWT(uev)

	w.Header().Add("authorization", "Bearer " + token)
	HyperText.HttpResponse(w, http.StatusOK, HyperText.CustomResponses["success-Login"])
	return
}

//____________________________ UPDATE SINGLE _________________________________//
func (c *UserEntityController) UpdateSingle(w http.ResponseWriter, r *http.Request) {

	ueus.Username = strings.ToLower(mux.Vars(r)["username"])

	err := HyperText.BodyValidate(r, &ueus)

	if err != nil {
		HyperText.HttpErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	err = ValidateUpdateSingle(c, ueus)

	if err != nil {
		HyperText.HttpErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	err = c.UserEntityRepository.UpdateSingleUserEntity(ueus)

	if err != nil {
		HyperText.HttpErrorResponse(w, http.StatusBadRequest, HyperText.CustomResponses["error-update"])
		return
	}

	HyperText.HttpResponse(w, http.StatusOK, HyperText.CustomResponses["success-update"])
	return
}

//____________________________ UPDATE ________________________________________//
func (c *UserEntityController) Update(w http.ResponseWriter, r *http.Request) {

	ueu.Username = strings.ToLower(mux.Vars(r)["username"])

	err := HyperText.BodyValidate(r, &ueu)

	if err != nil {
		HyperText.HttpErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	err = ValidateUpdate(ueu)

	if err != nil {
		HyperText.HttpErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	err = c.UserEntityRepository.UpdateUserEntity(ueu)

	if err != nil {
		HyperText.HttpErrorResponse(w, http.StatusBadRequest, HyperText.CustomResponses["error-update"])
		log.Println(err)
		return
	}

	HyperText.HttpResponse(w, http.StatusOK, HyperText.CustomResponses["success-update"])
	return
}

//____________________________ DISABLE _______________________________________//
func (c *UserEntityController) Disable(w http.ResponseWriter, r *http.Request) {

	uev.Username = strings.ToLower(mux.Vars(r)["username"])

	err := HyperText.BodyValidate(r, &uev)

	if err != nil {
		HyperText.HttpErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	err = c.UserEntityRepository.DisableUserEntity(uev)

	if err != nil {
		HyperText.HttpErrorResponse(w, http.StatusBadRequest, HyperText.CustomResponses["error-disable"])
		return
	}

	HyperText.HttpResponse(w, http.StatusOK, HyperText.CustomResponses["success-disabled"])
	return
}

//____________________________ GET UNIQUE ____________________________________//
func (c *UserEntityController) GetUnique(w http.ResponseWriter, r *http.Request) {

	uep.Username = strings.ToLower(mux.Vars(r)["username"])

	err := c.UserEntityRepository.GetUserEntity(&uep)

	if err != nil {
		HyperText.HttpErrorResponse(w, http.StatusBadRequest, HyperText.CustomResponses["not-found-entity"])
		return
	}

	HyperText.HttpResponse(w, http.StatusOK, uep)
	return
}

//____________________________ GET ALL ENABLED WHILE _________________________//
func (c *UserEntityController) GetAllEnabledWhile(w http.ResponseWriter, r *http.Request) {

	position := mux.Vars(r)["position"]
	value := mux.Vars(r)["value"]

	ues, err := c.UserEntityRepository.GetAllEnabledWhileUserEntities(position, value, ues)

	if err != nil {
		HyperText.HttpErrorResponse(w, http.StatusBadRequest, HyperText.CustomResponses["empty-Interface"])
		return
	}

	HyperText.HttpResponse(w, http.StatusOK, ues)
	return
}

//____________________________ GET ALL ENABLED _______________________________//
func (c *UserEntityController) GetAllEnabled(w http.ResponseWriter, r *http.Request) {

	ues, err := c.UserEntityRepository.GetAllEnabledUserEntities(ues)

	if err != nil {
		HyperText.HttpErrorResponse(w, http.StatusBadRequest, HyperText.CustomResponses["empty-Interface"])
		return
	}

	HyperText.HttpResponse(w, http.StatusOK, ues)
	return
}

//____________________________ GET ALL _______________________________________//
func (c *UserEntityController) GetAll(w http.ResponseWriter, r *http.Request) {

	ues, err := c.UserEntityRepository.GetAllUserEntities(ues)

	if err != nil {
		HyperText.HttpErrorResponse(w, http.StatusBadRequest, HyperText.CustomResponses["empty-Interface"])
		return
	}

	HyperText.HttpResponse(w, http.StatusOK, ues)
	return
}

//____________________________ DELETE ________________________________________//
func (c *UserEntityController) Delete(w http.ResponseWriter, r *http.Request) {

	username := strings.ToLower(mux.Vars(r)["username"])

	err := c.UserEntityRepository.DeleteUserEntity(username)

	if err != nil {
		HyperText.HttpErrorResponse(w, http.StatusBadRequest, HyperText.CustomResponses["not-found-entity"])
		return
	}

	HyperText.HttpResponse(w, http.StatusOK, HyperText.CustomResponses["success-Delete"])
	return
}
