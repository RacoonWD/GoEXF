package Protection

import (
	"os/exec"
	"strconv"
	"strings"
)

type Process struct {
	Name   string
	Pid    int
	Memory int
}

func standardizeSpaces(s string) string {
	return strings.Join(strings.Fields(s), " ")
}

func getRunningProcess() []Process {
	processList := []Process{}
	Output, _ := exec.Command("powershell", "Get-Process").Output()
	raw := strings.Split(string(Output), "\n")
	indexName := 0
	indexPid := 0
	for _, v := range raw {
		v = strings.Replace(v, "\n", "", -1)
		v = standardizeSpaces(v)
		line := strings.Split(v, " ")
		// fmt.Println(line)
		pr := Process{}

		for i, arg := range line {
			if len(arg) > 1 {
				if arg == "ProcessName" {
					indexName = i
				}
				if arg == "Id" {
					indexPid = i
				}
				if i == indexName {
					pr.Name = arg + ".exe"
				}
				if i == indexPid {
					pr.Pid, _ = strconv.Atoi(arg)
				}
			}

		}
		if pr.Pid != 0 {

			processList = append(processList, pr)
		}

	}
	return processList
}
