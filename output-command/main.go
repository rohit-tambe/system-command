package main

import (
	"fmt"
	"os/exec"
)

func main() {
	out, err := exec.Command("bash", "-c", "ps cax | grep node").Output()
	if err != nil {
		fmt.Printf("tomcat %s", err.Error())
	}
	fmt.Println("Command Successfully Executed")
	output := string(out[:])
	fmt.Println(output)
}
