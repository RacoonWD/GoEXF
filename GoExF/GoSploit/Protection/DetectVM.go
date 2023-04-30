package Protection

import (
	"errors"
	"os"
)

func checkFileExists(filePath string) bool {
	_, error := os.Stat(filePath)
	//return !os.IsNotExist(err)
	return !errors.Is(error, os.ErrNotExist)
}
func VmFile() (bool, string) {
	generalPath := []string{
		`c:\take_screenshot.ps1`,
		`c:\loaddll.exe`,
		`c:\symbols\aagmmc.pdb`,
	}

	prlPath := []string{
		`c:\windows\system32\drivers\prleth.sys`,
		`c:\windows\system32\drivers\prlfs.sys`,
		`c:\windows\system32\drivers\prlmouse.sys`,
		`c:\windows\system32\drivers\prlvideo.sys`,
		`c:\windows\system32\drivers\prltime.sys`,
		`c:\windows\system32\drivers\prl_pv32.sys`,
		`c:\windows\system32\drivers\prl_paravirt_32.sys`,
	}

	vboxPath := []string{
		`c:\windows\system32\drivers\VBoxMouse.sys`,
		`c:\windows\system32\drivers\VBoxGuest.sys`,
		`c:\windows\system32\drivers\VBoxSF.sys`,
		`c:\windows\system32\drivers\VBoxVideo.sys`,
		`c:\windows\system32\vboxdisp.dll`,
		`c:\windows\system32\vboxhook.dll`,
		`c:\windows\system32\vboxmrxnp.dll`,
		`c:\windows\system32\vboxogl.dll`,
		`c:\windows\system32\vboxoglarrayspu.dll`,
		`c:\windows\system32\vboxoglcrutil.dll`,
		`c:\windows\system32\vboxoglerrorspu.dll`,
		`c:\windows\system32\vboxoglfeedbackspu.dll`,
		`c:\windows\system32\vboxoglpackspu.dll`,
		`c:\windows\system32\vboxoglpassthroughspu.dll`,
		`c:\windows\system32\vboxservice.exe`,
		`c:\windows\system32\vboxtray.exe`,
		`c:\windows\system32\VBoxControl.exe`,
	}

	vmwarePath := []string{
		`c:\windows\system32\drivers\vmmouse.sys`,
		`c:\windows\system32\drivers\vmnet.sys`,
		`c:\windows\system32\drivers\vmxnet.sys`,
		`c:\windows\system32\drivers\vmhgfs.sys`,
		`c:\windows\system32\drivers\vmx86.sys`,
		`c:\windows\system32\drivers\hgfs.sys`,
	}

	virtualpcPath := []string{
		`c:\windows\system32\drivers\vmsrvc.sys`,
		`c:\windows\system32\drivers\vpc-s3.sys`,
	}

	allPath := [][]string{virtualpcPath, prlPath, vmwarePath, vboxPath, generalPath}

	for _, paths := range allPath {
		for _, p := range paths {
			if checkFileExists(p) {
				return true, p
			}
		}
	}

	return false, "none"
}
