package GameEntity

import (
	"strings"
	"be/Interface"
)

const DOCNAME = "game_entity"

//-------------------------------- USR --------------------------------//
func (r GameEntityRepository) InsertGameEntity(age *AdminGameEntity) (err error) {
	var ge GameEntity
	ge.Name = strings.ToLower(age.Name)
	ge.Desc = strings.ToLower(age.Desc)
	ge.Enable = true
	err = Interface.InsertDB(&ge, DOCNAME)
	return
}

//-------------------------------- USR --------------------------------//
func (r GameEntityRepository) EnableGameEntity(ge GameEntity) error {
	if ge.Enable == false {
		ge.Enable = true
		if _, err := Interface.UpdateDB(ge, DOCNAME, ge.Name); err != nil {
			return err
		}
		return nil
	}
	return nil
}

//-------------------------------- USR --------------------------------//
func (r GameEntityRepository) DisableGameEntity(ge GameEntity) error {
	if ge.Enable == true {
		ge.Enable = false
		if _, err := Interface.UpdateDB(ge, DOCNAME, ge.Name); err != nil {
			return err
		}
		return nil
	}
	return nil
}
