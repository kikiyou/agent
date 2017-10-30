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

	"github.com/gin-gonic/gin"
	"github.com/kikiyou/agent/g"
	shellwords "github.com/mattn/go-shellwords"
	cache "github.com/patrickmn/go-cache"
)

func execScriptsGetJSON(module string) (string, error) {
	var out string
	var err error
	cmd := fmt.Sprintf("%s %s", g.TempScriptsFile, module)
	appConfig.cache = 2
	path := "/tmp"
	shell, params, err := getShellAndParams(cmd, appConfig)
	if err != nil {
		return "", err
	}
	out, err = execCommand(appConfig, path, shell, params, CacheTTL)
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
func getShellAndParams(cmd string, appConfig Config) (shell string, params []string, err error) {
	shell, params = appConfig.defaultShell, []string{appConfig.defaultShOpt, cmd} // sh -c "cmd"

	// custom shell
	switch {
	case appConfig.shell != appConfig.defaultShell && appConfig.shell != "":
		shell = appConfig.shell
	case appConfig.shell == "":
		cmdLine, err := shellwords.Parse(cmd)
		if err != nil {
			return shell, params, fmt.Errorf("Parse '%s' failed: %s", cmd, err)
		}

		shell, params = cmdLine[0], cmdLine[1:]
	}

	return shell, params, nil
}

// execCommand - execute shell command, returns bytes out and error
func execCommand(appConfig Config, path string, shell string, params []string, cacheTTL *cache.Cache) (string, error) {
	var (
		out string
		err error
	)
	if path == "" {
		path = g.PublicPath
	}
	// appConfig.cache = 1
	// log.Println("###############################\n")
	// log.Println(appConfig.cache)
	fingerStr := fmt.Sprintln(path, shell, strings.Join(params[:], ","))
	fingerPrint := g.MD5(fingerStr)

	if appConfig.cache > 0 {
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
	timeout := time.After(10 * time.Second)

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
	if appConfig.cache > 0 {
		cacheTTL.Set(fingerPrint, out, time.Duration(appConfig.cache)*time.Second)
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

func initializeRoutes() {

	// Use the setUserStatus middleware for every route to set a flag
	// indicating whether the request was from an authenticated user or not
	router.Use(setUserStatus())

	// Handle the index route
	authorized := gin.BasicAuth(gin.Accounts{
		appConfig.authUser: appConfig.authPass,
	})
	// router.GET("/", ensureLoggedIn(), showIndexPage)
	router.GET("/", authorized, showIndexPage)
	router.GET("/server", ModulesRoutes)

	// Set a lower memory limit for multipart forms (default is 32 MiB)
	// router.MaxMultipartMemory = 8 << 20  // 8 MiB
	// router.GET("/upload", authorized, func(c *gin.Context) {
	// 	result := `<html><body><form method=POST action=/upload enctype=multipart/form-data><input type=file name=file><input type=submit></form>`
	// 	c.Header("Content-Type", "text/html; charset=utf-8")
	// 	c.String(http.StatusOK, result)
	// })
	router.GET("/upload", authorized, func(c *gin.Context) {
		render(c, gin.H{"title": "Create New Article"}, "upload.html")
	})
	router.POST("/upload", func(c *gin.Context) {
		// single file
		form, _ := c.MultipartForm()
		log.Println(form)
		files := form.File["files"]
		for _, file := range files {
			log.Println(file.Filename)
			dst := filepath.Join(*publicSharePath, file.Filename)
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
	router.GET("/command", authorized, func(c *gin.Context) {
		// result := "rrr"
		c.Header("Content-Type", "text/html; charset=utf-8")
		render(c, gin.H{"defaultPath": g.PublicPath}, "command.html")
		// c.String(http.StatusOK, result)
	})
	//设置了个2s的容错cache 两秒内同一个命令，只输出一次的结果
	router.POST("/command", func(c *gin.Context) {
		if cmd, ok := c.GetPostForm("command"); ok {
			path := ""
			if r, ok := c.GetPostForm("path"); ok {
				path = r
			}
			// fmt.Println(cmd)
			appConfig.shell = "sh"
			appConfig.defaultShOpt = "-c"
			shell, params, err := getShellAndParams(cmd, appConfig)
			if err != nil {
				return
			}
			// fmt.Printf("shell->%s,params-%s", shell, params)
			appConfig.cache = 2
			shellOut, err := execCommand(appConfig, path, shell, params, CacheTTL)
			// fmt.Println(shellOut)
			// terminal.wrapPreview()
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
