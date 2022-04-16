package main

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
)

var cmdToGetClass = `xprop -id $(xprop -root 32x '\t$0' _NET_ACTIVE_WINDOW | cut -f 2) WM_CLASS | awk -F "\"" '{print $4}'`

func main() {

	getWindowId := exec.Command("xprop", "-root", "32x", "'\t$0'", "_NET_ACTIVE_WINDOW")
	xpropOutput, err := getWindowId.Output()
	if err != nil {
		log.Fatalln("Can't get current class: ", err.Error())
	}
	windowFields := strings.Fields(string(xpropOutput))

	if len(windowFields) < 2 {
		log.Fatalln("Can't get current class: Not enough fields in the output")
	}

	windowId := windowFields[1]

	cmd := exec.Command("xprop", "-id", windowId, "WM_CLASS")
	currentClass, _ := cmd.Output()
	if err != nil {
		log.Fatalln("Can't get current class: ", err.Error())
	}

	fmt.Println(string(currentClass))
}
