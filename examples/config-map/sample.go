package main

import (
	"log"

	"github.com/kimprado/sllog/pkg/logging"
)

func main() {

	levels := map[string]string{
		"ROOT":                      "ERROR",
		"levelA":                    "INFO",
		"levelA.levelB.CONCURRENCE": "TRACE",
	}

	logger := logging.NewLogger("", levels)                                 // Using ROOT level
	loggerA := logging.NewLogger("levelA", levels)                          // Using levelA level
	loggerBConcur := logging.NewLogger("levelA.levelB.CONCURRENCE", levels) // Using levelA.levelB.CONCURRENCE level
	loggerC := logging.NewLogger("levelA.levelC", levels)                   //Using levelA level by inheritance "{levelA}.levelC"

	logger.Errorf("message 1")
	logger.Infof("message 3")  //no show. Only ERROR
	logger.Debugf("message 4") //no show. Only ERROR

	log.Print()
	loggerA.Errorf("message 1")
	loggerA.Warnf("message 2")
	loggerA.Infof("message 3")
	loggerA.Tracef("message 4") //no show. Only INFO or above

	log.Print()
	loggerBConcur.Errorf("message 1")
	loggerBConcur.Debugf("message 2")

	log.Print()
	loggerC.Errorf("message 1") //Using levelA level by inheritance
	loggerC.Warnf("message 2")  //Using levelA level by inheritance
	loggerC.Infof("message 3")  //Using levelA level by inheritance
	loggerC.Debugf("message 4") //no show. //Using levelA level by inheritance
	loggerC.Tracef("message 5") //no show. //Using levelA level by inheritance
}
