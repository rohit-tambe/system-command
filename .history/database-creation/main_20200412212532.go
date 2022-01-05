package main

import (
	"bytes"
	"log"
	"os/exec"
)

func main() {
	psqlUser := "postgres"
	psqlHost := "localhost"
	psqlDBName := "message_routing"
	cmd := exec.Command("psql", "-U", psqlUser, "-h", psqlHost, "-d", psqlDBName, "-a", "-f", sqlFilePath)

	var out, stderr bytes.Buffer

	cmd.Stdout = &out
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		log.Fatalf("Error executing query. Command Output: %+v\n: %+v, %v", out.String(), stderr.String(), err)
	}
}
