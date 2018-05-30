package UserEntity

type UserEntityRepository struct {
}

type UserEntityController struct {
	UserEntityRepository UserEntityRepository
}

type UserEntity struct {
	Username string `json:"username" validate:"required,username-length,username-used"`
	Email    string `json:"email" validate:"required,email,email-used"`
	Password string `json:"password" validate:"required,password-length"`
	Image    string `json:"image" validate:"omitempty,url"`
	Admin    bool   `json:"admin"`
	Enable   bool   `json:"enable"`
}

type UserEntityUpdate struct {
	Username string `json:"username" bson:"username,omitempty"`
	Email    string `json:"email" bson:"email,omitempty"`
	Image    string `json:"image" bson:"image,omitempty"`
	Admin    bool   `json:"admin" bson:"admin,omitempty"`
	Enable   bool   `json:"enable" bson:"enable,omitempty"`
}

type UserEntityUpdateSingle struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Position string `json:"position" validate:"required"`
	Value    string `json:"value" validate:"required"`
}

type UserEntityVerify struct {
	Username string `json:"username"`
	Password string `json:"password" validate:"required,password-length"`
	Admin    bool   `json:"admin"`
	Enable   bool   `json:"enable"`
}

type UserEntityProtected struct {
	Username string `json:"username"`
	Email    string `json:"email" validate="required"`
	Image    string `json:"image" validate="omitempty,required"`
	Admin    bool   `json:"admin"`
	Enable   bool   `json:"enable"`
}

type UserEntities []UserEntityProtected

type UserUsername struct {
	Username string `json:"username" validate:"required,username-length,username-exist"`
}

type UserEmail struct {
	Email string `json:"email" validate:"required,email,email-used"`
}

type UserImage struct {
	Image string `json:"image" validate:"omitempty,url"`
}

type UserPassword struct {
	Password string `json:"password" validate:"required,password-length"`
}
