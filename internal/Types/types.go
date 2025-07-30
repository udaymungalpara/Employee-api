package types

type Employee struct {
	Id         int
	Name       string `validate:"required"`
	Email      string `validate:"required,email"`
	Gender     string `validate:"oneof=male female prefer_not_to"`
	Department string `validate:"required"`
	Age        int    `validate:"required"`
}
