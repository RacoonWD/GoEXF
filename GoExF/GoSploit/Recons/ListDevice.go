package Recons

import (
	"os/exec"
	"strings"
)

type Device struct {
	Name   string
	Status string
	Class  string
}

func ListDevice() []Device {
	//create a new array of devices
	var devices []Device

	//start power shell and scan for Devices
	out, err := exec.Command("powershell", "Get-PnpDevice").Output()
	if err != nil {
		panic(err)
	}
	powerShellOutput := string(out)
	//split the output into lines
	//remove whitespace from the beginning and end of each line
	//remove empty lines

	powerShellOutput = strings.Replace(powerShellOutput, "\r\n", "\n", -1)
	powerShellOutput = strings.Replace(powerShellOutput, "\r", "\n", -1)

	lines := strings.Split(powerShellOutput, "\n")
	//loop through the lines
	for i, line := range lines {

		line = standardizeSpaces(line)
		var device Device
		if len(line) < 3 && i < 2 {
		} else {
			//split line into parts
			parts := strings.Split(line, " ")

			//check if parts is empty
			if len(parts) > 1 {

				//feed device

				device.Class = parts[1]
				device.Status = parts[0]

				//name is the last part of the line before last space
				name := strings.Join(parts[2:], " ")
				//remove everything after the last space

				//split the name into parts
				nameParts := strings.Split(name, " ")
				//remove the last part of the name
				nameParts = nameParts[:len(nameParts)-1]
				//join the parts back together
				device.Name = strings.Join(nameParts, " ")

				//add the device to the array
				devices = append(devices, device)
			}
		}

	}

	return devices
}

func GetDeviceByClass(class string) []Device {
	var devices []Device
	for _, v := range ListDevice() {
		if v.Class == class {
			devices = append(devices, v)
		}
	}
	return devices
}

func GetDeviceByName(name string) []Device {
	var devices []Device
	for _, v := range ListDevice() {
		if strings.Contains(v.Name, name) {
			devices = append(devices, v)
		}
	}
	return devices
}

func standardizeSpaces(s string) string {
	return strings.Join(strings.Fields(s), " ")
}
