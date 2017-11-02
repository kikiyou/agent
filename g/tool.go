package g

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	shellwords "github.com/mattn/go-shellwords"
	cache "github.com/patrickmn/go-cache"
	// "html/template"
)

// 生成32位MD5
func MD5(text string) string {
	ctx := md5.New()
	ctx.Write([]byte(text))
	return hex.EncodeToString(ctx.Sum(nil))
}

// Render one of HTML, JSON or CSV based on the 'Accept' header of the request
// If the header doesn't specify this, HTML is rendered, provided that
// the template name is present
func Render(c *gin.Context, data gin.H, templateName string) {
	loggedInInterface, _ := c.Get("is_logged_in")
	data["is_logged_in"] = loggedInInterface.(bool)
	// fmt.Println("uuuuuuuuuuuuuuu")
	switch c.Request.Header.Get("Accept") {
	case "application/json":
		// Respond with JSON
		c.JSON(http.StatusOK, data["payload"])
	case "application/xml":
		// Respond with XML
		c.XML(http.StatusOK, data["payload"])
	default:
		// Respond with HTML
		c.HTML(http.StatusOK, templateName, data)

	}
}

// getShellAndParams - get default shell and command
func GetShellAndParams(cmd string, appConfig Config) (shell string, params []string, err error) {
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
func ExecCommand(appConfig Config, path string, shell string, params []string, cacheTTL *cache.Cache) (string, error) {
	var (
		out string
		err error
	)
	if path == "" {
		path = appConfig.PublicDir
	}
	// AppConfig.cache = 1
	// log.Println("###############################\n")
	// log.Println(AppConfig.cache)
	fingerStr := fmt.Sprintln(path, shell, strings.Join(params[:], ","))
	fingerPrint := MD5(fingerStr)

	if appConfig.Cache > 0 {
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
	if appConfig.Cache > 0 {
		cacheTTL.Set(fingerPrint, out, time.Duration(appConfig.Cache)*time.Second)
	}
	return out, err
}
