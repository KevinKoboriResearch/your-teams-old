package UserEntity

type UserEntityRepository struct{
}

type UserEntityController struct {
	UserEntityRepository UserEntityRepository
}

type ByOne struct {
	Position string `json:"position" validate:"required"`
	Value    string `json:"value" validate:"required"`
}

type UserEntity struct {
	Username string `json:"username" validate:"required,username-length,username-used"`
	Email    string `json:"email" validate:"required,email,email-used"`
	Password string `json:"password" validate:"required,password-length"`
	Image    string `json:"image" validate:"omitempty,url"`
	Admin    bool   `json:"admin"`
	Enable   bool   `json:"enable"`
}

type UserEntities []UserEntityProtected

type UserEntityProtected struct {
	Username string `json:"username"`
	Email    string `json:"email" validate="required"`
	Image    string `json:"image" validate="omitempty,required"`
	Admin    bool   `json:"admin"`
	Enable   bool   `json:"enable"`
}

type UserEntityUpdate struct {
	Username string `json:"username"`
	Email    string `json:"email" validate:"email,email-used"`
	Password string `json:"password" validate:"required,password-length"`
	Image    string `json:"image" validate:"omitempty,url"`
	Enable   bool   `json:"enable"`
}

type UserEntityUpdateSingle struct {
	Username string `json:"username"`
	Password string `json:"password" validate:"required,password-length"`
	Position string `json:"position" validate:"required"`
	Value    string `json:"value" validate:"required"`
}

type UserEntityVerify struct {
	Username string `json:"username"`
	Password string `json:"password" validate:"required"`
	Admin    bool   `json:"admin"`
	Enable   bool   `json:"enable"`
}

type UserUsername struct {
	Username string `json:"username" validate:"required,username-length,username-used"`
}

type UserEmail struct {
	Email string `json:"email" validate:"required,email,email-used"`
}
