// routes.go

package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/kikiyou/agent/controllers"
	"github.com/kikiyou/agent/g"
	shellwords "github.com/mattn/go-shellwords"
	_ "github.com/mattn/go-sqlite3"
	cache "github.com/patrickmn/go-cache"
)

func execScriptsGetJSON(module string) (string, error) {
	var out string
	var err error
	cmd := fmt.Sprintf("%s %s", g.TempScriptsFile, module)
	g.AppConfig.Cache = 2
	path := "/tmp"
	shell, params, err := getShellAndParams(cmd, g.AppConfig)
	if err != nil {
		return "", err
	}
	out, err = execCommand(g.AppConfig, path, shell, params, CacheTTL)
	return out, err
}

func ModulesRoutes(c *gin.Context) {
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

// getShellAndParams - get default shell and command
func getShellAndParams(cmd string, appConfig g.Config) (shell string, params []string, err error) {
	shell, params = appConfig.DefaultShell, []string{appConfig.DefaultShOpt, cmd} // sh -c "cmd"

	// custom shell
	switch {
	case appConfig.Shell != appConfig.DefaultShell && appConfig.Shell != "":
		shell = appConfig.Shell
	case appConfig.Shell == "":
		cmdLine, err := shellwords.Parse(cmd)
		if err != nil {
			return shell, params, fmt.Errorf("Parse '%s' failed: %s", cmd, err)
		}

		shell, params = cmdLine[0], cmdLine[1:]
	}

	return shell, params, nil
}

// execCommand - execute shell command, returns bytes out and error
func execCommand(appConfig g.Config, path string, shell string, params []string, cacheTTL *cache.Cache) (string, error) {
	var (
		out string
		err error
	)
	if path == "" {
		path = g.AppConfig.PublicDir
	}
	// AppConfig.cache = 1
	// log.Println("###############################\n")
	// log.Println(AppConfig.cache)
	fingerStr := fmt.Sprintln(path, shell, strings.Join(params[:], ","))
	fingerPrint := g.MD5(fingerStr)

	if g.AppConfig.Cache > 0 {
		if cacheData, found := cacheTTL.Get(fingerPrint); !found {
			// log.Printf("get from cache failed: %s", err)
			// log.Println("no cache")
		} else if found {
			// cache hit
			log.Println("cache hit %s", fingerStr)
			out, _ = cacheData.(string)
			// out, _ = fmt.Fprintln(os.Stdout, cacheData)
			return out, nil
		}
	}
	cmd := exec.Command(shell, params...)
	cmd.Dir = path
	// Use a bytes.Buffer to get the output
	var buf bytes.Buffer
	cmd.Stdout = &buf

	cmd.Start()

	// Use a channel to signal completion so we can use a select statement
	done := make(chan error)
	go func() { done <- cmd.Wait() }()

	// Start a timer
	timeout := time.After(20 * time.Minute)

	// The select statement allows us to execute based on which channel
	// we get a message from first.
	select {
	case <-timeout:
		// Timeout happened first, kill the process and print a message.
		cmd.Process.Kill()
		log.Printf("%s module: timeout,fail to killed", shell)
		out = fmt.Sprintf("%s module: timeout,fail to killed", shell)
	case err := <-done:
		// Command completed before timeout. Print output and error if it exists.
		out = buf.String()
		if err != nil {
			log.Printf("%s modele: Non-zero exit code:%s", shell, err)
			out = fmt.Sprintf("%s modele: Non-zero exit code:%s", shell, err)
		}
	}
	if g.AppConfig.Cache > 0 {
		cacheTTL.Set(fingerPrint, out, time.Duration(g.AppConfig.Cache)*time.Second)
	}
	return out, err
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
		u.POST("/login", ensureNotLoggedIn(), user.Login)
	}
	// Handle the index route
	// router.GET("/", ensureLoggedIn(), showIndexPage)
	router.GET("/", ensureLoggedIn(), func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", "")
	})
	router.GET("/index", ensureLoggedIn(), func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", "")
	})
	router.GET("/dash", ensureLoggedIn(), func(c *gin.Context) {
		c.HTML(http.StatusOK, "dash.html", "")
	})
	router.GET("/server", ensureLoggedIn(), ModulesRoutes)

	router.GET("/upload", ensureLoggedIn(), func(c *gin.Context) {
		g.Render(c, gin.H{"title": "Create New Article"}, "upload.html")
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
		c.Header("Content-Type", "text/html; charset=utf-8")
		fmt.Println(g.AppConfig.PublicDir)
		g.Render(c, gin.H{"defaultPath": g.AppConfig.PublicDir}, "command.html")
		// c.String(http.StatusOK, result)
	})
	// router.GET("/command/:commandID", func(c *gin.Context) {
	// 	name := c.Param("name")
	// 	c.String(http.StatusOK, "Hello %s", name)
	// 	// c.Header("Content-Type", "text/html; charset=utf-8")
	// 	// render(c, gin.H{"defaultPath": g.PublicPath}, "command.html")
	// 	// c.String(http.StatusOK, result)
	// })
	//设置了个2s的容错cache 两秒内同一个命令，只输出一次的结果
	router.POST("/command", func(c *gin.Context) {
		var (
			shellOut string
			path     string
			cmd      string
		)

		if command, ok := c.GetPostForm("command"); ok {
			if r, ok := c.GetPostForm("path"); ok {
				path = r
			}
			cmd = command
		}
		if commandID, ok := c.GetPostForm("commandID"); ok {
			fmt.Println(commandID)
			//查询数据
			// db, err := sql.Open("sqlite3", "db/command_set.sqlite3")
			// g.CheckErr(err)

			// rows, err := db.Query("SELECT * FROM COMMANDS")
			// g.CheckErr(err)

			// for rows.Next() {
			// 	var ID int
			// 	var COMMAND string
			// 	var LABEL string
			// 	var ISDYNAMIC int
			// 	err = rows.Scan(&ID, &COMMAND, &LABEL, &ISDYNAMIC)
			// 	g.CheckErr(err)
			// 	fmt.Println(ID)
			// 	fmt.Println(COMMAND)
			// 	fmt.Println(LABEL)
			// 	fmt.Println(ISDYNAMIC)
			// }
			cmd = "ls -l /"

		}
		if cmd != "" {
			g.AppConfig.Shell = "sh"
			g.AppConfig.DefaultShOpt = "-c"
			shell, params, err := getShellAndParams(cmd, g.AppConfig)
			if err != nil {
				return
			}
			g.AppConfig.Cache = 2
			shellOut, err = execCommand(g.AppConfig, path, shell, params, CacheTTL)

			if _, ok := c.GetPostForm("html"); ok {
				s := bytes.Replace([]byte(CommandTemplate), []byte("CONTENT"), []byte(shellOut), 1)
				shellOut = string(s)
				c.String(http.StatusOK, shellOut)
			} else {
				c.String(http.StatusOK, shellOut)
			}
		}

	})

}
