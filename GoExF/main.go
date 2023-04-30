package main

import (
	"fmt"

	"github.com/F-r-o-i-d/GoExF/GoExF/GoSploit"
	"github.com/F-r-o-i-d/GoExF/GoExF/GoSploit/Recons"
)

func main() {
	GoSploit.Initialize()
	for _, v := range Recons.SshKnowHost() {
		fmt.Println("Name   :", v.Host)
		fmt.Println("")
	}

}
