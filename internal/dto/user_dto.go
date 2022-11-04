package dto

type UserDTO struct {
	Fullname string `json:"fullname" form:"fullname"`
	Phone    string `json:"phone" form:"phone"`
	Address  string `json:"address" form:"address"`
	Role     string `json:"role" form:"role"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}
