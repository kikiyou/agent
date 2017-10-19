package collector

type Collector interface {
	// Get new metrics and expose them via prometheus registry.
	Update() (res interface{}, err error)
}

var Factories = make(map[string]func() (Collector, error))

// const Namespace = "node"

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

// func RenderJson(v interface{}) {
// 	bs, err := json.Marshal(v)
// 	c.JSON(http.StatusOK, bs)
// }
