package controllers

import (
	"net/http"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/kikiyou/agent/forms"
	"github.com/kikiyou/agent/g"
	"github.com/kikiyou/agent/models"
)

type UserController struct{}

var userModel = new(models.UserModel)

func (ctrl UserController) ShowLoginPage(c *gin.Context) {
	// Call the render function with the name of the template to render
	g.Render(c, gin.H{
		"title": "Login",
	}, "login.html")
}

//Login ...
func (ctrl UserController) Login(c *gin.Context) {
	var LoginForm forms.LoginForm

	if c.Bind(&LoginForm) != nil {
		c.JSON(406, gin.H{"message": "无效的提交", "form": LoginForm})
		c.Abort()
		return
	}

	user, err := userModel.Login(LoginForm)
	if err == nil {
		c.Set("is_logged_in", true)
		session := sessions.Default(c)
		session.Set("user_id", user.ID)
		session.Set("user_name", user.Name)
		session.Set("user_role", user.Role)
		session.Save()
		// cc := session.Get("user_name")
		// c.SetCookie("token", "token", 3600, "", "", false, true)
		// fmt.Println(cc)
		// fmt.Println(user)
		// c.Set("is_logged_in", true)
		c.HTML(http.StatusOK, "index.html", "")
		// c.RedirectRedirects()
		// c.Redirect(301, "/index")
		// c.JSON(200, gin.H{"message": "User signed in", "user": user})
	} else {
		c.JSON(406, gin.H{"message": "登录失败", "error": err.Error()})
		c.Abort()
		return
	}

}
