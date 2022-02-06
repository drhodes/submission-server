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

func setup_logger() *os.File {
	// log to custom file
	LOG_FILE := "./submitter_log"
	// open log file
	logFile, err := os.OpenFile(LOG_FILE, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Panic(err)
	}

	// Set log out put and enjoy :)
	log.SetOutput(logFile)

	// optional: log date-time, filename, and line number
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	log.Println("Logging to custom file")
	return logFile
}

func main() {
	logFile := setup_logger()
	defer logFile.Close()
	flag.Parse()
	envOrDie("STAFF_SUBMITTER_USERID")
	envOrDie("STAFF_SUBMITTER_PASSWD")
	envOrDie("ANSWER_SERVER")
	envOrDie("ANSWER_SERVER_USERID")
	envOrDie("ANSWER_SERVER_PASSWD")

	fmt.Println("Go Web App Started on Port 3000")
	serve()

}
