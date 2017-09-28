package http

import (
	"fmt"
	"net/http"

	"github.com/toolkits/nux"
)

func configMemoryRoutes() {
	http.HandleFunc("/server/", func(w http.ResponseWriter, r *http.Request) {
		module := r.URL.Query().Get("module")
		if module == "" {
			http.Error(w, "No module specified, or requested module doesn't exist.", 406)
			return
		}

		mem, err := nux.MemInfo()
		if err != nil {
			RenderMsgJson(w, err.Error())
			return
		}

		memFree := mem.MemFree + mem.Buffers + mem.Cached
		memUsed := mem.MemTotal - memFree
		var t uint64 = 1024 * 1024
		fmt.Println("###################")
		RenderDataJson(w, []interface{}{mem.MemTotal / t, memUsed / t, memFree / t})
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
