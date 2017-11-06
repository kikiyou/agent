package main

import (
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/kikiyou/agent/collector"
	"github.com/kikiyou/agent/g"
	"github.com/kikiyou/agent/shell"
	"github.com/kikiyou/agent/templates"
)

var router *gin.Engine

// var authorized gin.HandlerFunc

func loadCollectors(appConfig g.Config) (map[string]collector.Collector, error) {
	collectors := map[string]collector.Collector{}

	for _, name := range strings.Split(g.AppConfig.EnabledCollectors, ",") {
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

// func tadd(t *template.Template, files ...string) {

// }
func main() {
	// Set Gin to production mode
	// gin.SetMode(gin.ReleaseMode)

	collectors, err := loadCollectors(g.AppConfig)
	if err != nil {
		log.Fatalf("Couldn't load config and collectors: %s", err)
	}
	log.Printf("Enabled collectors:")
	for n, _ := range collectors {
		log.Printf(" - %s", n)
	}

	g.Collectors = collectors

	// from the disk again. This makes serving HTML pages very fast.
	bytes, err := templates.Asset("templates/webshell.tmpl") // 根据地址获取对应内容
	if err != nil {
		log.Println(err)
		// return
	}
	t, err := template.New("webshell.html").Parse(string(bytes)) // 比如用于模板处理

	bytes, err = templates.Asset("templates/dash.tmpl") // 根据地址获取对应内容
	if err != nil {
		log.Println(err)
		// return
	}
	t, err = t.New("dash.html").Parse(string(bytes)) // 比如用于模板处理

	// bytes, err = templates.Asset("templates/command.html") // 根据地址获取对应内容
	// if err != nil {
	// 	log.Println(err)
	// 	// return
	// }
	// t, err = t.New("command.html").Parse(string(bytes))

	bytes, err = templates.Asset("templates/upload.tmpl") // 根据地址获取对应内容
	if err != nil {
		log.Println(err)
		// return
	}
	t, err = t.New("upload.html").Parse(string(bytes))

	bytes, err = templates.Asset("templates/login.tmpl") // 根据地址获取对应内容
	if err != nil {
		log.Println(err)
		// return
	}
	t, err = t.New("login.html").Parse(string(bytes))

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

	// 添加测试模板
	// router.LoadHTMLFiles("templates/dash.html", "templates/upload.html", "templates/login.html", "templates/webshell.html")
	// router.LoadHTMLGlob("./templates/*")
	// router.Static("/static", "./static")

	// _publicDir := *publicSharePath
	// fmt.Println(_publicDir)
	if _, err := os.Stat(g.AppConfig.PublicDir); os.IsNotExist(err) {
		os.MkdirAll(g.AppConfig.PublicDir, os.ModePerm)
	}
	router.StaticFS("/public", http.Dir(g.AppConfig.PublicDir))

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	//

	// Initialize the routes
	initializeRoutes()

	// m := melody.New()

	router.Run(g.AppConfig.ListeningAddress)
}
