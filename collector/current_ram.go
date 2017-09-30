package collector

import (
	"encoding/json"
	"log"

	"github.com/shirou/gopsutil/mem"
)

func init() {
	Factories["current_ram"] = NewCurrentRamCollector
}

type current_ram struct {
	Total     uint64 `json:"total"`
	Used      uint64 `json:"used"`
	Available uint64 `json:"available"`
}

func (m current_ram) String() string {
	s, _ := json.Marshal(m)
	return string(s)
}
func NewCurrentRamCollector() (Collector, error) {
	var c *current_ram = new(current_ram)
	return c, nil
}

func Newcurrent_ram(Total uint64, Used uint64, Available uint64) *current_ram {

	return &current_ram{Total, Used, Available}
}

func (c *current_ram) Update() (res interface{}, err error) {
	log.Println("Update-----------------\n")
	// load, err := getLoad1()
	if err != nil {
		// return fmt.Errorf("Couldn't get load: %s", err error)
		log.Println(err)

	}
	v, _ := mem.VirtualMemory()
	var t uint64 = 1024 * 1024
	current_ram := &current_ram{v.Total / t, v.Used / t, v.Available / t}
	log.Printf("Total: %v, Free:%v, UsedPercent:%f%%\n", v.Total, v.Available, v.UsedPercent)

	// load1.Collect(ch)
	return current_ram, nil
}
