package forms

//LoginForm ...
type LoginForm struct {
	Name     string `form:"name" json:"name" binding:"required,max=100"`
	Password string `form:"password" json:"password" binding:"required"`
}

// //SignupForm ...
// type LogupForm struct {
// 	Name     string `form:"name" json:"name" binding:"required,max=100"`
// 	Password string `form:"password" json:"password" binding:"required"`
// }
