package tools

import (
	"encoding/json"
	"fmt"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/mem"
	"nctr/pkg/alog"
	"nctr/pkg/constant"
	"nctr/pkg/service"
	"io/ioutil"
	"net/http"
	"runtime"
	"strings"
	"time"
)

type deviceHW struct {
	LogicCores    int
	TotalMemory   uint64
	FreeDisk      uint64
	DownloadSpeed float64
	UploadSpeed   float64
}
type deviceInfo struct {
	OS             string
	TotalMemory    uint64
	FreeMemory     uint64
	TotalDisk      uint64
	UsedDisk       uint64
	FreeDisk       uint64
	LogicCores     int
	NumOfCores     int
	ModelName      string
	CpuSpeed       float64
	HostName       string
	UpTime         uint64
	Platform       string
	HostId         string
	DownloadSpeed  float64
	UploadSpeed    float64
	CheckBee       bool
	Valid          bool
	ReadCount      uint64
	WriteCount     uint64
	Timestamp      int64
	MinimalRequire bool
	Arch           string
	HW             int
}

var (
	DeviceInfo deviceInfo
	beePort    string

	LowHW = deviceHW{
		LogicCores:    8,
		TotalMemory:   uint64(4294967296),
		FreeDisk:      uint64(107374182400),
		DownloadSpeed: float64(50),
		UploadSpeed:   float64(5),
	}

	MediumHW = deviceHW{
		LogicCores:    16,
		TotalMemory:   uint64(4294967296 * 2),
		FreeDisk:      uint64(483183820800),
		DownloadSpeed: float64(50 * 2),
		UploadSpeed:   float64(5 * 2),
	}
	HighHW = deviceHW{
		LogicCores:    32,
		TotalMemory:   uint64(4294967296 * 4),
		FreeDisk:      uint64(1020054732800),
		DownloadSpeed: float64(50 * 6),
		UploadSpeed:   float64(5 * 10),
	}
)

func SetBeePort(port string) {
	beePort = port

}
func DealWithErr(err error) {
	if err != nil {
		fmt.Println(err.Error())

	}
}

//simple check bee node
func checkBeen() bool {

	resp, err := http.Get("http://localhost" + service.Node.SwarmPort + "/addresses")
	if err == nil {
		b, err2 := ioutil.ReadAll(resp.Body)
		if err2 == nil && strings.Contains(string(b), "ethereum") {
			return true
		}
	}
	DealWithErr(err)
	return false
}

func GetHardwareData() deviceInfo {
	runtimeOS := runtime.GOOS
	vmStat, err := mem.VirtualMemory()
	DealWithErr(err)
	DealWithErr(err)
	var totalDisk uint64 = 0
	var totalDiskFree uint64 = 0
	var totalDiskUsed uint64 = 0
	path := service.Node.DataDir
	diskStat, err := disk.Usage(path)
	if diskStat != nil {
		totalDiskUsed = totalDiskUsed + diskStat.Used
		totalDisk = totalDisk + diskStat.Total
		totalDiskFree = totalDiskFree + diskStat.Free
		ioStat, _ := disk.IOCounters(path)
		ioStatCount, ok := ioStat[path]
		if ok {
			DeviceInfo.ReadCount = ioStatCount.ReadBytes * uint64(1000) / uint64(1024) / uint64(1024) / ioStatCount.ReadTime
			DeviceInfo.WriteCount = ioStatCount.WriteBytes * uint64(1000) / uint64(1024) / uint64(1024) / ioStatCount.WriteTime
		}

	}
	DealWithErr(err)
	cpuStat, err := cpu.Info()
	DealWithErr(err)
	hostStat, err := host.Info()
	DealWithErr(err)
	DeviceInfo.OS = runtimeOS
	DeviceInfo.Timestamp = time.Now().Unix()
	DeviceInfo.TotalMemory = vmStat.Total
	DeviceInfo.FreeMemory = vmStat.Free
	DeviceInfo.TotalDisk = totalDisk
	DeviceInfo.UsedDisk = totalDiskUsed
	DeviceInfo.FreeDisk = totalDiskFree
	DeviceInfo.LogicCores = len(cpuStat)
	DeviceInfo.NumOfCores = int(cpuStat[0].Cores)
	DeviceInfo.ModelName = cpuStat[0].ModelName
	DeviceInfo.CpuSpeed = cpuStat[0].Mhz
	DeviceInfo.HostName = hostStat.Hostname
	DeviceInfo.HostId = hostStat.HostID
	DeviceInfo.UpTime = hostStat.Uptime
	DeviceInfo.Platform = hostStat.Platform
	DeviceInfo.Arch = hostStat.KernelArch
	DeviceInfo.CheckBee = checkBeen()
	if DeviceInfo.DownloadSpeed <= 0 || time.Now().Minute() == 0 {
		startTestSpeed()
	}
	hw := getHW(DeviceInfo.LogicCores, DeviceInfo.TotalMemory, DeviceInfo.FreeDisk, DeviceInfo.DownloadSpeed, DeviceInfo.UploadSpeed)
	DeviceInfo.HW = hw
	if hw > 0 {
		DeviceInfo.Valid = true
		DeviceInfo.MinimalRequire = true
	} else {
		DeviceInfo.Valid = false
		DeviceInfo.MinimalRequire = false
		alog.Error(constant.SYSTEM_NOT_SUITABLE)
		alog.Info("Minimal requirements are:")
		lowS, err := json.MarshalIndent(LowHW, "", "\t")
		DealWithErr(err)
		alog.Info(string(lowS))
	}
	dataString, err := json.MarshalIndent(DeviceInfo, "", "\t")
	DealWithErr(err)
	alog.Info("finished check and check data:" + string(dataString))
	return DeviceInfo
}

func startTestSpeed() {
	err := retry(4, 1*time.Second, func(i int) error {
		return TestSpeed(i)
	})
	DealWithErr(err)
}

func TestSpeed(i int) error {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Errorf("request timeout！")
		}
	}()
	down, up := float64(0), float64(0)
	workDone := make(chan struct{}, 1)
	go func() {

		timeout := make(chan bool)
		go func() {
			time.Sleep(60 * time.Second)
			timeout <- true
		}()
		down, up = testSpeed(i)
		workDone <- struct{}{}
	}()
	select {
	case <-workDone:
		if down <= 0 || up <= 0 {
			return fmt.Errorf("timeout speed 0！")
		}
		DeviceInfo.DownloadSpeed = down
		DeviceInfo.UploadSpeed = up
		return nil
	case <-time.After(120 * time.Second):
		return fmt.Errorf("request timeout！")
	}

}
func retry(attempts int, sleep time.Duration, f func(index int) error) (err error) {
	for i := 0; i < attempts; i++ {
		if i > 0 {
			alog.Info("retrying after error:")
			DealWithErr(err)
			time.Sleep(sleep)
		}
		err = f(i)
		if err == nil {
			return nil
		}
	}
	return fmt.Errorf("after %d attempts, last error: %s", attempts, err)
}

func getHW(cpu int, memory uint64, disk uint64, downloadSpeed float64, uploadSpeed float64) int {
	a := suitLevel(cpu, memory, disk, downloadSpeed, uploadSpeed, HighHW)
	if a {
		return 3
	}
	b := suitLevel(cpu, memory, disk, downloadSpeed, uploadSpeed, MediumHW)
	if b {
		return 2
	}
	c := suitLevel(cpu, memory, disk, downloadSpeed, uploadSpeed, LowHW)
	if c {
		return 1
	} else {
		return 0
	}
}

func suitLevel(cpu int, memory uint64, disk uint64, downloadSpeed float64, uploadSpeed float64, hw deviceHW) bool {
	return cpu >= hw.LogicCores && memory >= hw.TotalMemory && (disk+DirSize(service.Node.DataDir)) >= hw.FreeDisk && downloadSpeed >= hw.DownloadSpeed && uploadSpeed >= hw.UploadSpeed
}
func GetHwDiskFree(hw int) uint64 {
	if hw == 1 {
		return LowHW.FreeDisk
	} else if hw == 2 {
		return MediumHW.FreeDisk
	} else {
		return HighHW.FreeDisk
	}

}
