package g

import (
	"log"
	"runtime"
)

var AppConfig Config

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	InitRootDir()
	AppConfig, _ = getConfig()
}
