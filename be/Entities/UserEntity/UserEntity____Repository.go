package UserEntity

import (
	"be/Interface"
	"log"
	"strings"
)

//___________________________ INSERT _________________________________________//
func (r UserEntityRepository) InsertUserEntity(ue UserEntity) (err error) {

	ue.Username = strings.ToLower(ue.Username)
	ue.Email = strings.ToLower(ue.Email)
	ue.Password, _ = GenerateHashPassword(ue.Password)
	ue.Enable = true

	err = Interface.InsertDB(DOCNAME, ue)

	log.Print("[ERROR] failed to delete entity: ", err)

	return
}

//___________________________ VERIFY _________________________________________//
func (r UserEntityRepository) VerifyUserEntity(c *UserEntityController, uev UserEntityVerify) string {

	uev.Username = strings.ToLower(uev.Username)

	err := Interface.FindDB(DOCNAME, ID_USERNAME, uev.Username, &ue)

	if err != nil {
		return err.Error()
	}

	boolean := CheckPasswordHash(uev.Password, ue.Password)

	if boolean == false {
		return "Wrong Password"
	}

	err = c.UserEntityRepository.EnableUserEntity(ue)

	if err != nil {
		return err.Error()
	}

	return ""
}

//___________________________ UPDATE SINGLE __________________________________//
func (r UserEntityRepository) UpdateSingleUserEntity(ueus UserEntityUpdateSingle) (err error) {

	ueus.Username = strings.ToLower(ueus.Username)
	ueus.Position = strings.ToLower(ueus.Position)

	if ueus.Position == "password" {
		ueus.Value, _ = GenerateHashPassword(ueus.Value)
		err = Interface.UpdateSingleDB(DOCNAME, ID_USERNAME, ueus.Username, ueus.Position, ueus.Value)
		return
	}

	ueus.Value = strings.ToLower(ueus.Value)

	err = Interface.UpdateSingleDB(DOCNAME, ID_USERNAME, ueus.Username, ueus.Position, ueus.Value)

	return
}

//___________________________ UPDATE _________________________________________//
func (r UserEntityRepository) UpdateUserEntity(ueu UserEntityUpdate) (err error) {

	ueu.Username = strings.ToLower(ueu.Username)
	ueu.Email = strings.ToLower(ueu.Email)

	err = Interface.UpdateDB(DOCNAME, ID_USERNAME, ueu.Username, ueu)

	return
}

//___________________________ GET ENTITY _____________________________________//
func (r UserEntityRepository) GetUserEntity(uep *UserEntityProtected) (err error) {

	err = Interface.FindDB(DOCNAME, ID_USERNAME, uep.Username, &uep)

	return
}

//___________________________ GET ALL ENABLED WHILE __________________________//
func (r UserEntityRepository) GetAllEnabledWhileUserEntities(position string, value string, ues UserEntities) (UserEntities, error) {

	err := Interface.FindAllEnabledWhileDB(DOCNAME, position, value, &ues)

	return ues, err
}

//___________________________ GET ALL ENABLED ________________________________//
func (r UserEntityRepository) GetAllEnabledUserEntities(ues UserEntities) (UserEntities, error) {

	err := Interface.FindAllEnabledDB(DOCNAME, &ues)

	return ues, err
}

//___________________________ GET ALL ________________________________________//
func (r UserEntityRepository) GetAllUserEntities(ues UserEntities) (UserEntities, error) {

	err := Interface.FindAllDB(DOCNAME, &ues)

	return ues, err
}

//___________________________ DELETE _________________________________________//
func (r UserEntityRepository) DeleteUserEntity(username string) (err error) {

	err = Interface.DeleteDB(DOCNAME, ID_USERNAME, username, ue)

	return
}

//___________________________ ENABLE _________________________________________//
func (r UserEntityRepository) EnableUserEntity(ue UserEntity) (err error) {

	if ue.Enable != true {
		err = Interface.UpdateSingleDB(DOCNAME, ID_USERNAME, ue.Username, "enable", true)
	}

	return
}

//___________________________ DISABLE ________________________________________//
func (r UserEntityRepository) DisableUserEntity(uev UserEntityVerify) string {

	uev.Username = strings.ToLower(uev.Username)

	err := Interface.FindDB(DOCNAME, ID_USERNAME, uev.Username, &ue)

	if err != nil {
		return err.Error()
	}

	boolean := CheckPasswordHash(uev.Password, ue.Password)

	if boolean == false {
		return "Wrong Password"
	}

	err = Interface.UpdateSingleDB(DOCNAME, ID_USERNAME, uev.Username, "enable", false)

	if err != nil {
		return err.Error()
	}

	return ""
}

//__ ADMIN VERIFY ____________________________________________________________//
func AdminVerify(uev UserEntityVerify) (UserEntity, bool) {

	uev.Username = strings.ToLower(uev.Username)

	if err := Interface.FindDB(DOCNAME, ID_USERNAME, uev.Username, &ue); err != nil {
		return ue, false
	}

	if boolean := CheckPasswordHash(uev.Password, ue.Password); boolean == true {
		if ue.Admin == true {
			return ue, true
		}
	}

	return ue, false
}
