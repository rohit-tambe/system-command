package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // user postgres driver
)

func main() {

	// log.Printf(psqlInfo)
	// psqlUser := "postgres"
	// psqlHost := "localhost"
	// psqlDBName := "message_routing"
	// sqlFilePath := "message_routing.sql"
	// cmd := exec.Command("psql", "-U", psqlUser, "-h", psqlHost, "-d", psqlDBName, "-a", "-f", sqlFilePath)

	// var out, stderr bytes.Buffer

	// cmd.Stdout = &out
	// cmd.Stderr = &stderr

	// err := cmd.Run()
	// if err != nil {
	// 	log.Fatalf("Error executing query. Command Output: %+v\n: %+v, %v", out.String(), stderr.String(), err)
	// }
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s sslmode=disable",
		"localhost", 5432, "postgres", "root")
	log.Printf(psqlInfo)
	con, err := gorm.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err.Error())
	}
	file, err := ioutil.ReadFile("create_message_routing.sql")
	if err != nil {
		// handle error
	}

	requests := strings.Split(string(file), ";")
	for _, request := range requests {
		_ = con.Exec(request)
		// do whatever you need with result and error
	}

	psqlInfo = fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		"localhost", 5432, "postgres", "root", "message_routing")
	file, err = ioutil.ReadFile("message_routing.sql")
	if err != nil {
		// handle error
	}

	requests = strings.Split(string(file), ";")
	for _, request := range requests {
		_ = con.Exec(request)
		// do whatever you need with result and error
	}
}
