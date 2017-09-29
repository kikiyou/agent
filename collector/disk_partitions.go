package collector

import (
	"encoding/json"
	"log"
	"strconv"

	"github.com/shirou/gopsutil/disk"
)

func init() {
	Factories["disk_partitions"] = NewCollector_disk_partitions
}

type disk_partition struct {
	File_system string  `json:"file_system"`
	Size        string  `json:"size"`
	Used        string  `json:"used"`
	Avail       string  `json:"avail"`
	UsedPercent float64 `json:"used%"`
	Mounted     string  `json:"mounted"`
}

type disk_partitions []disk_partition

func (m disk_partition) String() string {
	s, _ := json.Marshal(m)
	return string(s)
}
func NewCollector_disk_partitions() (Collector, error) {
	var c *disk_partitions = new(disk_partitions)
	return c, nil
}

func (c *disk_partitions) Update() (res interface{}, err error) {
	log.Println("Update-----------------\n")

	// log.Println(get_disk_partition())
	return get_disk_partition(), nil
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func get_disk_partition() *disk_partitions {
	parts, err := disk.Partitions(false)
	check(err)

	// var usage []*disk.UsageStat
	var usage disk_partitions

	for _, part := range parts {
		u, err := disk.Usage(part.Mountpoint)
		check(err)
		d := disk_partition{
			File_system: u.Fstype,
			Size:        strconv.FormatUint(u.Total/1024/1024/1024, 10) + " GiB",
			Used:        strconv.FormatUint(u.Used/1024/1024/1024, 10) + " GiB",
			Avail:       strconv.FormatUint(u.Free/1024/1024/1024, 10) + " GiB",
			UsedPercent: u.UsedPercent + 5,
			Mounted:     u.Path,
		}
		usage = append(usage, d)
	}
	return &usage
}
