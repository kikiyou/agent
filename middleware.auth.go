// middleware.auth.go

package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/kikiyou/agent/g"
)

func ensureLoggedIn() gin.HandlerFunc {
	return func(c *gin.Context) {
		// If there's an error or if the token is empty
		// the user is not logged in
		loggedInInterface, _ := c.Get("is_logged_in")
		loggedIn := loggedInInterface.(bool)
		fmt.Printf("ensureLoggedIn %s", loggedIn)
		if !loggedIn {
			//if token, err := c.Cookie("token"); err != nil || token == "" {
			g.Render(c, gin.H{
				"title": "Login",
			}, "login.html")
			// c.AbortWithStatus(http.StatusUnauthorized)
			// c.HTML(http.StatusOK, "index.html", "")
			// fmt.Println("cccc")
			c.Abort()
		}
	}
}

// This middleware ensures that a request will be aborted with an error
// if the user is already logged in
func ensureNotLoggedIn() gin.HandlerFunc {
	return func(c *gin.Context) {
		// If there's no error or if the token is not empty
		// the user is already logged in
		loggedInInterface, _ := c.Get("is_logged_in")
		loggedIn := loggedInInterface.(bool)
		fmt.Printf("ensureNotLoggedIn %s", loggedIn)
		if loggedIn {
			// c.HTML(http.StatusOK, "index.html", "")
			// c.Abort()
			c.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}

// This middleware sets whether the user is logged in or not
func setUserStatus() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		if session.Get("user_name") != nil {
			c.Set("is_logged_in", true)
		} else {
			c.Set("is_logged_in", false)
		}
	}
}
