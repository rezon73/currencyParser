package main

import (
	"currencyParser/command"
	"currencyParser/service/config"
	"currencyParser/service/logService"
	"currencyParser/service/mainDatabase"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
	"time"
)

func main() {
	defer mainDatabase.Close()

	if len(os.Args) <= 1 {
		logService.Fatal("set command name")
	}

	commandName := os.Args[1]
	if commandName == "" {
		logService.Fatal("set command name")
	}

	logService.SetJobName("daemon-" + strings.ReplaceAll(strings.Join(os.Args[1:], "_"), "-", ""))

	logService.Info("Start signal")

	command.ArgumentUtil{}.ExceptArgument(&os.Args, 1)

	var interval int
	intervalString, err := command.ArgumentUtil{}.GetFlag("interval")
	if err != nil {
		logService.Fatal(err)
	}
	interval, err = strconv.Atoi(intervalString)
	if interval == 0 {
		logService.Fatal("Set interval")
	}

	comm := command.Factory{
		CommandName:  commandName,
		Config:       config.GetConfig(),
		MainDatabase: mainDatabase.GetInstance(0),
	}.CreateCommand()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

LOOP:
	for {
		select {
		case <- sigChan:
			handleDaemonShutDown(sigChan)
			break LOOP
		default:
			time.Sleep(time.Duration(interval) * time.Second)
			err = comm.Exec()
			if err != nil {
				logService.Error(err)
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