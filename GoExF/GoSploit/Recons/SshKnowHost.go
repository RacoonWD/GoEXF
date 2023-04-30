package Recons

import (
	"os/exec"
	"strings"
)

type SshHost struct {
	Host       string
	Protection string
}

func SshKnowHost() []SshHost {
	var hosts []SshHost

	out, err := exec.Command("powershell", "Get-Content", "$env:USERPROFILE/.ssh/known_hosts").Output()
	if err != nil {
		panic(err)
	}
	powerShellOutput := string(out)
	//split by line
	powerShellOutput = strings.Replace(powerShellOutput, "\r\n", "\n", -1)
	powerShellOutput = strings.Replace(powerShellOutput, "\r", "\n", -1)
	lines := strings.Split(powerShellOutput, "\n")
	//loop through the lines
	for _, line := range lines {
		line = standardizeSpaces(line)
		var host SshHost
		if len(line) < 3 {
		} else {
			//split line into parts
			parts := strings.Split(line, " ")
			//check if parts is empty
			if len(parts) > 1 {
				//feed device
				host.Host = parts[0]
				host.Protection = parts[1]
				hosts = append(hosts, host)
			}
		}
	}
	return hosts
}
