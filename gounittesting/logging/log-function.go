package logging

import (
	"fmt"
	"log"
	"os"
)

func handleClose(f *os.File) {
	err := f.Close()
	if err != nil {
		fmt.Printf("error closing file: %v", err)
	}
}

func LogStuff(logName, logInputData string ) {
	f, err := os.OpenFile(logName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("error opening file: %v", err)
	}
	defer handleClose(f)

	log.SetOutput(f)

	log.Printf("%v", logInputData)
}
