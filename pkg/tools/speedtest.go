package tools

import (
	"encoding/xml"
	"errors"
	"fmt"
	"github.com/kellydunn/golang-geo"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"runtime"
	"sort"
	"strings"
	"sync"
	"time"
)

type CliFlags struct {
	List        bool
	Server      int
	Interactive bool // Not a direct flag, this is derived from whether a user has or has not selected a machine readable output
	Json        bool
	Xml         bool
	Csv         bool
	Simple      bool
	Source      string
	Timeout     int64
	Share       bool
}

type Latency struct {
	Length float64 `xml:"testlength,attr"`
}

type Configuration struct {
	Client       Client       `xml:"client"`
	ServerConfig ServerConfig `xml:"server-config"`
	Times        Times        `xml:"times"`
	Download     Download     `xml:"socket-download"`
	Upload       Upload       `xml:"socket-upload"`
	Latency      Latency      `xml:"socket-latency"`
}

type Server struct {
	CC        string        `xml:"cc,attr" json:"cc"`
	Country   string        `xml:"country,attr" json:"country"`
	ID        int           `xml:"id,attr" json:"id"`
	Latitude  float64       `xml:"lat,attr" json:"lat"`
	Longitude float64       `xml:"lon,attr" json:"lon"`
	Name      string        `xml:"name,attr" json:"name"`
	Sponsor   string        `xml:"sponsor,attr" json:"sponsor"`
	URL       string        `xml:"url,attr" json:"url"`
	URL2      string        `xml:"url2,attr" json:"url2"`
	Host      string        `xml:"host,attr" json:"host"`
	Distance  float64       `xml:"distance,attr" json:"distance"`
	Latency   time.Duration `xml:"latency,attr" json:"latency"`
	speedtest *Speedtest
	tcpAddr   *net.TCPAddr
}

type Servers struct {
	Servers []Server `xml:"servers>server"`
}

type Results struct {
	XMLName   xml.Name  `json:"-" xml:"results"`
	Download  float64   `json:"download" xml:"download"`
	Upload    float64   `json:"upload" xml:"upload"`
	Latency   float64   `json:"latency" xml:"latency"`
	Server    *Server   `json:"server" xml:"server"`
	Timestamp time.Time `json:"timestamp" xml:"timestamp"`
	Share     string    `json:"share" xml:"share"`
}

type Speedtest struct {
	Configuration *Configuration
	Servers       *Servers
	CliFlags      *CliFlags
	Results       *Results
	Source        *net.TCPAddr
	Timeout       time.Duration
}

type Client struct {
	IP        string  `xml:"ip,attr"`
	ISP       string  `xml:"isp,attr"`
	Latitude  float64 `xml:"lat,attr"`
	Longitude float64 `xml:"lon,attr"`
}

type ServerConfig struct {
	IgnoreIDs   string `xml:"ignoreids,attr"`
	ThreadCount string `xml:"threadcount,attr"`
}

type Times struct {
	DownloadOne   int `xml:"dl1,attr"`
	DownloadTwo   int `xml:"dl2,attr"`
	DownloadThree int `xml:"dl3,attr"`
	UploadOne     int `xml:"ul1,attr"`
	UploadTwo     int `xml:"ul2,attr"`
	UploadThree   int `xml:"ul3,attr"`
}

type Download struct {
	Length       float64 `xml:"testlength,attr"`
	PacketLength int     `xml:"packetlength,attr"`
}

type Upload struct {
	Length       float64 `xml:"testlength,attr"`
	PacketLength int     `xml:"packetlength,attr"`
}

var SIndex int

// Helper function to make it easier for printing and exiting
func errorf(text string, a ...interface{}) error {
	if !strings.HasSuffix(text, "\n") {
		text += "\n"
	}
	return fmt.Errorf(text, a...)
}

// Established connection with local address and timeout support
func dialTimeout(network string, laddr *net.TCPAddr, raddr *net.TCPAddr, timeout time.Duration) (net.Conn, error) {
	dialer := &net.Dialer{
		Timeout:   timeout,
		LocalAddr: laddr,
	}

	conn, err := dialer.Dial(network, raddr.String())
	return conn, err
}

func NewCliFlags() *CliFlags {
	return &CliFlags{
		Interactive: true,
	}
}

func NewResults() *Results {
	return &Results{
		Timestamp: time.Now(),
	}
}

func NewSpeedtest() *Speedtest {
	return &Speedtest{
		Configuration: &Configuration{},
		Servers:       &Servers{},
		CliFlags:      NewCliFlags(),
		Results:       NewResults(),
	}
}

func (s *Speedtest) Printf(text string, a ...interface{}) {
	if !s.CliFlags.Interactive {
		return
	}
	fmt.Printf(text, a...)
}

func (s *Speedtest) GetConfiguration() (*Configuration, error) {
	res, err := http.Get("https://www.speedtest.net/speedtest-config.php")
	if err != nil {
		return s.Configuration, errors.New("Error retrieving Speedtest.net configuration: " + err.Error())
	}
	defer res.Body.Close()
	settingsBody, _ := ioutil.ReadAll(res.Body)
	xml.Unmarshal(settingsBody, &s.Configuration)
	return s.Configuration, nil
}

func (s *Speedtest) GetServers(serverId int) (*Servers, error) {
	res, err := http.Get("https://www.speedtest.net/speedtest-servers.php")
	if err != nil {
		return s.Servers, errors.New("Error retrieving Speedtest.net servers: " + err.Error())
	}
	defer res.Body.Close()
	serversBody, _ := ioutil.ReadAll(res.Body)
	var allServers Servers
	xml.Unmarshal(serversBody, &allServers)
	for _, server := range allServers.Servers {
		server.speedtest = s
		if serverId == 0 || server.ID == serverId {
			s.Servers.Servers = append(s.Servers.Servers, server)
		}
	}

	return s.Servers, nil
}

func (s *Servers) SortServersByDistance() {
	ps := &serverSorter{
		servers: s.Servers,
		by: func(s1, s2 *Server) bool {
			return s1.Distance < s2.Distance
		},
	}
	sort.Sort(ps)
}

func (s *Servers) SortServersByLatency() {
	ps := &serverSorter{
		servers: s.Servers,
		by: func(s1, s2 *Server) bool {
			// Latency should never be 0 unless we didn't test latency for that server
			if s1.Latency == 0 {
				return false
			}
			return s1.Latency < s2.Latency
		},
	}
	sort.Sort(ps)
}

// serverSorter joins a By function and a slice of Servers to be sorted.
type serverSorter struct {
	servers []Server
	by      func(s1, s2 *Server) bool // Closure used in the Less method.
}

// Len is part of sort.Interface.
func (s *serverSorter) Len() int {
	return len(s.servers)
}

// Swap is part of sort.Interface.
func (s *serverSorter) Swap(i, j int) {
	s.servers[i], s.servers[j] = s.servers[j], s.servers[i]
}

// Less is part of sort.Interface. It is implemented by calling the "by" closure in the sorter.
func (s *serverSorter) Less(i, j int) bool {
	return s.by(&s.servers[i], &s.servers[j])
}

func (s *Servers) SetDistances(latitude, longitude float64) {
	me := geo.NewPoint(latitude, longitude)
	for i, server := range s.Servers {
		serverPoint := geo.NewPoint(server.Latitude, server.Longitude)
		distance := me.GreatCircleDistance(serverPoint)
		s.Servers[i].Distance = distance
	}
}

func (s *Servers) TestLatency() (*Server, error) {
	var servers []Server
	s.SortServersByDistance()

	if len(s.Servers) >= 5 {
		servers = s.Servers[:5]
	} else {
		servers = s.Servers[:len(s.Servers)]
	}

	for i, server := range servers {
		addr, err := net.ResolveTCPAddr("tcp", server.Host)
		s.Servers[i].tcpAddr = addr
		if err != nil {
			server.speedtest.Printf("%s\n", err.Error())
			continue
		}

		conn, err := dialTimeout("tcp", server.speedtest.Source, addr, server.speedtest.Timeout)
		if err != nil {
			server.speedtest.Printf("%s\n", err.Error())
			continue
		}
		if conn == nil {
			return nil, errorf("conn faild")
		}
		defer func() {
			if conn != nil {
				conn.Close()
			}
		}()

		conn.Write([]byte("HI\n"))
		hello := make([]byte, 1024)
		conn.Read(hello)

		sum := time.Duration(0)
		for j := 0; j < 3; j++ {
			resp := make([]byte, 1024)
			start := time.Now()
			conn.Write([]byte(fmt.Sprintf("PING %d\n", start.UnixNano()/1000000)))
			conn.Read(resp)
			total := time.Since(start)
			sum += total
		}
		s.Servers[i].Latency = sum / 3
	}
	s.SortServersByLatency()
	if SIndex > (len(s.Servers) - 1) {
		SIndex = len(s.Servers) - 1
	}
	if len(s.Servers) < 1 {
		return nil, errorf("can't not find server,please check net")
	}
	return &s.Servers[SIndex], nil
}

func (s *Server) Downloader(ci chan int, co chan []int, wg *sync.WaitGroup, start time.Time, length float64) {
	defer wg.Done()

	conn, err := dialTimeout("tcp", s.speedtest.Source, s.tcpAddr, s.speedtest.Timeout)
	if err != nil {
		errorf("\nCannot connect to %s\n", s.tcpAddr.String())
	}
	if conn == nil {
		return
	}
	conn.Write([]byte("HI\n"))
	hello := make([]byte, 1024)
	conn.Read(hello)
	var ask int
	tmp := make([]byte, 1024)

	var out []int

	for size := range ci {
		s.speedtest.Printf(".")
		remaining := size

		for remaining > 0 && time.Since(start).Seconds() < length {

			if remaining > 1000000 {
				ask = 1000000
			} else {
				ask = remaining
			}
			down := 0

			conn.Write([]byte(fmt.Sprintf("DOWNLOAD %d\n", ask)))

			for down < ask {
				n, err := conn.Read(tmp)
				if err != nil {
					if err != io.EOF {
						fmt.Printf("ERR: %v\n", err)
					}
					break
				}
				down += n
			}
			out = append(out, down)
			remaining -= down

		}
		s.speedtest.Printf(".")
	}
	defer conn.Close()

	defer func() {
		err := recover()
		if err != nil {
			errorf("conn faild")
		}
	}()
	go func(co chan []int, out []int) {
		co <- out
	}(co, out)

}

func (s *Server) TestDownload(length float64) (float64, time.Duration) {
	ci := make(chan int)
	co := make(chan []int)
	wg := new(sync.WaitGroup)
	sizes := []int{245388, 505544, 1118012, 1986284, 4468241, 7907740, 12407926, 17816816, 24262167, 31625365}
	start := time.Now()

	for i := 0; i < 8; i++ {
		wg.Add(1)
		go s.Downloader(ci, co, wg, start, length)
	}

	for _, size := range sizes {
		for i := 0; i < 4; i++ {
			ci <- size
		}
	}

	close(ci)
	wg.Wait()

	total := time.Since(start)
	s.speedtest.Printf("\n")

	var totalSize int
	for i := 0; i < 8; i++ {
		chunks := <-co
		for _, chunk := range chunks {
			totalSize += chunk
		}
	}
	defer func() {
		err := recover()
		if err != nil {
			errorf("conn faild")
		}
	}()
	return float64(totalSize) * 8, total
}

func (s *Server) Uploader(ci chan int, co chan []int, wg *sync.WaitGroup, start time.Time, length float64) {
	defer wg.Done()
	conn, err := dialTimeout("tcp", s.speedtest.Source, s.tcpAddr, s.speedtest.Timeout)
	if err != nil {
		errorf("\nCannot connect to %s\n", s.tcpAddr.String())
	}
	if conn == nil {
		return
	}
	conn.Write([]byte("HI\n"))
	hello := make([]byte, 1024)
	conn.Read(hello)
	defer conn.Close()
	var give int
	var out []int
	for size := range ci {
		s.speedtest.Printf(".")
		remaining := size

		for remaining > 0 && time.Since(start).Seconds() < length {
			if remaining > 100000 {
				give = 100000
			} else {
				give = remaining
			}
			header := []byte(fmt.Sprintf("UPLOAD %d 0\n", give))
			data := make([]byte, give-len(header))

			conn.Write(header)
			conn.Write(data)
			up := make([]byte, 24)
			conn.Read(up)

			out = append(out, give)
			remaining -= give
		}
		s.speedtest.Printf(".")
	}
	defer func() {
		err := recover()
		if err != nil {
			errorf("conn faild")
		}
	}()
	go func(co chan []int, out []int) {
		co <- out
	}(co, out)

}

func (s *Server) TestUpload(length float64) (float64, time.Duration) {
	ci := make(chan int)
	co := make(chan []int)
	wg := new(sync.WaitGroup)
	sizes := []int{32768, 65536, 131072, 262144, 524288, 1048576, 7340032}
	start := time.Now()

	for i := 0; i < 8; i++ {
		wg.Add(1)
		go s.Uploader(ci, co, wg, start, length)
	}

	var tmp int
	for _, size := range sizes {
		for i := 0; i < 4; i++ {
			tmp += size
			ci <- size
		}
	}
	close(ci)
	wg.Wait()

	total := time.Since(start)
	s.speedtest.Printf("\n")

	var totalSize int
	for i := 0; i < 8; i++ {
		chunks := <-co
		for _, chunk := range chunks {
			totalSize += chunk
		}
	}

	return float64(totalSize) * 8, total
}

func testSpeed(index int) (float64, float64) {
	SIndex = index
	speedtest := NewSpeedtest()

	speedtest.Timeout = time.Duration(speedtest.CliFlags.Timeout) * time.Second

	if speedtest.CliFlags.Source != "" {
		source, err := net.ResolveTCPAddr("tcp", speedtest.CliFlags.Source+":0")
		if err != nil {
			errorf("Could not parse source IP address %s: %s", speedtest.CliFlags.Source, err.Error())
		} else {
			speedtest.Source = source
		}
	} else {
		speedtest.Source = nil
	}

	if speedtest.CliFlags.Json || speedtest.CliFlags.Xml || speedtest.CliFlags.Csv || speedtest.CliFlags.Simple {
		speedtest.CliFlags.Interactive = false
	}

	// ALL THE CPUS!
	runtime.GOMAXPROCS(runtime.NumCPU())

	config, err := speedtest.GetConfiguration()
	if err != nil {
		errorf(err.Error())
	}

	servers, err := speedtest.GetServers(speedtest.CliFlags.Server)
	if err != nil {
		errorf(err.Error())
	} else if len(servers.Servers) == 0 {
		err = fmt.Errorf("Failed to retrieve servers or invalid server ID specified")
	}

	servers.SetDistances(config.Client.Latitude, config.Client.Longitude)

	if speedtest.CliFlags.List {
		servers.SortServersByDistance()
		for _, server := range servers.Servers {
			speedtest.Printf("%5d) %s (%s, %s) [%0.2f km]\n", server.ID, server.Sponsor, server.Name, server.Country, server.Distance)
		}
		os.Exit(0)
	}

	speedtest.Results.Server, err = servers.TestLatency()
	if err != nil {
		fmt.Println(err)
		return 0, 0
	}
	speedtest.Results.Latency = float64(speedtest.Results.Server.Latency.Nanoseconds()) / 1000000.0
	if speedtest.Results.Server.Latency == 0 {
		errorf("Unable to test server latency, this may be caused by a connection failure")
	}

	speedtest.Printf("Testing Download Speed")
	downBits, downDuration := speedtest.Results.Server.TestDownload(config.Download.Length)
	speedtest.Results.Download = downBits / downDuration.Seconds()
	speedtest.Printf("Download: %0.2f Mbit/s\n", speedtest.Results.Download/1000/1000)

	speedtest.Printf("Testing Upload Speed")
	upBits, upDuration := speedtest.Results.Server.TestUpload(config.Upload.Length)
	speedtest.Results.Upload = upBits / upDuration.Seconds()
	speedtest.Printf("Upload: %0.2f Mbit/s\n", speedtest.Results.Upload/1000/1000)
	down, up := speedtest.Results.Download/1000/1000, speedtest.Results.Upload/1000/1000
	defer func() {
		err := recover()
		if err != nil {
			down = 0
			up = 0
		}
	}()
	return down, up
}
