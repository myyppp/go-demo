package logger_test

import (
	"log"
	"testing"

	"github.com/myyppp/go-demo/logger"
)

func Test_Logger(t *testing.T) {
	log.Println("message")
	// 调用 Println() 后调用 os.Exit(1)
	// log.Fatalln("fatal message")
	// 调用 Println() 后调用 panic()
	// log.Panicln("panic message")

	logger.Trace.Println("I have something standard to say")
	logger.Info.Println("Special Information")
	logger.Warning.Println("There is something you need to know about")
	logger.Error.Println("Something has failed")
}
