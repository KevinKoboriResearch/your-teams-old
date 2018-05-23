package UserEntity

import (
	"strings"
	"be/Database"
)

const DOCNAME = "user_games"

//-------------------------------- BG --------------------------------//
func (r UserEntityRepository) InsertUserEntity(ue *UserEntity) (err error) {
	ue.Username = strings.ToLower(ue.Username)
	ue.Email = strings.ToLower(ue.Email)
	ue.Enable = true
	ue.Password, _ = GenerateHashPassword(ue.Password)
	if err = Database.InsertDB(&ue, DOCNAME); err != nil {
		return
	}
	return
}

//-------------------------------- BG --------------------------------//
func (r UserEntityRepository) VerifyUserEntity(uev UserEntityVerify) (UserEntityUpdate, bool) {
	ueu := UserEntityUpdate{}
	if err := Database.FindUnitDB(uev.Username, &ueu, DOCNAME); err != nil {
		return ueu, false
	}
	boolean := CheckPasswordHash(uev.Password, ueu.Password)
	return ueu, boolean
}

//-------------------------------- BG --------------------------------//
func (r UserEntityRepository) EnableUserEntity(ueu UserEntityUpdate) error {
	if ueu.Enable == false {
		ueu.Enable = true
		if _, err := Database.UpdateDB(ueu, DOCNAME, ueu.Username); err != nil {
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
		if _, err := Database.UpdateDB(ueu, DOCNAME, ueu.Username); err != nil {
			return err
		}
		return nil
	}
	return nil
}
