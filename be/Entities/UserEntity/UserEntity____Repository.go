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
func (r UserEntityRepository) VerifyUserEntity(uev UserEntityVerify) (UserEntity, bool) {
	uev.Username = strings.ToLower(uev.Username)
	ue := UserEntity{}
	if err := Interface.FindUnitDB(uev.Username, &ue, DOCNAME); err != nil {
		return ue, false
	}
	boolean := CheckPasswordHash(uev.Password, ue.Password)
	return ue, boolean
}

//-------------------------------- BG --------------------------------//
func (r UserEntityRepository) EnableUserEntity(ue UserEntity) error {
	if ue.Enable == false {
		ue.Enable = true
		if _, err := Interface.UpdateDB(ue, DOCNAME, ue.Username); err != nil {
			return err
		}
		return nil
	}
	return nil
}

//-------------------------------- BG --------------------------------//
func (r UserEntityRepository) DisableUserEntity(ue UserEntity) error {
	if ue.Enable == true {
		ue.Enable = false
		if _, err := Interface.UpdateDB(ue, DOCNAME, ue.Username); err != nil {
			return err
		}
		return nil
	}
	return nil
}
