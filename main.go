package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	//"time"
)

// make sure varname is an environmental variable.
func envOrDie(varname string) {
	val, ok := os.LookupEnv(varname)
	if !ok || val == "" {
		log.Fatalf("ERROR: missing environement variable: %s", varname)
	}
	log.Println(varname, val)
}

func main() {
	flag.Parse()
	envOrDie("STAFF_SUBMITTER_USERID")
	envOrDie("STAFF_SUBMITTER_PASSWD")
	envOrDie("ANSWER_SERVER")
	envOrDie("ANSWER_SERVER_USERID")
	envOrDie("ANSWER_SERVER_PASSWD")

	fmt.Println("Go Web App Started on Port 3000")
	serve()
}
