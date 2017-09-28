package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/kikiyou/agent/collector"
)

var router *gin.Engine

var (
	// configFile        = flag.String("config", "node_exporter.conf", "config file.")
	// memProfile        = flag.String("memprofile", "", "write memory profile to this file")
	// listeningAddress  = flag.String("listen", ":8080", "address to listen on")
	enabledCollectors = flag.String("enabledCollectors", "user_accounts,current_ram", "comma-seperated list of collectors to use")
	printCollectors   = flag.Bool("printCollectors", true, "If true, print available collectors and exit")



func loadCollectors() (map[string]collector.Collector, error) {
	collectors := map[string]collector.Collector{}

	for _, name := range strings.Split(*enabledCollectors, ",") {
		fn, ok := collector.Factories[name]
		// config := &collector.Config{}
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
	flag.Parse()
	// printCollectors = "888"
	if *printCollectors {
		fmt.Printf("Available collectors:\n")
		for n, _ := range collector.Factories {
			fmt.Printf(" - %s\n", n)
		}
		// return
	}
	// Set Gin to production mode
	// gin.SetMode(gin.ReleaseMode)

	// Set the router as the default one provided by Gin
	router = gin.Default()

	// Process the templates at the start so that they don't have to be loaded
	// from the disk again. This makes serving HTML pages very fast.
	router.LoadHTMLGlob("templates/*")
	router.Static("/static", "static")

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
	fmt.Println("ccccccccccccccccccccc")
	fmt.Println(collectors)
	// MakeNodeCollector := collectors
	// collector.RegisterNodeCollector(n)

	initializeRoutes()

	// Start serving the application
	router.Run(":9090")
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
