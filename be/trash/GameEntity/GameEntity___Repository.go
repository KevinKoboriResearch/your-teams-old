package GameEntity

import (
	"be/Interface"
	"strings"
)

//-------------------------------- USR --------------------------------//
func (r GameEntityRepository) InsertGameEntity(age *AdminGameEntity) (err error) {
	var ge GameEntity
	ge.Name = strings.ToLower(age.Name)
	ge.Abbre = strings.ToLower(age.Abbre)
	ge.Desc = strings.ToLower(age.Desc)
	ge.Enable = true
	err = Interface.InsertDB(DOCNAME, &ge)
	return
}

//-------------------------------- USR --------------------------------//
func (r GameEntityRepository) EnableGameEntity(ge GameEntity) (err error) {
	if ge.Enable == false {
		ge.Enable = true
		if err = Interface.UpdateDB(DOCNAME, ID_NAME, ge.Name, ge); err != nil {
			return
		}
		return
	}
	return
}

//-------------------------------- USR --------------------------------//
func (r GameEntityRepository) DisableGameEntity(ge GameEntity) (err error) {
	if ge.Enable == true {
		ge.Enable = false
		if err = Interface.UpdateDB(DOCNAME, ID_NAME, ge.Name, ge); err != nil {
			return
		}
		return
	}
	return
}
