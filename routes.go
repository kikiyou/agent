// routes.go

package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kikiyou/agent/g"
	shellwords "github.com/mattn/go-shellwords"
	cache "github.com/patrickmn/go-cache"
)

func execScriptsGetJSON(module string) (string, error) {
	var out string
	var err error

	// var buf bytes.Buffer
	// cmd.Stdout = &buf

	cmd := exec.Command(g.TempScriptsFile, module)
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
		log.Printf("%s module: timeout,fail to killed", module)
		out = fmt.Sprintf("%s module: timeout,fail to killed", module)
	case err := <-done:
		// Command completed before timeout. Print output and error if it exists.
		out = buf.String()
		if err != nil {
			log.Printf("%s modele: Non-zero exit code:%s", module, err)
			out = fmt.Sprintf("%s modele: Non-zero exit code:%s", module, err)
		}
	}
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
		// fmt.Println("output77")
		// fmt.Println(output)
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

// execShellCommand - execute shell command, returns bytes out and error
func execShellCommand(appConfig Config, shell string, params []string, cacheTTL *cache.Cache) (string, error) {

	var (
		out string
		err error
	)
	// appConfig.cache = 1
	log.Println("###############################\n")
	log.Println(appConfig.cache)
	if appConfig.cache > 0 {
		if cacheData, found := cacheTTL.Get("foo"); !found {
			// log.Printf("get from cache failed: %s", err)
			log.Println("no cache")
		} else if found {
			// cache hit
			log.Println("cache hit %s", cacheData)
			out, _ = cacheData.(string)
			// out, _ = fmt.Fprintln(os.Stdout, cacheData)
			return out, nil
		}
	}

	// ctx := req.Context()
	// if appConfig.timeout > 0 {
	// 	var cancelFn context.CancelFunc
	// 	ctx, cancelFn = context.WithTimeout(ctx, time.Duration(appConfig.timeout)*time.Second)
	// 	defer cancelFn()
	// }
	// osExecCommand := exec.CommandContext(ctx, shell, params...) // #nosec

	// proxySystemEnv(osExecCommand, appConfig)

	// finalizer := func() {}
	// if appConfig.setForm {
	// 	var err error
	// 	if finalizer, err = getForm(osExecCommand, req); err != nil {
	// 		log.Printf("parse form failed: %s", err)
	// 	}
	// }

	// if appConfig.setCGI {
	// 	setCGIEnv(osExecCommand, req, appConfig)

	// 	// get POST data to stdin of script (if not parse form vars above)
	// 	if req.Method == "POST" && !appConfig.setForm {
	// 		if stdin, pipeErr := osExecCommand.StdinPipe(); pipeErr != nil {
	// 			log.Println("write POST data to shell failed:", pipeErr)
	// 		} else {
	// 			waitPipeWrite = true
	// 			go func() {
	// 				if _, pipeErr := io.Copy(stdin, req.Body); pipeErr != nil {
	// 					pipeErrCh <- pipeErr
	// 					return
	// 				}
	// 				pipeErrCh <- stdin.Close()
	// 			}()
	// 		}
	// 	}
	// }

	// if appConfig.includeStderr {
	// 	shellOut, err = osExecCommand.CombinedOutput()
	// } else {
	// 	osExecCommand.Stderr = os.Stderr
	// 	shellOut, err = osExecCommand.Output()
	// }

	// if waitPipeWrite {
	// 	if pipeErr := <-pipeErrCh; pipeErr != nil {
	// 		log.Println("write POST data to shell failed:", pipeErr)
	// 	}
	// }

	// finalizer()

	cmd := exec.Command(shell, params...)
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
		// if cacheErr := cacheTTL.SetBytes(req.RequestURI, shellOut, appConfig.cache); cacheErr != nil {
		// 	log.Printf("set to cache failed: %s", cacheErr)
		// }
		cacheTTL.Set("foo", out, cache.NoExpiration)
	}
	// out = "-"
	return out, err
}

func initializeRoutes() {

	// // Set the value of the key "foo" to "bar", with the default expiration time
	// CacheTTL.Set("foo", "bar", cache.DefaultExpiration)

	// // Set the value of the key "baz" to 42, with no expiration time
	// // (the item won't be removed until it is re-set, or removed using
	// // c.Delete("baz")
	// CacheTTL.Set("baz", 42, cache.NoExpiration)

	// // Get the string associated with the key "foo" from the cache
	// foo, found := CacheTTL.Get("foo")
	// if found {
	// 	fmt.Println(foo)
	// }

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
	router.GET("/upload", authorized, func(c *gin.Context) {
		result := `<html><body><form method=POST action=/upload enctype=multipart/form-data><input type=file name=file><input type=submit></form>`
		c.Header("Content-Type", "text/html; charset=utf-8")
		c.String(http.StatusOK, result)
	})
	router.POST("/upload", func(c *gin.Context) {
		// single file
		file, _ := c.FormFile("file")
		log.Println(file.Filename)

		// Upload the file to specific dst.
		dst := filepath.Join(*publicSharePath, file.Filename)
		c.SaveUploadedFile(file, dst)

		c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
	})
	//download GET 请求输入页面
	//download post请求真的下载 GET 请求输入页面
	//command get 命令 post请求 真的执行
	router.POST("/command", func(c *gin.Context) {
		if cmd, ok := c.GetPostForm("command"); ok {
			fmt.Println(cmd)
			shell, params, err := getShellAndParams(cmd, appConfig)
			if err != nil {
				return
			}
			fmt.Printf("shell->%s,params-%s", shell, params)
			shellOut, err := execShellCommand(appConfig, shell, params, CacheTTL)
			fmt.Println(shellOut)
			// getShellHandler(appConfig, path, shell, params, cacheTTL)
		}
		// c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
	})

}
