# slLog: Simple Levels Logging in Go
 Golang Logging Levels like Spring boot configuration.
 
 Defines default(ROOT), personalizeds, and loggers hierarchy.

## Installing

Install slLog by running:

```shell
go get github.com/kimprado/sllog/pkg/logging
```

Or define in Module

```shell
github.com/kimprado/sllog v1.0.0
```

## Sample

``` go
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
```

Output:

```shell
2019/04/05 09:15:42 ERROR  message 1
2019/04/05 09:15:42 
2019/04/05 09:15:42 ERROR levelA message 1
2019/04/05 09:15:42 WARN levelA message 2
2019/04/05 09:15:42 INFO levelA message 3
2019/04/05 09:15:42 
2019/04/05 09:15:42 ERROR levelA.levelB.CONCURRENCE message 1
2019/04/05 09:15:42 DEBUG levelA.levelB.CONCURRENCE message 2
2019/04/05 09:15:42 
2019/04/05 09:15:42 ERROR levelA.levelC message 1
2019/04/05 09:15:42 WARN levelA.levelC message 2
2019/04/05 09:15:42 INFO levelA.levelC message 3
```

## Project status

Production ready
