package form

type FormLogin struct {
	Name string `form:"name" binding:"required"`
	Password string `form:"password" binding:"required"`
}