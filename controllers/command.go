package controllers

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kikiyou/agent/forms"
	"github.com/kikiyou/agent/g"
	"github.com/kikiyou/agent/models"
	"golang.org/x/crypto/bcrypt"
)

type CommandController struct{}

// var CacheTTL g.CacheTTL
var CommandsModel = new(models.COMMANDS)

func (ctrl CommandController) Command(c *gin.Context) {
	var (
		shellOut    string
		path        string
		cmd         string
		CommandForm forms.CommandForm
	)
	if c.Bind(&CommandForm) != nil {
		c.JSON(406, gin.H{"message": "无效的提交", "form": CommandForm})
		c.Abort()
		return
	}

	if hashedToken, ok := c.GetPostForm("token"); ok {
		tokenStr := g.GetTokenStr()
		err := bcrypt.CompareHashAndPassword([]byte(hashedToken), []byte(tokenStr))
		if err != nil {
			c.JSON(406, gin.H{"message": "无效的token"})
			c.Abort()
			return
		}

		if command, ok := c.GetPostForm("command"); ok {
			if r, ok := c.GetPostForm("path"); ok {
				path = r
			}
			cmd = command
		}
		if commandID, ok := c.GetPostForm("commandID"); ok {
			fmt.Println(commandID)
			//查询数据
			intCommandID, _ := g.StrToInt64(commandID)
			if err != nil {
				return
			}
			comand, err := CommandsModel.GetCommandAndIsDynamicByID(intCommandID)
			// db, err := sql.Open("sqlite3", "db/command_set.sqlite3")
			if err != nil {
				return
			}
			// cc := CommandsModel.AddCommand("ss", "monkey22")
			// cc := CommandsModel.DeleteCommand(intCommandID)
			// fmt.Println(cc)
			cmd = comand.COMMAND

		}
		if cmd != "" {
			// g.AppConfig.Shell = "sh"
			// g.AppConfig.DefaultShOpt = "-c"
			shell, params, err := g.GetShellAndParams(cmd, g.AppConfig)
			if err != nil {
				return
			}
			// g.AppConfig.Cache = 2
			shellOut, err = g.ExecCommand(g.AppConfig, path, shell, params, g.AppConfig.CacheTTL)

			if _, ok := c.GetPostForm("html"); ok {
				s := bytes.Replace([]byte(CommandTemplate), []byte("CONTENT"), []byte(shellOut), 1)
				shellOut = string(s)
				c.String(http.StatusOK, shellOut)
			} else {
				c.String(http.StatusOK, shellOut)
			}
		}
	}
}

func execScriptsGetJSON(module string) (string, error) {
	var out string
	var err error
	shell := g.TempScriptsFile
	params := []string{module}
	// fmt.Println(params)
	path := "/tmp"
	out, err = g.ExecCommand(g.AppConfig, path, shell, params, g.AppConfig.CacheTTL)
	return out, err
}

func (ctrl CommandController) ModulesRoutes(c *gin.Context) {
	module := c.Query("module")
	if module == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No module specified, or requested module doesn't exist."})
		return
	}
	// fmt.Println(g.Collectors)
	if fn, ok := g.Collectors[module]; ok {
		result, _ := fn.Update()
		c.JSON(http.StatusOK, result)

	} else {
		output, _ := execScriptsGetJSON(module)
		c.String(http.StatusOK, output)
	}

}

var CommandTemplate = `
<!DOCTYPE html>
<html>
	<head>
		<title>terminal-to-html Preview</title>
		<link rel="stylesheet" href="static/css/terminal.css">
	</head>
	<body>
		<div class="term-container">CONTENT</div>
	</body>
</html>
`
