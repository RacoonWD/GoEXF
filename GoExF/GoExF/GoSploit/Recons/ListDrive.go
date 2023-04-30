package Recons

import (
	"os"
	"time"
)

// key string, state bool
//
//             0 -> new 1 -> remove
func OnNewDriveThreaded(fc func(key string, state int)) {
	baseDisk := ListDisk()
	go func() {
		for {
			newDisk := ListDisk()
			if len(baseDisk) < len(newDisk) {
				fc(newDisk[len(baseDisk)], 0)
			}
			if len(newDisk) < len(baseDisk) {
				fc(baseDisk[len(newDisk)], 1)
			}
			baseDisk = newDisk
			time.Sleep(1 * time.Second)
		}
	}()
}

func OnNewDrive(fc func(key string, state int)) {
	baseDisk := ListDisk()
	for {
		newDisk := ListDisk()
		if len(baseDisk) < len(newDisk) {
			fc(newDisk[len(baseDisk)], 0)
		}
		if len(newDisk) < len(baseDisk) {
			fc(baseDisk[len(newDisk)], 1)
		}
		baseDisk = newDisk
		time.Sleep(1 * time.Second)
	}
}

func ListDisk() []string {
	listDrivel := []string{}
	for _, letter := range "abcdefghijklmnopqrstuvwxyz" {
		_, err := os.Stat(string(letter) + ":")
		if os.IsNotExist(err) {
		} else {
			listDrivel = append(listDrivel, string(letter)+":")
		}

	}
	return listDrivel
}
