package monitorModels

import (
	"fmt"
	"github.com/bzdanny/BaiZe/baize/utils/ipUtils"
	"github.com/bzdanny/BaiZe/baize/utils/timeUtils"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/util/gconv"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/load"
	"github.com/shirou/gopsutil/v3/mem"
	"os"
	"runtime"
	"strconv"
	"time"
)

type server struct {
	CpuNum          int              `json:"cpuNum"`
	CpuNumThread    int              `json:"cpuNumThread"`
	CpuUsed         float64          `json:"cpuUsed"`
	CpuAvg5         float64          `json:"cpuAvg5"`
	CpuAvg15        float64          `json:"cpuAvg15"`
	MemTotal        float64          `json:"memTotal"`
	MemUsed         float64          `json:"memUsed"`
	MemFree         float64          `json:"memFree"`
	MemUsage        float64          `json:"memUsage"`
	GoTotal         uint64           `json:"goTotal"`
	GoUsed          float64          `json:"goUsed,"`
	SysComputerIp   string           `json:"sysComputerIp,"`
	SysComputerName string           `json:"sysComputerName"`
	SysOsName       string           `json:"sysOsName"`
	SysOsArch       string           `json:"sysOsArch"`
	GoName          string           `json:"goName"`
	GoVersion       string           `json:"goVersion,"`
	GoStartTime     string           `json:"goStartTime"`
	GoRunTime       int64            `json:"goRunTime"`
	GoHome          string           `json:"goHome,"`
	GoUserDir       string           `json:"goUserDir"`
	DiskList        []disk.UsageStat `json:"diskList"`
}

var StartTime = gtime.Datetime()

func NewServer() *server {
	server := new(server)
	server.CpuNum, _ = cpu.Counts(false)      //cpu物理核心
	server.CpuNumThread, _ = cpu.Counts(true) //核心数

	cpuInfo, err := cpu.Percent(time.Duration(time.Second), false)
	if err == nil {
		server.CpuUsed, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", cpuInfo[0]), 64)
	}

	loadInfo, err := load.Avg()
	if err == nil {
		server.CpuAvg5, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", loadInfo.Load5), 64)
		server.CpuAvg15, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", loadInfo.Load15), 64)
	}

	v, err := mem.VirtualMemory()
	if err == nil {
		server.MemTotal = gconv.Float64(fmt.Sprintf("%.2f", gconv.Float64(v.Total)/1024/1024/1024))
		server.MemFree = gconv.Float64(fmt.Sprintf("%.2f", gconv.Float64(v.Free)/1024/1024/1024))
		server.MemUsed = gconv.Float64(fmt.Sprintf("%.2f", gconv.Float64(v.Used)/1024/1024/1024))
		server.MemUsage = gconv.Float64(fmt.Sprintf("%.2f", v.UsedPercent))
	}

	var gomem runtime.MemStats
	runtime.ReadMemStats(&gomem)
	server.GoUsed = gconv.Float64(fmt.Sprintf("%.2f", gconv.Float64(gomem.Sys)/1024/1024/1024))

	ip, err := ipUtils.GetLocalIP()
	if err == nil {
		server.SysComputerIp = ip
	}

	sysInfo, err := host.Info()

	if err == nil {
		server.SysComputerName = sysInfo.Hostname
		server.SysOsName = sysInfo.OS
		server.SysOsArch = sysInfo.KernelArch
	}

	server.GoName = "GoLang"             //语言环境
	server.GoVersion = runtime.Version() //版本
	gtime.Date()
	server.GoStartTime = StartTime //启动时间

	server.GoRunTime = timeUtils.GetHourDiffer(StartTime, gtime.Datetime()) //运行时长
	server.GoHome = runtime.GOROOT()                                        //安装路径

	curDir, err := os.Getwd()

	if err == nil {
		server.GoUserDir = curDir
	}

	//服务器磁盘信息

	diskInfo, err := disk.Partitions(true) //所有分区
	server.DiskList = make([]disk.UsageStat, 0, len(diskInfo))
	if err == nil {
		for _, p := range diskInfo {
			diskDetail, err := disk.Usage(p.Mountpoint)
			if err == nil {
				diskDetail.UsedPercent, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", diskDetail.UsedPercent), 64)
				diskDetail.Total = diskDetail.Total / 1024 / 1024
				diskDetail.Used = diskDetail.Used / 1024 / 1024
				diskDetail.Free = diskDetail.Free / 1024 / 1024
				server.DiskList = append(server.DiskList, *diskDetail)
			}
		}
	}
	return server
}
