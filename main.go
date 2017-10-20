package main

import (
	"flag"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/kikiyou/agent/shell"
	"github.com/kikiyou/agent/templates"

	"github.com/gin-gonic/gin"
	"github.com/kikiyou/agent/collector"
	"github.com/kikiyou/agent/g"
	// "github.com/prometheus/node_exporter/collector"
)

var router *gin.Engine

var (
	// configFile        = flag.String("config", "node_exporter.conf", "config file.")
	// memProfile        = flag.String("memprofile", "", "write memory profile to this file")
	// listeningAddress  = flag.String("listen", ":8080", "address to listen on")
	// enabledCollectors = flag.String("enabledCollectors", "user_accounts,disk_partitions", "comma-seperated list of collectors to use")
	enabledCollectors = flag.String("enabledCollectors", "current_ram,load_avg", "comma-seperated list of collectors to use")
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

// // Implements Collector.
// type NodeCollector struct {
// 	collectors map[string]collector.Collector
// }
func check(e error) {
	if e != nil {
		panic(e)
	}
}
func main() {
	flag.Parse()
	// printCollectors = "888"
	if *printCollectors {
		log.Printf("Available collectors:\n")
		for n, _ := range collector.Factories {
			log.Printf(" - %s\n", n)
		}
		return
	}
	// Set Gin to production mode
	// gin.SetMode(gin.ReleaseMode)

	// Set the router as the default one provided by Gin
	router = gin.Default()

	// Process the templates at the start so that they don't have to be loaded
	// from the disk again. This makes serving HTML pages very fast.
	// router.LoadHTMLGlob("templates/*")
	bytes, err := shell.Asset("shell/linux_json_api.sh") // 根据地址获取对应内容
	if err != nil {
		fmt.Println(err)
		// return "", "xx"
	}
	err = ioutil.WriteFile("/tmp/linux_json_api.sh", bytes, 0755)
	check(err)
	// -pkg=asset, 打包的包名是 asset
	bytes, err = templates.Asset("templates/index.html") // 根据地址获取对应内容
	if err != nil {
		fmt.Println(err)
		return
	}
	t, err := template.New("index.html").Parse(string(bytes)) // 比如用于模板处理
	router.SetHTMLTemplate(t)
	// fmt.Println(string(bytes))
	// Start serving the application

	router.StaticFS("/static", assetFS())

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	// Initialize the routes
	collectors, err := loadCollectors()
	if err != nil {
		log.Fatalf("Couldn't load config and collectors: %s", err)
	}
	log.Printf("Enabled collectors:")
	for n, _ := range collectors {
		log.Printf(" - %s", n)
	}
	// log.Println("ccccccccccccccccccccc")
	// log.Println(collectors)
	g.Collectors = collectors
	// NNodeCollector := NodeCollector{collectors: collectors}
	// fmt.Println(NNodeCollector)
	// MakeNodeCollector := collectors
	// collector.RegisterNodeCollector(n)

	initializeRoutes()

	router.Run(":8899")
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
