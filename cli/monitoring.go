package main

import (
	"currencyParser/monitoring"
	"currencyParser/service/config"
	"currencyParser/service/logService"
	"currencyParser/service/mainDatabase"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const EXECUTE_EACH_SECONDS = 5

func main() {
	defer mainDatabase.Close()

	logService.SetJobName("monitor")
	logService.Info("Start monitor")

	var monitors []monitoring.Monitor
	monitors = append(monitors, monitoring.LastUpdateMonitor{
		Config:          config.GetConfig(),
		MainDatabase:    mainDatabase.GetInstance(0),
		DelayErrorLevel: 10,
	})

	var errs []error

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

LOOP:
	for {
		select {
		case <- sigChan:
			handleDaemonShutDown(sigChan)
			break LOOP
		default:
			time.Sleep(time.Duration(EXECUTE_EACH_SECONDS) * time.Second)

			for _, monitor := range monitors {
				errs = append(errs, monitor.Check()...)
			}

			for _, err := range errs {
				logService.Alert(err)
			}
		}
	}
}

func handleDaemonShutDown(signalChan chan os.Signal) {
	logService.Info("Caught signal", map[string]interface{}{
		"signal":    signalChan,
		"operation": "terminating service",
	})

	mainDatabase.Close()
}