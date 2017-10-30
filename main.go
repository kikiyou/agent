package main

import (
	"flag"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kikiyou/agent/collector"
	"github.com/kikiyou/agent/g"
	"github.com/kikiyou/agent/shell"
	"github.com/kikiyou/agent/templates"
	cache "github.com/patrickmn/go-cache"
)

var router *gin.Engine

// var authorized gin.HandlerFunc

// VERSION - version
const VERSION = "0.0.3.1"

// shBasicAuthVar - name of env var for basic auth credentials
const shBasicAuthVar = "SH_BASIC_AUTH"

// defaultShellPOSIX - shell executable by default in POSIX systems
const defaultShellPOSIX = "sh"

// defaultShellWindows - shell executable by default in Windows
const defaultShellWindows = "cmd"

// defaultShellPlan9 - shell executable by default in Plan9
const defaultShellPlan9 = "rc"

// Config - config struct
type Config struct {
	port          int    // server port
	cache         int    // caching command out (in seconds)
	timeout       int    // timeout for shell command (in seconds)
	host          string // server host
	exportVars    string // list of environment vars for export to script
	shell         string // custom shell
	defaultShell  string // shell by default
	defaultShOpt  string // shell option for one-liner (-c or /C)
	cert          string // SSL certificate
	key           string // SSL private key path
	authUser      string // basic authentication user name
	authPass      string // basic authentication password
	exportAllVars bool   // export all current environment vars
	setCGI        bool   // set CGI variables
	setForm       bool   // parse form from URL
	noIndex       bool   // dont generate index page
	addExit       bool   // add /exit command
	oneThread     bool   // run each shell commands in one thread
	showErrors    bool   // returns the standard output even if the command exits with a non-zero exit code
	includeStderr bool   // also returns output written to stderr (default is stdout only)
}

var appConfig Config
var CacheTTL *cache.Cache
var (
	// configFile        = flag.String("config", "node_exporter.conf", "config file.")
	// memProfile        = flag.String("memprofile", "", "write memory profile to this file")
	basicAuth        = flag.String("basic-auth", "admin:admin", "setup HTTP Basic Authentication (\"user_name:password\")")
	showVersion      = flag.Bool("version", false, "Print version information.")
	listeningAddress = flag.String("listen", ":8899", "address to listen on")
	// enabledCollectors = flag.String("enabledCollectors", "user_accounts,disk_partitions", "comma-seperated list of collectors to use")
	enabledCollectors = flag.String("enabledCollectors", "current_ram,load_avg", "comma-seperated list of collectors to use")
	publicSharePath   = flag.String("public", filepath.Join(g.Root, "public"), "public share dir")
	printCollectors   = flag.Bool("printCollectors", false, "If true, print available collectors and exit")
)

func loadCollectors() (map[string]collector.Collector, error) {
	collectors := map[string]collector.Collector{}

	for _, name := range strings.Split(*enabledCollectors, ",") {
		fn, ok := collector.Factories[name]
		if !ok {
			log.Fatalf("Collector '%s' not available", name)
		}
		c, err := fn()
		if err != nil {
			return nil, err
		}
		collectors[name] = c
	}
	return collectors, nil
}

// getConfig - parse arguments
func getConfig() (appConfig Config, err error) {
	// var (
	// 	// logFilename string
	// 	basicAuth string
	// )

	if *basicAuth == "" && len(os.Getenv(shBasicAuthVar)) > 0 {
		*basicAuth = os.Getenv(shBasicAuthVar)
	}

	if len(*basicAuth) > 0 {
		basicAuthParts := strings.SplitN(*basicAuth, ":", 2)
		if len(basicAuthParts) != 2 {
			return Config{}, fmt.Errorf("HTTP basic authentication must be in format: name:password, got: %s", basicAuth)
		}
		appConfig.authUser, appConfig.authPass = basicAuthParts[0], basicAuthParts[1]
	}

	// if appConfig.shell != "" && appConfig.shell != appConfig.defaultShell {
	// 	if _, err := exec.LookPath(appConfig.shell); err != nil {
	// 		return Config{}, fmt.Errorf("an error has occurred while searching for shell executable %q: %s", appConfig.shell, err)
	// 	}
	// }

	// // need >= 2 arguments and count of it must be even
	// args := flag.Args()
	// if len(args) < 2 || len(args)%2 == 1 {
	// 	return Config{}, fmt.Errorf("requires a pair of path and shell command")
	// }
	// fmt.Println("-------------------")
	// fmt.Println(args)
	// for i := 0; i < len(args); i += 2 {
	// 	path, cmd := args[i], args[i+1]
	// 	if path[0] != '/' {
	// 		return Config{}, fmt.Errorf("the path %q does not begin with the prefix /", path)
	// 	}
	// 	cmdHandlers = append(cmdHandlers, Command{path: path, cmd: cmd})
	// }

	return appConfig, nil
}

func main() {
	// Set Gin to production mode
	// gin.SetMode(gin.ReleaseMode)

	flag.Parse()
	g.PublicPath = *publicSharePath
	if *showVersion {
		fmt.Fprintln(os.Stdout, VERSION)
		os.Exit(0)
	}
	log.Println("Starting version", VERSION)
	if *printCollectors {
		log.Printf("Available collectors:\n")
		for n, _ := range collector.Factories {
			log.Printf(" - %s\n", n)
		}
		return
	}
	appConfig, _ = getConfig()

	collectors, err := loadCollectors()
	if err != nil {
		log.Fatalf("Couldn't load config and collectors: %s", err)
	}
	log.Printf("Enabled collectors:")
	for n, _ := range collectors {
		log.Printf(" - %s", n)
	}

	g.Collectors = collectors

	// Process the templates at the start so that they don't have to be loaded

	// var tmpl *Template
	// if t == nil {
	// 	t = New(name)
	// }
	// if name == t.Name() {
	// 	tmpl = t
	// } else {
	// 	tmpl = t.New(name)
	// }
	// _, err = tmpl.Parse(s)
	// if err != nil {
	// 	return nil, err
	// }

	// from the disk again. This makes serving HTML pages very fast.
	bytes, err := templates.Asset("templates/index.html") // 根据地址获取对应内容
	if err != nil {
		log.Println(err)
		// return
	}
	t, err := template.New("index.html").Parse(string(bytes)) // 比如用于模板处理

	bytes, err = templates.Asset("templates/command.html") // 根据地址获取对应内容
	if err != nil {
		log.Println(err)
		// return
	}
	t, err = t.New("command.html").Parse(string(bytes))

	// bytes, err = templates.Asset("templates/upload.html") // 根据地址获取对应内容
	// if err != nil {
	// 	log.Println(err)
	// 	// return
	// }
	// t, err = t.New("upload.html").Parse(string(bytes))

	bytes, err = shell.Asset("shell/linux_json_api.sh") // 根据地址获取对应内容
	if err != nil {
		log.Println(err)
		// return "", "xx"
	}

	g.InitTempScriptsFile()
	// fmt.Printf("g.TempScriptsFile:%s", g.TempScriptsFile)
	if err := ioutil.WriteFile(g.TempScriptsFile, bytes, 0700); err != nil {
		g.CheckErr(err)
	}

	// Set the router as the default one provided by Gin
	router = gin.Default()
	// authorized = gin.BasicAuth(gin.Accounts{
	// 	"admin": "admin",
	// })
	router.SetHTMLTemplate(t)
	// router.StaticFS("/static", assetFS())

	// 添加测试模板
	router.LoadHTMLFiles("templates/upload.html")
	router.Static("/static", "./static")
	// fmt.Println("##############")
	// fmt.Println(*publicSharePath)
	_publicDir := *publicSharePath
	fmt.Println(_publicDir)
	if _, err := os.Stat(_publicDir); os.IsNotExist(err) {
		os.MkdirAll(_publicDir, os.ModePerm)
	}
	router.StaticFS("/public", http.Dir(_publicDir))

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	//
	// var CacheTTL *cache.Cache
	appConfig.cache = 1
	if appConfig.cache > 0 {
		CacheTTL = cache.New(5*time.Minute, 10*time.Minute)
	}
	// Initialize the routes
	initializeRoutes()

	router.Run(*listeningAddress)
}

// Render one of HTML, JSON or CSV based on the 'Accept' header of the request
// If the header doesn't specify this, HTML is rendered, provided that
// the template name is present
func render(c *gin.Context, data gin.H, templateName string) {
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
