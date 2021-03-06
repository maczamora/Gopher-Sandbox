package main

import (
	"fmt"
	"log"
	"os"
)

type messageType int

const (
	INFO messageType = 0 + iota
	WARNING
	ERROR
	FATAL
)

func main() {
	fmt.Println("-*-*-*-*-*-*-*-*-*-*-* Formatting Log Output -*-*-*-*-*-*-*-*-*-*-*")
	writeLog(INFO, "this is an information message!")
	writeLog(INFO, "this is an information message!")
	writeLog(WARNING, "There is something abnormal happening!")
	writeLog(ERROR, "Error, something failed!")
	writeLog(FATAL, "something crashed")
}

func writeLog(messagetype messageType, message string) {
	file, err := os.OpenFile("log.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(file)

	switch messagetype {
	case INFO:
		logger := log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
		logger.Println(message)
	case WARNING:
		logger := log.New(file, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
		logger.Println(message)
	case ERROR:
		logger := log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
		logger.Println(message)
	case FATAL:
		logger := log.New(file, "FATAL: ", log.Ldate|log.Ltime|log.Lshortfile)
		logger.Fatal(message)
	}
}

// Some Basic Error Levels
// Information - tell the user information or verify something occured
// Warning - something bad happened but program will still run
// Error - more serious, something expected to happen didn't happen program can continue on but needs an operator
// Fatal - something very bad happened and the program needs to close immediately
