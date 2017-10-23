// routes.go

package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kikiyou/agent/g"
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
	timeout := time.After(2 * time.Second)

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

func initializeRoutes() {

	// Use the setUserStatus middleware for every route to set a flag
	// indicating whether the request was from an authenticated user or not
	router.Use(setUserStatus())

	// Handle the index route
	// router.GET("/", ensureLoggedIn(), showIndexPage)
	router.GET("/", showIndexPage)
	router.GET("/server", ModulesRoutes)
}
