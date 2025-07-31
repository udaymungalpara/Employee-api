package types

type Employee struct {
	Id         int64  `json:"id"`
	Name       string `json:"name" validate:"required"`
	Email      string `json:"email" validate:"required,email"`
	Gender     string `json:"gender" validate:"oneof=male female prefer_not_to"`
	Department string `json:"department" validate:"required"`
	Age        int    `json:"age" validate:"required"`
}
