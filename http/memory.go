package http

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/shirou/gopsutil/mem"
	"github.com/toolkits/nux"
)

type current_ram struct {
	Total     uint64 `json:"total"`
	Used      uint64 `json:"used"`
	Available uint64 `json:"available"`
}

func (m current_ram) String() string {
	s, _ := json.Marshal(m)
	return string(s)
}

func Newcurrent_ram(Total uint64, Used uint64, Available uint64) *current_ram {
	// if fd < 0 {
	// 	return nil
	// }

	return &current_ram{Total, Used, Available}
}

func configMemoryRoutes() {
	http.HandleFunc("/server/", func(w http.ResponseWriter, r *http.Request) {
		module := r.URL.Query().Get("module")
		if module == "" {
			http.Error(w, "No module specified, or requested module doesn't exist.", 406)
			return
		}

		v, _ := mem.VirtualMemory()
		var t uint64 = 1024 * 1024
		current_ram := &current_ram{v.Total / t, v.Used / t, v.Free / t}
		fmt.Printf("Total: %v, Free:%v, UsedPercent:%f%%\n", v.Total, v.Free, v.UsedPercent)

		RenderJson(w, current_ram)
	})

	http.HandleFunc("/page/memory", func(w http.ResponseWriter, r *http.Request) {
		mem, err := nux.MemInfo()
		if err != nil {
			RenderMsgJson(w, err.Error())
			return
		}

		memFree := mem.MemFree + mem.Buffers + mem.Cached
		memUsed := mem.MemTotal - memFree
		var t uint64 = 1024 * 1024
		RenderDataJson(w, []interface{}{mem.MemTotal / t, memUsed / t, memFree / t})
	})

	http.HandleFunc("/proc/memory", func(w http.ResponseWriter, r *http.Request) {
		mem, err := nux.MemInfo()
		if err != nil {
			RenderMsgJson(w, err.Error())
			return
		}

		memFree := mem.MemFree + mem.Buffers + mem.Cached
		memUsed := mem.MemTotal - memFree

		RenderDataJson(w, map[string]interface{}{
			"total": mem.MemTotal,
			"free":  memFree,
			"used":  memUsed,
		})
	})
}
