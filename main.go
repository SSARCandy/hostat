package main

import (
	"fmt"
	"flag"
    "os/exec"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
)

/*
 * hostname  | CPUs | 1m Load | 5m Load | 15m Load | memory % | disk % | users | uptime | status | jobs
 * cluster01 |    8 |     1.1 |     1.5 |      2.1 |     60 % |   56 % |     2 | 16days |  alloc | ssarcandy(8)
 * cluster02 |    8 |     5.1 |     5.5 |      5.1 |     20 % |   96 % |     0 | 16days |  alloc | ssarcandy(8)
 */

func main() {
	m, _ := mem.VirtualMemory()
	l, _ := load.Avg()
	c, _ := cpu.Counts(true)
	i, _ := host.Info()
	d, _ := disk.Usage("/")
	t, _ := host.Uptime()

	header := flag.Bool("header", true, "Print Header or not")
	thresMemory := flag.Int("thres_mem", 80, "Threshold for Memory. Render red color if >= thres")
	thresDisk := flag.Int("thres_disk", 80, "Threshold for Disk. Render red color if >= thres")
	thresLoad := flag.Int("thres_load", c, "Threshold for Load. Render red color if >= thres")
	hosts := flag.String("hosts", "localhost", "Target hosts in range expression, i.e. host[01-99]")
	flag.Parse()

	_, err := exec.LookPath("sinfo")

	if *header {
		fmt.Printf("%-10s |%5s |%7s |%7s |%7s |%9s |%7s |%7s |", "hostname", "CPUs", "1m", "5m", "15m", "memory %", "disk %", "UpTime")
		if err == nil {
			fmt.Printf("%6s | %s", "State", "Jobs")
		}
		fmt.Println("")
	}

	fmt.Printf("%-10s |", i.Hostname)
	fmt.Printf("%5v |", c)
	fmt.Printf("%7.1f |%7.1f |%7.1f |", RedScale(l.Load1, *thresLoad), RedScale(l.Load5, *thresLoad), RedScale(l.Load15, *thresLoad))
	fmt.Printf("%7.0f %% |", RedScale(m.UsedPercent, *thresMemory))
	fmt.Printf("%5.0f %% |", RedScale(d.UsedPercent, *thresDisk))
	fmt.Printf("%7s |", fmt.Sprintf("%v d", t/86400))

	if err == nil {
		PrintSlurmInfo(i.Hostname)
		PrintSlurmQueue(i.Hostname)
	}

	fmt.Println("")

	// Example usage
	strings, err := ExpandRange(*hosts)
	if err != nil {
		fmt.Printf("error: %s\n", err)
	} else {
		fmt.Println(strings)
	}
}

