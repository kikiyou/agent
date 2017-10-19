package collector

import (
	"encoding/json"
	"log"
	"runtime"

	"github.com/shirou/gopsutil/load"
)

func init() {
	Factories["load_avg"] = NewCollector_laod_avg
}

type laod_avg struct {
	Load1  float64 `json:"1_min_avg"`
	Load5  float64 `json:"5_min_avg"`
	Load15 float64 `json:"15_min_avg"`
}

// type laod_avg []laod_avg

func (m laod_avg) String() string {
	s, _ := json.Marshal(m)
	return string(s)
}
func NewCollector_laod_avg() (Collector, error) {
	var c *laod_avg = new(laod_avg)
	return c, nil
}

// func check(err error) {
// 	if err != nil {
// 		panic(err)
// 	}
// }

func (c *laod_avg) Update() (res interface{}, err error) {
	// log.Println("Update-----------------\n")

	loads, err := load.Avg()
	if err != nil {
		log.Fatalln("load_avg:,", err)
	}
	// var NumCPU float64
	NumCPU := runtime.NumCPU()

	return laod_avg{
		Load1:  loads.Load1 * 100 / float64(NumCPU),
		Load5:  loads.Load5 * 100 / float64(NumCPU),
		Load15: loads.Load15 * 100 / float64(NumCPU),
	}, nil
}
