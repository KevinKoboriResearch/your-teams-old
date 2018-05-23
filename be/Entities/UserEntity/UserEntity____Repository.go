package UserEntity

import (
	"strings"
	"be/Interface"
)

const DOCNAME = "user_entity"

//-------------------------------- BG --------------------------------//
func AdminVerify(uev UserEntityVerify) (UserEntity, bool) {
	uev.Username = strings.ToLower(uev.Username)
	ue := UserEntity{}
	if err := Interface.FindUnitDB(uev.Username, &ue, DOCNAME); err != nil {
		return ue, false
	}
	if boolean := CheckPasswordHash(uev.Password, ue.Password); boolean == true {
		if ue.Admin == true {
			return ue, true
		}
	}
	return ue, false
}

//-------------------------------- BG --------------------------------//
func (r UserEntityRepository) InsertUserEntity(ue *UserEntity) (err error) {
	ue.Username = strings.ToLower(ue.Username)
	ue.Email = strings.ToLower(ue.Email)
	ue.Enable = true
	ue.Password, _ = GenerateHashPassword(ue.Password)
	err = Interface.InsertDB(&ue, DOCNAME)
	return
}

//-------------------------------- BG --------------------------------//
func (r UserEntityRepository) UpdateSingleUserEntity(userE *UserEntity) (err error) {
	userE.Username = strings.ToLower(userE.Username)
	userE.Email = strings.ToLower(userE.Email)
	userE.Enable = true
	userE.Password, _ = GenerateHashPassword(userE.Password)
	err = Interface.InsertDB(&userE, DOCNAME)
	return
}

//-------------------------------- BG --------------------------------//
func (r UserEntityRepository) UpdatePartialUserEntity(ue *UserEntity) (err error) {
	ue.Username = strings.ToLower(ue.Username)
	ue.Email = strings.ToLower(ue.Email)
	ue.Enable = true
	ue.Password, _ = GenerateHashPassword(ue.Password)
	err = Interface.InsertDB(&ue, DOCNAME)
	return
}

//-------------------------------- BG --------------------------------//
func (r UserEntityRepository) UpdateUserEntity(ue *UserEntity) (err error) {
	ue.Username = strings.ToLower(ue.Username)
	ue.Email = strings.ToLower(ue.Email)
	ue.Enable = true
	ue.Password, _ = GenerateHashPassword(ue.Password)
	err = Interface.InsertDB(&ue, DOCNAME)
	return
}

//-------------------------------- BG --------------------------------//
func (r UserEntityRepository) VerifyUserEntity(uev UserEntityVerify) (UserEntityUpdate, bool) {
	uev.Username = strings.ToLower(uev.Username)
	ueu := UserEntityUpdate{}
	if err := Interface.FindUnitDB(uev.Username, &ueu, DOCNAME); err != nil {
		return ueu, false
	}
	boolean := CheckPasswordHash(uev.Password, ueu.Password)
	return ueu, boolean
}

//-------------------------------- BG --------------------------------//
func (r UserEntityRepository) EnableUserEntity(ueu UserEntityUpdate) error {
	if ueu.Enable == false {
		ueu.Enable = true
		if _, err := Interface.UpdateDB(ueu, DOCNAME, ueu.Username); err != nil {
			return err
		}
		return nil
	}
	return nil
}

//-------------------------------- BG --------------------------------//
func (r UserEntityRepository) DisableUserEntity(ueu UserEntityUpdate) error {
	if ueu.Enable == true {
		ueu.Enable = false
		if _, err := Interface.UpdateDB(ueu, DOCNAME, ueu.Username); err != nil {
			return err
		}
		return nil
	}
	return nil
}
