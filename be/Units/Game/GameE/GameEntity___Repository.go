package GameEntity

import (
	"strings"
	"be/Database"
)

const DOCNAME = "user_entity"

//-------------------------------- USR --------------------------------//
func (r GameEntityRepository) InsertGameEntity(ue *GameEntity) (err error) {
	ue.Username = strings.ToLower(ue.Username)
	ue.Email = strings.ToLower(ue.Email)
	ue.Enable = true
	ue.Password, _ = GenerateHashPassword(ue.Password)
	err = Database.InsertDB(&ue, DOCNAME)
	if err != nil {
		return
	}
	return
}

//-------------------------------- USR --------------------------------//
func (r GameEntityRepository) VerifyGameEntity(uev GameEntityVerify) (GameEntityUpdate, bool) {
	ueu := GameEntityUpdate{}
	err := Database.FindOneDB(uev.Username, &ueu, DOCNAME)
	if err != nil {
		return ueu, false
	} else if boolean := CheckPasswordHash(uev.Password, ueu.Password); boolean == true {
		return ueu, true
	}
	return ueu, false
}

//-------------------------------- USR --------------------------------//
func (r GameEntityRepository) EnableGameEntity(ueu GameEntityUpdate) error {
	if ueu.Enable == false {
		ueu.Enable = true
		if _, err := Database.UpdateDB(ueu, DOCNAME, ueu.Username); err != nil {
			return err
		}
		return nil
	}
	return nil
}

//-------------------------------- USR --------------------------------//
func (r GameEntityRepository) DisableGameEntity(ueu GameEntityUpdate) error {
	if ueu.Enable == true {
		ueu.Enable = false
		if _, err := Database.UpdateDB(ueu, DOCNAME, ueu.Username); err != nil {
			return err
		}
		return nil
	}
	return nil
}
