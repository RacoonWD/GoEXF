package Regedit

import "os/exec"

type keyType struct {
	name string
}

type Key struct {
	Name    string
	Path    string
	Value   string
	KeyType keyType
}

var REG_MULTI_SZ = keyType{name: "REG_MULTI_SZ"}
var REG_SZ = keyType{name: "REG_SZ"}
var REG_DWORD = keyType{name: "REG_DWORD"}
var REG_QWORD = keyType{name: "REG_QWORD"}
var REG_BINARY = keyType{name: "REG_BINARY"}

func (k *Key) CreateKey() {
	exec.Command("reg.exe", "add", k.Path, "/v", ""+k.Name+"", "/t", k.KeyType.name, "/d", k.Value).Output()

}

func (k *Key) DeleteKey() {

}
