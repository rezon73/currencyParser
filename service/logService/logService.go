package logService

import (
	"currencyParser/service/telegramSender"
	"fmt"
	"log"
	"log/syslog"
)

var logInstance *syslog.Writer

func init() {
	initClient("index")
}

func SetJobName(jobName string) {
	logInstance.Close()
	initClient(jobName)
}

func Info(logMessage ...interface{}) {
	log.Println(fmt.Sprintf("Info: %s", logMessage))
}

func Warn(logMessage ...interface{}) {
	log.Println(fmt.Sprintf("Warning: %s", logMessage))
}

func Error(logMessage ...interface{}) {
	log.Println(fmt.Sprintf("Error: %s", logMessage))
}

func Fatal(logMessage ...interface{}) {
	msg := fmt.Sprintf("Fatal: %s", logMessage)
	fmt.Println(msg)
	log.Fatalln(msg)
}

func Alert(logMessage ...interface{}) {
	msg := fmt.Sprintf("Alert: %s", logMessage)
	log.Println(msg)
	telegramSender.GetSender().SendMessage(msg)
}

func initClient(jobName string) {
	var err error
	logInstance, err = syslog.New(syslog.LOG_NOTICE, "[" + jobName + "]")
	if err == nil {
		log.SetOutput(logInstance)
	}
}