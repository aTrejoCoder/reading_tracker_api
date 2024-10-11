package dtos

type SignupDTO struct {
	Username            string           `json:"username" validate:"required,min=3,max=32"`
	Email               string           `json:"email" validate:"required,email"`
	Password            string           `json:"password" validate:"required,min=8"`
	ProprofileInsertDTO ProfileInsertDTO `json:"profile_insert_dto"`
}

type LoginDTO struct {
	LoginIdentfier string `json:"login_identifier" validate:"required"` //username or email
	Password       string `json:"password" validate:"required"`
}
