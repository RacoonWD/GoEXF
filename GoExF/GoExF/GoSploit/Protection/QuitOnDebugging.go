package Protection

import "os"

func QuitOnDebugging() {
	go func() {
		if Debuging() {
			os.Exit(1)
		}
	}()
}
