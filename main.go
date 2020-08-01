package main

import (
	"fmt"

  "github.com/fatih/color"
  "github.com/shirou/gopsutil/mem"
  "github.com/shirou/gopsutil/load"
)

/**
 * hostname  | CPUs | 1m Load | 5m Load | 15m Load | memory % | disk % | users | uptime | status | jobs
 * cluster01 |    8 |     1.1 |     1.5 |      2.1 |     60 % |   56 % |     2 | 16days |  alloc | ssarcandy(8)
 * cluster02 |    8 |     5.1 |     5.5 |      5.1 |     20 % |   96 % |     0 | 16days |  alloc | ssarcandy(8)
 */

func main() {
  red := color.New(color.FgRed).SprintFunc()
  m, _ := mem.VirtualMemory()
  l, _ := load.Avg()

  fmt.Printf("Total: %v, Free:%v, UsedPercent:%f%%\n", red(m.Total), m.Free, m.UsedPercent)
  fmt.Printf("1m: %v, 5m: %v, 15m: %v\n", l.Load1, l.Load5, l.Load15)
}
