package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s sslmode=disable",
		"localhost", 5432, "postgres", "root")
	log.Printf(psqlInfo)
	psqlUser := "postgres"
	psqlHost := "localhost"
	psqlDBName := "message_routing"
	sqlFilePath := "message_routing.sql"
	cmd := exec.Command("psql", "-U", psqlUser, "-h", psqlHost, "-d", psqlDBName, "-a", "-f", sqlFilePath)

	var out, stderr bytes.Buffer

	cmd.Stdout = &out
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		log.Fatalf("Error executing query. Command Output: %+v\n: %+v, %v", out.String(), stderr.String(), err)
	}
}
