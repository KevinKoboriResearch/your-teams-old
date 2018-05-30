package UserEntity

import (
	"be/HyperText"
)

const (
	DOCNAME     = "user_entity"
	ID_USERNAME = "username"
)

var (
	controller = &UserEntityController{UserEntityRepository: UserEntityRepository{}}
	routes     = HyperText.Routes{{}}
	ue UserEntity
	uev UserEntityVerify
	ueus UserEntityUpdateSingle
	ueu UserEntityUpdate
	uep UserEntityProtected
	ues UserEntities
)
