package Protection

import (
	"github.com/F-r-o-i-d/GoExF/GoExF/GoSploit"
)

var ProcessListToCheck = []string{"processhacker.exe", "procmon.exe", "pestudio.exe", "procmon64.exe", "x32dbg.exe", "x64dbg.exe", "CFF Explorer.exe", "procexp64.exe",
	"procexp.exe", "pslist.exe", "tcpview.exe", "tcpvcon.exe", "dbgview.exe", "RAMMap.exe", "RAMMap64.exe", "vmmap.exe", "ollydbg.exe", "agent.py", "autoruns.exe",
	"autorunsc.exe", "filemon.exe", "regmon.exe", "idaq.exe", "idaq64.exe", "ImmunityDebugger.exe", "Wireshark.exe", "dumpcap.exe", "HookExplorer.exe", "ImportREC.exe", "PETools.exe",
	"LordPE.exe", "SysInspector.exe", "proc_analyzer.exe", "sysAnalyzer.exe", "sniff_hit.exe", "windbg.exe", "joeboxcontrol.exe", "joeboxserver.exe", "joeboxserver.exe", "ResourceHacker.exe", "Fiddler.exe",
	"httpdebugger.exe"}

func Debuging() bool {

	for _, process := range getRunningProcess() {
		for _, target := range ProcessListToCheck {
			if target == process.Name {
				return true
			}
		}
	}
	flag, _, _ := GoSploit.IsDebuggerPresent.Call()
	if flag != 0 {
		return true
	} else {
		return false
	}

}
