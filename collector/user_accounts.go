// +build !noloadavg

package collector

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/prometheus/client_golang/prometheus"
	// "github.com/prometheus/client_golang/prometheus"
)

const (
	procLoad = "/proc/loadavg"
)

var (
	load1 = prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: Namespace,
		Name:      "load1",
		Help:      "1m load average.",
	})
)

type loadavgCollector struct {
	config string
}

func init() {
	Factories["user_accounts"] = NewLoadavgCollector
}

// Takes a config struct and prometheus registry and returns a new Collector exposing
// load, seconds since last login and a list of tags as specified by config.
func NewLoadavgCollector() (Collector, error) {
	// config := &Config{}
	c := loadavgCollector{
		config: "ccc",
	}
	return &c, nil
}

func (c *loadavgCollector) Update() (res interface{}, err error) {
	println("Update-----------------\n")
	load, err := getLoad1()
	if err != nil {
		// return fmt.Errorf("Couldn't get load: %s", err error)
		fmt.Println(err)

	}
	// glog.V(1).Infof("Set node_load: %f", load)
	load1.Set(load)
	// load1.Collect(ch)
	return "000", nil
}

func getLoad1() (float64, error) {
	data, err := ioutil.ReadFile(procLoad)
	if err != nil {
		return 0, err
	}
	return parseLoad(string(data))
}

func parseLoad(data string) (float64, error) {
	parts := strings.Fields(data)
	load, err := strconv.ParseFloat(parts[0], 64)
	if err != nil {
		return 0, fmt.Errorf("Could not parse load '%s': %s", parts[0], err)
	}
	return load, nil
}
