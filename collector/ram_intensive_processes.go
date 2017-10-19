package collector

import (
	"encoding/json"
	"log"
	"sort"
	"strconv"

	"github.com/shirou/gopsutil/process"
)

func init() {
	Factories["ram_intensive_processes"] = NewCollector_ram_intensive_processes
}

type ram_intensive_processes struct {
	Pid  int32  `json:"pid"`
	User string `json:"user"`
	Mem  string `json:"mem%"`
	Rss  uint64 `json:"rss"`
	Vsz  uint64 `json:"vsz"`
	Cmd  string `json:"cmd"`
}
type ram_intensive_processesSlice []ram_intensive_processes

type sortMem struct {
	Id  int32
	Mem float32
}
type sortMemSlice []sortMem

func (a sortMemSlice) Len() int      { return len(a) }
func (a sortMemSlice) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a sortMemSlice) Less(i, j int) bool {
	return a[i].Mem > a[j].Mem
}

func (m ram_intensive_processes) String() string {
	s, _ := json.Marshal(m)
	return string(s)
}
func NewCollector_ram_intensive_processes() (Collector, error) {
	var c *ram_intensive_processes = new(ram_intensive_processes)
	return c, nil
}

func New_ram_intensive_processes(Pid int32, User string, Mem string, Rss uint64, Vsz uint64, Cmd string) *ram_intensive_processes {

	return &ram_intensive_processes{Pid, User, Mem, Rss, Vsz, Cmd}
}

func (c *ram_intensive_processes) Update() (res interface{}, err error) {
	log.Println("Update---ram_intensive_processes\n")
	// fmt.Println(get_ram_intensive_processes())

	return get_ram_intensive_processes(), nil
}

func get_ram_intensive_processes() *ram_intensive_processesSlice {
	pids, err := process.Pids()
	if err != nil {
		log.Fatalln("Could not read processes,", err)
	}
	var members []sortMem
	var s_ram_intensive_processesSlice ram_intensive_processesSlice
	for _, pid := range pids {
		proc, err := process.NewProcess(pid)
		if err != nil {
			continue
			log.Fatalln("Could not create1 ,", err)
		}

		mem, err := proc.MemoryPercent()
		if err != nil {
			continue
			log.Fatalln("Could not create2 ,", err)
		}
		members = append(members, sortMem{pid, mem})
	}
	sort.Sort(sortMemSlice(members))

	for i := range members[:15] {
		// println(members[i].Id)
		proc, err := process.NewProcess(members[i].Id)
		if err != nil {
			continue
			log.Fatalln("Could not create3 ,", err)
		}
		user, err := proc.Username()
		if err != nil {
			continue
			log.Fatalln("Could not read process name,", err)
		}
		mem := members[i].Mem
		meminfo, err := proc.MemoryInfo()
		if err != nil {
			continue
			log.Fatalln("Could not read process name,", err)
		}
		cmd, err := proc.Name()
		if err != nil {
			continue
			log.Fatalln("Could not read process name,", err)
		}
		var t uint64 = 1024 * 1024
		rss := meminfo.RSS / t
		vms := meminfo.VMS / t

		// log.Println("pid:", members[i].Id, "user:", user, "mem:", mem, "rss:", rss, "vms:", vms, "cmd:", cmd)
		s_ram_intensive_processesSlice = append(s_ram_intensive_processesSlice, ram_intensive_processes{members[i].Id, user, strconv.FormatFloat(float64(mem), 'f', 1, 32), rss, vms, cmd})
	}
	return &s_ram_intensive_processesSlice
}
