package Persistance

import (
	"fmt"
	"os"
	"os/exec"
)

func GetLocaleName() string {
	//get locale name with enviroment variables
	return os.Getenv("USERNAME")

}

func injectIntoOpera() {
	//inject into opera
	fmt.Println("Injecting into opera...")
	out, err := exec.Command("powershell", "Start-Process Opera").Output()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(string(out))
}
