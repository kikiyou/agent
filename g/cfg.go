package g

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"github.com/kikiyou/agent/collector"
	cacheGo "github.com/patrickmn/go-cache"
)

// var (
// 	CacheTTL *cacheGo.Cache
// )

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
	Collectors        map[string]collector.Collector
	Secret            string
	CacheTTL          *cacheGo.Cache
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
		cache             = flag.Int("cache", 2, "设定cache默认是:2")
		secret            = flag.String("secret", "secret", "随机token加密因子")
	)
	flag.Parse()

	switch runtime.GOOS {
	case "plan9":
		appConfig.DefaultShell, appConfig.DefaultShOpt = defaultShellPlan9, "-c"
	case "windows":
		appConfig.DefaultShell, appConfig.DefaultShOpt = defaultShellWindows, "/C"
	default:
		appConfig.DefaultShell, appConfig.DefaultShOpt = defaultShellPOSIX, "-c"
	}

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
	appConfig.Shell = appConfig.DefaultShell
	appConfig.Cache = *cache
	appConfig.Secret = *secret

	var CacheTTL *cacheGo.Cache
	if appConfig.Cache > 0 {
		CacheTTL = cacheGo.New(5*time.Minute, 10*time.Minute)
	}
	appConfig.CacheTTL = CacheTTL
	// appConfig.DefaultShOpt = "-c"
	// var c Config
	// lock.Lock()
	// defer lock.Unlock()

	// AppConfig = &appConfig
	return appConfig, nil
}
