package collector

import (
	"fmt"
	"log"
	"net/http"

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
	for k, _ := range Factories {

		if module == k {
			println("#####cccc#####")
			println(k)

			fn, ok := Factories[module]
			// config := &Config{}
			if !ok {
				log.Printf("Collector '%s' not available", module)
			}
			cf, err := fn()
			fmt.Println(err)
			// cf.Update()
			cc, _ := cf.Update()
			fmt.Println(cc)
			// c.String(http.StatusOK, "Hello")
			c.JSON(http.StatusOK, cc)
			println("#####vvv####")
		}
		// println(k)
	}
}

// func RenderJson(v interface{}) {
// 	bs, err := json.Marshal(v)
// 	c.JSON(http.StatusOK, bs)
// }
