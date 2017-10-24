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

	"github.com/gin-gonic/gin"
	"github.com/kikiyou/agent/collector"
	"github.com/kikiyou/agent/g"
	"github.com/kikiyou/agent/shell"
	"github.com/kikiyou/agent/templates"
)

var router *gin.Engine

// func init (
// 	defaultPublicDir = filepath.Join(g.Root, "public")

// )
var (
	// configFile        = flag.String("config", "node_exporter.conf", "config file.")
	// memProfile        = flag.String("memprofile", "", "write memory profile to this file")
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

func main() {
	// Set Gin to production mode
	// gin.SetMode(gin.ReleaseMode)

	flag.Parse()
	//init
	// g.InitRootDir()
	//

	if *printCollectors {
		log.Printf("Available collectors:\n")
		for n, _ := range collector.Factories {
			log.Printf(" - %s\n", n)
		}
		return
	}
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
	// from the disk again. This makes serving HTML pages very fast.
	bytes, err := templates.Asset("templates/index.html") // 根据地址获取对应内容
	if err != nil {
		log.Println(err)
		// return
	}
	t, err := template.New("index.html").Parse(string(bytes)) // 比如用于模板处理

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
	router.SetHTMLTemplate(t)
	router.StaticFS("/static", assetFS())
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
