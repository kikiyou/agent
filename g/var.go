package g

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/kikiyou/agent/collector"
)

var (
	Root            string
	Collectors      map[string]collector.Collector
	TempScriptsFile string
	// PublicPath      string
)

func CheckErr(e error) {
	if e != nil {
		panic(e)
	}
}

func InitRootDir() {
	var err error
	Root, err = os.Getwd()
	if err != nil {
		log.Fatalln("getwd fail:", err)
	}
}

func InitTempScriptsFile() {
	var err error
	tmpfile, err := ioutil.TempFile("/tmp", "linux_json_api")
	if err != nil {
		log.Fatalln("TempScriptsFile fail:", err)
	}
	TempScriptsFile = tmpfile.Name()

	defer os.Remove(TempScriptsFile) // clean up

}
