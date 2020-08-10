package main

import (
	"fmt"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/disk"
)

/**
 * hostname  | CPUs | 1m Load | 5m Load | 15m Load | memory % | disk % | users | uptime | status | jobs
 * cluster01 |    8 |     1.1 |     1.5 |      2.1 |     60 % |   56 % |     2 | 16days |  alloc | ssarcandy(8)
 * cluster02 |    8 |     5.1 |     5.5 |      5.1 |     20 % |   96 % |     0 | 16days |  alloc | ssarcandy(8)
 */

func main() {
	// red := color.New(color.FgRed).SprintFunc()
	m, _ := mem.VirtualMemory()
	l, _ := load.Avg()
	c, _ := cpu.Counts(true)
	i, _ := host.Info()
	d, _ := disk.Usage("/")

	// print header
	fmt.Printf("%-10s |%5s |%7s |%7s |%7s |%8s |%6s |\n", "hostname", "CPUs", "1m", "5m", "15m", "memory%", "disk%")

	fmt.Printf("%-10s |", i.Hostname)
	fmt.Printf("%5v |", c)
	fmt.Printf("%7.1f |%7.1f |%7.1f |", l.Load1, l.Load5, l.Load15)
	fmt.Printf("%7.1f%% |", m.UsedPercent)
	fmt.Printf("%5.1f%% |\n", d.UsedPercent)
}
