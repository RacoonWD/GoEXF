package Recons

import (
	"io/ioutil"
	"log"
)

func ListDir(path string) []string {
	listFile := []string{}
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}
	for _, v := range files {
		listFile = append(listFile, v.Name())

	}
	return listFile
}
