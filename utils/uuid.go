
package utils

import (
	"strings"
    "log"
    "os/exec"
)
 
func TimeUUID() string {
    out, err := exec.Command("uuidgen").Output()
    if err != nil {
        log.Fatal(err)
    }
    return strings.Replace(string(out),"\n","",-1)
}