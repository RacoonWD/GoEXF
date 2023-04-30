package Visibility

import (
	"os/exec"
)

func HideConsole() {
	cmd := exec.Command("cmd", "/C", "start", "/MIN", "cmd", "/C", "exit")
	cmd.Start()
}
