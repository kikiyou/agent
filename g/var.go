package g

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/kikiyou/agent/collector"
	// "github.com/kikiyou/agent/collector"
	// "github.com/kikiyou/agent/collector"
)

var (
	Root            string
	Collectors      map[string]collector.Collector
	TempScriptsFile string
	PublicPath      string
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
	// if err := tmpfile.Close(); err != nil {
	// 	log.Fatal(err)
	// }

}

// var LocalIp string

// func InitLocalIp() {
// 	if Config().Heartbeat.Enabled {
// 		conn, err := net.DialTimeout("tcp", Config().Heartbeat.Addr, time.Second*10)
// 		// log.Println(Config().Heartbeat.Addr)
// 		if err != nil {
// 			log.Println("get local addr failed !")
// 		} else {
// 			LocalIp = strings.Split(conn.LocalAddr().String(), ":")[0]
// 			conn.Close()
// 		}
// 	} else {
// 		log.Println("hearbeat is not enabled, can't get localip")
// 	}
// }

// var (
// 	HbsClient *SingleConnRpcClient
// )

// func InitRpcClients() {
// 	if Config().Heartbeat.Enabled {
// 		HbsClient = &SingleConnRpcClient{
// 			RpcServer: Config().Heartbeat.Addr,
// 			Timeout:   time.Duration(Config().Heartbeat.Timeout) * time.Millisecond,
// 		}
// 	}
// }

// func SendToTransfer(metrics []*model.MetricValue) {
// 	if len(metrics) == 0 {
// 		return
// 	}

// 	debug := Config().Debug

// 	if debug {
// 		log.Printf("=> <Total=%d> %v\n", len(metrics), metrics[0])
// 	}

// 	var resp model.TransferResponse
// 	SendMetrics(metrics, &resp)

// 	if debug {
// 		log.Println("<=", &resp)
// 	}
// }

// var (
// 	reportUrls     map[string]string
// 	reportUrlsLock = new(sync.RWMutex)
// )

// func ReportUrls() map[string]string {
// 	reportUrlsLock.RLock()
// 	defer reportUrlsLock.RUnlock()
// 	return reportUrls
// }

// func SetReportUrls(urls map[string]string) {
// 	reportUrlsLock.RLock()
// 	defer reportUrlsLock.RUnlock()
// 	reportUrls = urls
// }

// var (
// 	reportPorts     []int64
// 	reportPortsLock = new(sync.RWMutex)
// )

// func ReportPorts() []int64 {
// 	reportPortsLock.RLock()
// 	defer reportPortsLock.RUnlock()
// 	return reportPorts
// }

// func SetReportPorts(ports []int64) {
// 	reportPortsLock.Lock()
// 	defer reportPortsLock.Unlock()
// 	reportPorts = ports
// }

// var (
// 	duPaths     []string
// 	duPathsLock = new(sync.RWMutex)
// )

// func DuPaths() []string {
// 	duPathsLock.RLock()
// 	defer duPathsLock.RUnlock()
// 	return duPaths
// }

// func SetDuPaths(paths []string) {
// 	duPathsLock.Lock()
// 	defer duPathsLock.Unlock()
// 	duPaths = paths
// }

// var (
// 	// tags => {1=>name, 2=>cmdline}
// 	// e.g. 'name=falcon-agent'=>{1=>falcon-agent}
// 	// e.g. 'cmdline=xx'=>{2=>xx}
// 	reportProcs     map[string]map[int]string
// 	reportProcsLock = new(sync.RWMutex)
// )

// func ReportProcs() map[string]map[int]string {
// 	reportProcsLock.RLock()
// 	defer reportProcsLock.RUnlock()
// 	return reportProcs
// }

// func SetReportProcs(procs map[string]map[int]string) {
// 	reportProcsLock.Lock()
// 	defer reportProcsLock.Unlock()
// 	reportProcs = procs
// }

// var (
// 	ips     []string
// 	ipsLock = new(sync.Mutex)
// )

// func TrustableIps() []string {
// 	ipsLock.Lock()
// 	defer ipsLock.Unlock()
// 	return ips
// }

// func SetTrustableIps(ipStr string) {
// 	arr := strings.Split(ipStr, ",")
// 	ipsLock.Lock()
// 	defer ipsLock.Unlock()
// 	ips = arr
// }

// func IsTrustable(remoteAddr string) bool {
// 	ip := remoteAddr
// 	idx := strings.LastIndex(remoteAddr, ":")
// 	if idx > 0 {
// 		ip = remoteAddr[0:idx]
// 	}

// 	if ip == "127.0.0.1" {
// 		return true
// 	}

// 	return slice.ContainsString(TrustableIps(), ip)
// }

//////////////////////////
// // Implements Collector.
// type NodeCollector struct {
// 	collectors map[string]collector.Collector
// }

// // Implements Collector.
// func (n NodeCollector) Describe(ch chan<- *prometheus.Desc) {
// 	scrapeDurations.Describe(ch)
// }

// Implements Collector.
// func (n NodeCollector) Collect(ch chan<- prometheus.Metric) {
// 	wg := sync.WaitGroup{}
// 	wg.Add(len(n.collectors))
// 	for name, c := range n.collectors {
// 		go func(name string, c collector.Collector) {
// 			Execute(name, c, ch)
// 			wg.Done()
// 		}(name, c)
// 	}
// 	wg.Wait()
// 	// scrapeDurations.Collect(ch)
// }

// func Execute(name string, c collector.Collector, ch chan<- prometheus.Metric) {
// 	begin := time.Now()
// 	err := c.Update(ch)
// 	duration := time.Since(begin)
// 	var result string

// 	if err != nil {
// 		glog.Infof("ERROR: %s failed after %fs: %s", name, duration.Seconds(), err)
// 		result = "error"
// 	} else {
// 		glog.Infof("OK: %s success after %fs.", name, duration.Seconds())
// 		result = "success"
// 	}
// 	scrapeDurations.WithLabelValues(name, result).Observe(duration.Seconds())
// }

// var (
// 	NodeCollector NodeCollector
// )

// func RegisterNodeCollector(n NodeCollector) {
// 	NodeCollector = n
// }
