// routes.go

package main

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/kikiyou/agent/controllers"
	"github.com/kikiyou/agent/g"
	"github.com/kikiyou/agent/models"
	// _ "github.com/mattn/go-sqlite3"
)

var CommandsModel = new(models.COMMANDS)

//CORSMiddleware ...
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "X-Requested-With, Content-Type, Origin, Authorization, Accept, Client-Security-Token, Accept-Encoding, x-access-token")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			fmt.Println("OPTIONS")
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	}
}

func initializeRoutes() {

	// 防止跨站攻击
	router.Use(CORSMiddleware())

	store := sessions.NewCookieStore([]byte("secret"))
	router.Use(sessions.Sessions("fsv_agent", store))

	// Use the setUserStatus middleware for every route to set a flag
	// indicating whether the request was from an authenticated user or not
	router.Use(setUserStatus())

	// 简单认证
	// authorized := gin.BasicAuth(gin.Accounts{
	// 	AppConfig.authUser: AppConfig.authPass,
	// })
	u := router.Group("/u")
	{
		user := new(controllers.UserController)
		u.GET("/login", ensureNotLoggedIn(), user.ShowLoginPage)
		u.POST("/login", user.Login)
		u.GET("/logout", user.Logout)
	}
	router.GET("/", ensureLoggedIn(), func(c *gin.Context) {
		session := sessions.Default(c)
		user_name := session.Get("user_name")
		user_nameStr, _ := user_name.(string)
		var cli bool
		if user_name == "admin" {
			cli = true
		}
		g.Render(c, gin.H{"cli": cli, "user_name": user_nameStr}, "dash.html")
	})
	command := new(controllers.CommandController)

	router.GET("/server", ensureLoggedIn(), command.ModulesRoutes)

	router.GET("/upload", ensureLoggedIn(), func(c *gin.Context) {
		session := sessions.Default(c)
		user_name := session.Get("user_name")
		user_nameStr, _ := user_name.(string)
		var cli bool
		if user_name == "admin" {
			cli = true
		}
		g.Render(c, gin.H{"cli": cli, "user_name": user_nameStr}, "upload.html")
	})
	router.POST("/upload", func(c *gin.Context) {
		// single file
		form, _ := c.MultipartForm()
		// log.Println(form)
		files := form.File["files"]
		for _, file := range files {
			log.Println(file.Filename)
			dst := filepath.Join(g.AppConfig.PublicDir, file.Filename)
			if err := c.SaveUploadedFile(file, dst); err != nil {
				c.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
				return
			}
		}
		c.JSON(http.StatusOK, gin.H{"data": "Uploaded successfully"})
	})
	//download GET 请求输入页面
	//download post请求真的下载 GET 请求输入页面
	//command get 命令 post请求 真的执行

	router.GET("/cli", ensureLoggedIn(), func(c *gin.Context) {
		// result := "rrr"
		// c.Header("Content-Type", "text/html; charset=utf-8")
		// fmt.Println(g.AppConfig.PublicDir)
		CommandList, _ := CommandsModel.GetCommandList()

		g.Render(c, gin.H{"defaultPath": g.AppConfig.PublicDir, "token": g.GenerateToken(), "CommandList": CommandList}, "command.html")
		// c.String(http.StatusOK, result)
	})

	//设置了个2s的容错cache 两秒内同一个命令，只输出一次的结果
	router.POST("/command", command.Command)

}
