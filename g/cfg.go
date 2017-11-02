package g

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/kikiyou/agent/collector"
)

func LoadCollectors(appConfig Config) (map[string]collector.Collector, error) {
	collectors := map[string]collector.Collector{}

	for _, name := range strings.Split(appConfig.EnabledCollectors, ",") {
		fn, ok := collector.Factories[name]
		if !ok {
			log.Fatalf("Collector '%s' not available", name)
		}
		c, err := fn()
		if err != nil {
			return nil, err
		}
		collectors[name] = c
	}
	return collectors, nil
}

// Config - config struct
type Config struct {
	Port              int // server port
	ListeningAddress  string
	Cache             int    // caching command out (in seconds)
	Timeout           int    // timeout for shell command (in seconds)
	Host              string // server host
	Shell             string // custom shell
	DefaultShell      string // shell by default
	DefaultShOpt      string // shell option for one-liner (-c or /C)
	AuthUser          string // basic authentication user name
	AuthPass          string // basic authentication password
	AddExit           bool   // add /exit command
	PublicDir         string // 共享目录
	EnabledCollectors string
}

func getConfig() (appConfig Config, err error) {

	var (
		// configFile        = flag.String("config", "xxx.conf", "config file.")
		basicAuth        = flag.String("basic-auth", "admin:admin", "setup HTTP Basic Authentication (\"user_name:password\")")
		showVersion      = flag.Bool("version", false, "Print version information.")
		listeningAddress = flag.String("listen", ":8899", "address to listen on")
		// enabledCollectors = flag.String("enabledCollectors", "user_accounts,disk_partitions", "comma-seperated list of collectors to use")
		enabledCollectors = flag.String("enabledCollectors", "current_ram,load_avg", "comma-seperated list of collectors to use")
		publicSharePath   = flag.String("public", filepath.Join(Root, "public"), "public share dir")
		printCollectors   = flag.Bool("printCollectors", false, "If true, print available collectors and exit")
	)
	flag.Parse()

	if *showVersion {
		fmt.Fprintln(os.Stdout, VERSION)
		os.Exit(0)
	}
	log.Println("Starting version", VERSION)
	if *printCollectors {
		log.Printf("Available collectors:\n")
		for n, _ := range collector.Factories {
			log.Printf(" - %s\n", n)
		}
		return
	}

	if *basicAuth == "" && len(os.Getenv(BasicAuthVar)) > 0 {
		*basicAuth = os.Getenv(BasicAuthVar)
	}
	if len(*basicAuth) > 0 {
		basicAuthParts := strings.SplitN(*basicAuth, ":", 2)
		if len(basicAuthParts) != 2 {
			return Config{}, fmt.Errorf("HTTP basic authentication must be in format: name:password, got: %s", *basicAuth)
		}
		appConfig.AuthUser, appConfig.AuthPass = basicAuthParts[0], basicAuthParts[1]
	}
	appConfig.ListeningAddress = *listeningAddress
	appConfig.PublicDir = *publicSharePath
	appConfig.EnabledCollectors = *enabledCollectors
	// var c Config
	// lock.Lock()
	// defer lock.Unlock()

	// AppConfig = &appConfig
	return appConfig, nil
}
