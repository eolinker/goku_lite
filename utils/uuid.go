package utils

import (
	"log"
	"os/exec"
	"strings"
)

//TimeUUID 获取uuid
func TimeUUID() string {
	out, err := exec.Command("uuidgen").Output()
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(string(out), "\n", "", -1)
}
