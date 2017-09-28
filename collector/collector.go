package collector

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	// "github.com/kikiyou/agent/collector"
	// "github.com/kikiyou/agent/collector"
)

type Collector interface {
	// Get new metrics and expose them via prometheus registry.
	Update() (res interface{}, err error)
}

// // // Implements Collector.
// type NodeCollector struct {
// 	collectors map[string]Collector
// }

var Factories = make(map[string]func(Config) (Collector, error))

// Implements Collector.
// func Collect(ch chan<- prometheus.Metric) {
// 	wg := sync.WaitGroup{}
// 	wg.Add(len(Factories))
// 	for name, c := range Factories {
// 		c.Update(ch)
// 		// go func(name string, c Collector) {
// 		// 	// Execute(name, c, ch)
// 		// 	wg.Done()
// 		// }(name, c)
// 	}
// 	wg.Wait()
// 	// scrapeDurations.Collect(ch)
// }

// func Execute(name string, c Collector, ch chan<- prometheus.Metric) {
// 	// begin := time.Now()
// 	err := c.Update(ch)
// 	// duration := time.Since(begin)
// 	var result string

// 	// if err != nil {
// 	// 	log.Printf("ERROR: %s failed after %fs: %s", name, duration.Seconds(), err)
// 	// 	result = "error"
// 	// } else {
// 	// 	log.Printf("OK: %s success after %fs.", name, duration.Seconds())
// 	// 	result = "success"
// 	// }
// 	// // scrapeDurations.WithLabelValues(name, result).Observe(duration.Seconds())
// }

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
			config := &Config{}
			if !ok {
				log.Printf("Collector '%s' not available", module)
			}
			cf, err := fn(*config)
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
