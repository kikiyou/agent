// routes.go

package main

import (
	"fmt"
	"net/http"
	"os/exec"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kikiyou/agent/g"
)

func execScriptsGetJSON(module string) (string, error) {
	var out string
	var err error

	// defer f.Close()
	cmd := exec.Command(g.TempScriptsFile, module)

	outCh := make(chan []byte, 1)
	errCh := make(chan error, 1)

	// var out string
	// var err error

	go func() {
		_out, _err := cmd.Output()
		if _err != nil {
			errCh <- fmt.Errorf("failed to collect shell data: %s", _err)
			return
		}
		outCh <- _out
	}()

WAIT:
	for {
		select {
		case res := <-outCh:
			if res != nil {
				out = string(res)
			} else {
				out, err = "nil", fmt.Errorf("failed to collect shell data")
			}
			break WAIT
		case err = <-errCh:
			out = "nil"
			break WAIT
		case <-time.After(10 * time.Second):
			// Kill the process if it takes too long
			if killErr := cmd.Process.Kill(); killErr != nil {
				fmt.Printf("%s timeout,failed to kill:%s", module, killErr)
				return "module timeout", err
				// Force goroutine to exit
				<-outCh
			}
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
