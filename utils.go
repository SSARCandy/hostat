package main

import (
    "fmt"
    "os/exec"
    "strconv"
    "strings"

    "github.com/logrusorgru/aurora"
)

func ExpandRange(s string) ([]string, error) {
    var hosts []string
    hostList := strings.Split(s, ",")
    for _, host := range hostList {
        if strings.Contains(host, "[") {
            parts := strings.Split(host, "[")
            if len(parts) != 2 {
                return nil, fmt.Errorf("invalid expansion range: %s", host)
            }
            prefix := parts[0]
            suffix := parts[1]
            suffix = strings.TrimSuffix(suffix, "]")
            rangeParts := strings.Split(suffix, "-")
            if len(rangeParts) != 2 {
                return nil, fmt.Errorf("invalid expansion range: %s", host)
            }
            start, err := strconv.Atoi(rangeParts[0])
            if err != nil {
                return nil, fmt.Errorf("invalid expansion range: %s", host)
            }
            end, err := strconv.Atoi(rangeParts[1])
            if err != nil {
                return nil, fmt.Errorf("invalid expansion range: %s", host)
            }
            if start > end {
                return nil, fmt.Errorf("invalid expansion range: %s", host)
            }
            for i := start; i <= end; i++ {
                hosts = append(hosts, fmt.Sprintf("%s%02d", prefix, i))
            }
        } else {
            hosts = append(hosts, host)
        }
    }
    return hosts, nil
}

func RedScale(v float64, thres int) aurora.Value {
	if v >= float64(thres) {
		return aurora.BrightRed(v).Bold()
	}
	return aurora.Reset(v)
}

func PrintSlurmInfo(nodename string) {
	cmd := fmt.Sprintf("sinfo -o '%%N %%.6D %%P %%6t %%c' -N | grep %s | awk '{print $4}'", nodename)
	out, _ := exec.Command("bash", "-c", cmd).Output()
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
	cmd := fmt.Sprintf("squeue -o '%%u %%R' -h | awk '$2==\"%s\" {print $2\" \"$1}'  | sort | uniq -c", nodename)
	out, _ := exec.Command("bash", "-c", cmd).Output()
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
