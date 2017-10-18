package collector

import (
	"encoding/json"
	"fmt"
	"log"
	"sort"
	"strconv"

	"github.com/shirou/gopsutil/process"
)

func init() {
	Factories["cpu_intensive_processes"] = NewCollector_cpu_intensive_processes
}

type cpu_intensive_processes struct {
	Pid  int32  `json:"pid"`
	User string `json:"user"`
	Cpu  string `json:"cpu%"`
	Rss  uint64 `json:"rss"`
	Vsz  uint64 `json:"vsz"`
	Cmd  string `json:"cmd"`
}
type cpu_intensive_processesSlice []cpu_intensive_processes

type sortCpu struct {
	Id  int32
	Cpu float64
}
type sortCpuSlice []sortCpu

func (a sortCpuSlice) Len() int      { return len(a) }
func (a sortCpuSlice) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a sortCpuSlice) Less(i, j int) bool {
	return a[i].Cpu > a[j].Cpu
}

func (m cpu_intensive_processes) String() string {
	s, _ := json.Marshal(m)
	return string(s)
}
func NewCollector_cpu_intensive_processes() (Collector, error) {
	var c *cpu_intensive_processes = new(cpu_intensive_processes)
	return c, nil
}

func New_cpu_intensive_processes(Pid int32, User string, Cpu string, Rss uint64, Vsz uint64, Cmd string) *cpu_intensive_processes {

	return &cpu_intensive_processes{Pid, User, Cpu, Rss, Vsz, Cmd}
}

func (c *cpu_intensive_processes) Update() (res interface{}, err error) {
	log.Println("Update---cpu_intensive_processes\n")
	fmt.Println(get_cpu_intensive_processes())

	return get_cpu_intensive_processes(), nil
}

func get_cpu_intensive_processes() *cpu_intensive_processesSlice {
	pids, err := process.Pids()
	if err != nil {
		log.Fatalln("Could not read processes,", err)
	}
	var members []sortCpu
	var s_cpu_intensive_processesSlice cpu_intensive_processesSlice
	for _, pid := range pids {
		proc, err := process.NewProcess(pid)
		if err != nil {
			continue
			log.Fatalln("Could not create1 ,", err)
		}

		cpu, err := proc.CPUPercent()
		if err != nil {
			continue
			log.Fatalln("Could not create2 ,", err)
		}
		members = append(members, sortCpu{pid, cpu})
	}
	sort.Sort(sortCpuSlice(members))

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
		cpu := members[i].Cpu
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
		s_cpu_intensive_processesSlice = append(s_cpu_intensive_processesSlice, cpu_intensive_processes{members[i].Id, user, strconv.FormatFloat(cpu, 'f', 1, 64), rss, vms, cmd})
	}
	return &s_cpu_intensive_processesSlice
}
