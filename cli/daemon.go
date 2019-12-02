package main

import (
	"currencyParser/command"
	"currencyParser/service/config"
	"currencyParser/service/mainDatabase"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

func main() {
	defer mainDatabase.Close()

	if len(os.Args) <= 1 {
		panic("set command name")
	}

	commandName := os.Args[1]
	if commandName == "" {
		panic("set command name")
	}


	command.ArgumentUtil{}.ExceptArgument(&os.Args, 1)

	var interval int
	intervalString, err := command.ArgumentUtil{}.GetFlag("interval")
	if err != nil {
		panic(err)
	}
	interval, err = strconv.Atoi(intervalString)
	if interval == 0 {
		panic("Set interval")
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
				panic(err)
			}
		}

	}
}

func handleDaemonShutDown(signalChan chan os.Signal) {
	log.Print("Caught signal", map[string]interface{}{
		"signal":    signalChan,
		"operation": "terminating service",
	})

	mainDatabase.Close()
}