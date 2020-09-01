package main

import (
	"strings"
	"fmt"
	"os/exec"
	
    "flag"

	"github.com/logrusorgru/aurora"

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
}

func RedScale(v float64, thres int) aurora.Value {
	if v >= float64(thres) {
		return aurora.BrightRed(v).Bold()
	}
	return aurora.Reset(v)
}

func PrintSlurmInfo(nodename string) {
	cmd := fmt.Sprintf("sinfo -o '%%N %%.6D %%P %%6t %%c' -N | grep %s | awk '{print $4}'", nodename)
	out, _ := exec.Command("bash","-c",cmd).Output()
	state := strings.TrimSpace(string(out))

	color := aurora.Reset
	if state == "idle" || state == "mix" {
		color = aurora.BrightGreen
	} else if state == "drain" || state == "comp" {
		color = aurora.BrightBlack
	} else if strings.Contains(state, "*") {
		color = aurora.BrightRed
	}
	fmt.Printf("%6s |", color(state).Bold())
}

func PrintSlurmQueue(nodename string) {
	cmd := fmt.Sprintf("squeue -l | tail -n +3 | awk '$9 == \"%s\" {print $9\" \"$4}' | sort | uniq -c", nodename)
	out, _ := exec.Command("bash","-c",cmd).Output()
	lines := strings.Split(strings.TrimSpace(string(out)), "\n")

	var jobs []string
	for _, line := range lines {
		tokens := strings.Split(strings.TrimSpace(string(line)), " ")
		if len(tokens) < 3 {
			return
		}
		job := fmt.Sprintf("%s(%s)", tokens[2], tokens[0])
		jobs = append(jobs, job)
	}
	fmt.Printf(" %s", strings.Join(jobs, ", "))
}
