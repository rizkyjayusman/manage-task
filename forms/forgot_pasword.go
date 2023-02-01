package forms

type ForgotPassword struct {
	Email string `form:"email" binding:"required,email"`
}
