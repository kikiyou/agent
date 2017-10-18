package collector

import (
	"bytes"
	"fmt"
	"net/http"
	"os/exec"

	"github.com/gin-gonic/gin"
)

type Collector interface {
	// Get new metrics and expose them via prometheus registry.
	Update() (res interface{}, err error)
}

var Factories = make(map[string]func() (Collector, error))

const Namespace = "node"

// Interface a collector has to implement.

// TODO: Instead of periodically call Update, a Collector could be implemented
// as a real prometheus.Collector that only gathers metrics when
// scraped. (However, for metric gathering that takes very long, it might
// actually be better to do them proactively before scraping to minimize scrape
// time.)

type Config struct {
	Config     map[string]string `json:"config"`
	Attributes map[string]string `json:"attributes"`
}

func ModulesRoutes(c *gin.Context) {
	module := c.Query("module")
	// fmt.Println(module)
	if module == "" {
		// fmt.Println("cccc1")
		c.JSON(http.StatusBadRequest, gin.H{"error": "No module specified, or requested module doesn't exist."})
		return
	}
	if fn, ok := Factories[module]; ok {
		cf, _ := fn()
		// fmt.Println(err)
		// cf.Update()
		cc, _ := cf.Update()
		// fmt.Println(cc)
		// c.String(http.StatusOK, "Hello")
		c.JSON(http.StatusOK, cc)
		// println("#####vvv####")
		// println("#####vvv####")

	} else {
		fmt.Println(module)
		cmd := exec.Command("./linux_json_api.sh", module)
		var output bytes.Buffer
		cmd.Stdout = &output
		err := cmd.Run()
		if err != nil {
			fmt.Printf("Error executing '%s': %s\n\tScript output: %s\n", module, err.Error(), output.String())
			c.JSON(http.StatusBadRequest, gin.H{"error": "Unable to execute module"})
			return
		}
		// fmt.Println(module)
		// fmt.Println(output.String())
		c.String(http.StatusOK, output.String())
	}

}

// func RenderJson(v interface{}) {
// 	bs, err := json.Marshal(v)
// 	c.JSON(http.StatusOK, bs)
// }
