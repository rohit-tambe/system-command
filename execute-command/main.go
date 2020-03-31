package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
)

func main() {
	gopath := os.Getenv("GOPATH")
	fmt.Println(gopath)
	excePath := fmt.Sprintf("%v%v", gopath, "/src/github.com/rohit-tambe/garbage-collection/garbage.sh")
	path, err := exec.LookPath(excePath)
	if err != nil {
		log.Fatalf("didn't find  executable '%s'\n", path)
	} else {
		fmt.Printf("executable is in '%s'\n", path)
	}
	// out, err := exec.Command("/bin/bash", "/home/billcloud3/GoWorkspace/src/github.com/rohit-tambe/garbage-collection/garbage.sh").Output()
	// if err != nil {
	// 	fmt.Printf("tomcat %s", err.Error())
	// }
	// output := string(out[:])
	// fmt.Println(output)
	cmd := exec.Command("/bin/bash", path)
	var stdoutBuf, stderrBuf bytes.Buffer
	cmd.Stdout = io.MultiWriter(os.Stdout, &stdoutBuf)
	cmd.Stderr = io.MultiWriter(os.Stderr, &stderrBuf)

	err = cmd.Run()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	outStr, errStr := string(stdoutBuf.Bytes()), string(stderrBuf.Bytes())
	fmt.Printf("\nout:\n%s\nerr:\n%s\n", outStr, errStr)
}
