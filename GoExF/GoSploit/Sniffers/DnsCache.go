package Sniffers

import (
	"os/exec"
	"regexp"
	"strings"
)

type Recorded struct {
	Domain string
	Ip     string
}

func DnsCache() []Recorded {
	//Get Dns cache

	out, err := exec.Command("powershell", "Get-DnsClientCache").Output()
	if err != nil {
		panic(err)
	}
	//Parse output
	var records []Recorded
	for _, line := range strings.Split(string(out), "\n") {
		if strings.Contains(line, "Success") {
			Record := Recorded{}

			//remove duplicate spaces
			space := regexp.MustCompile(`\s+`)
			s := space.ReplaceAllString(line, " ")

			//split s by spaces
			splitted := strings.Split(s, " ")
			Record.Domain = splitted[1]
			Record.Ip = splitted[7]
			//append to records
			records = append(records, Record)
		}
	}
	return records
}

func HttpsCache() {
	panic("Not implemented")
}
