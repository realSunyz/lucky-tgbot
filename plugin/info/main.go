package info

import (
	"errors"
	"fmt"

	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/docker"
	"github.com/shirou/gopsutil/v4/host"
	"github.com/shirou/gopsutil/v4/mem"
	tele "gopkg.in/telebot.v3"
	"log"
)

func getCPUInfo() string {
	// Get CPU model name
	cpuName := func() string {
		cpuInfo, err := cpu.Info()
		if err != nil || len(cpuInfo) == 0 {
			log.Printf("Error fetching CPU info: %s", err)
			return "N/A"
		}
		return cpuInfo[0].ModelName
	}

	// Get CPU usage percentage
	cpuPercent := func() string {
		cpuUsage, err := cpu.Percent(0, false)
		if err != nil {
			log.Printf("Error fetching CPU usage: %s", err)
			return "N/A"
		}
		return fmt.Sprintf("%.2f%%", cpuUsage[0])
	}

	cpuInfo := fmt.Sprintf(
		"CPU Model: `%s`\nCPU Usage: `%s`\n", cpuName(), cpuPercent())

	return cpuInfo
}

func getRAMInfo() string {
	vmInfo := func() (string, string, string) {
		vmStat, err := mem.VirtualMemory()
		if err != nil || vmStat == nil {
			log.Printf("Error fetching memory info: %s", err)
			return "N/A", "N/A", "N/A"
		}
		vmPercent := fmt.Sprintf("%.2f%%", vmStat.UsedPercent)
		vmTotal := fmt.Sprintf("%.2f", float64(vmStat.Total)/1e9)
		vmUsed := fmt.Sprintf("%.2f", float64(vmStat.Used)/1e9)
		return vmPercent, vmUsed, vmTotal
	}

	vmPercent, vmUsed, vmTotal := vmInfo()

	ramInfo := fmt.Sprintf(
		"RAM Usage: `%s` (`%sGB` / `%sGB`)\n", vmPercent, vmUsed, vmTotal)

	return ramInfo
}

func getHostInfo() string {
	hostName := func() string {
		hostInfo, err := host.Info()
		if err != nil || hostInfo == nil {
			log.Printf("Error fetching host info: %s", err)
			return "N/A"
		}
		return hostInfo.Hostname
	}

	hostInfo := fmt.Sprintf(
		"Hostname: `%s`\n", hostName())

	return hostInfo
}

func getDocker() string {
	checkDocker := func() string {
		dockerStat, err := docker.GetDockerStat()
		if err == nil && len(dockerStat) > 0 {
			return "Yes"
		} else if errors.Is(err, docker.ErrDockerNotAvailable) {
			return "No"
		} else {
			log.Printf("Error fetching container info: %s", err)
			return "N/A"
		}
	}

	dockerInfo := fmt.Sprintf("Running inside a container: `%s`\n", checkDocker())

	return dockerInfo
}

func Execute(c tele.Context) error {
	cpuInfo := getCPUInfo()
	ramInfo := getRAMInfo()
	hostInfo := getHostInfo()
	dockerInfo := getDocker()

	outputText := fmt.Sprintf("%s%s%s%s", hostInfo, cpuInfo, ramInfo, dockerInfo)

	return c.Reply(outputText, &tele.SendOptions{
		ParseMode: "Markdown",
	})
}
