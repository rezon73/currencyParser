package main

import (
	"currencyParser/command"
	"currencyParser/service/config"
	"currencyParser/service/logService"
	"currencyParser/service/mainDatabase"
	"os"
	"strings"
)

func main() {
	defer mainDatabase.Close()

	commandName := os.Args[1]
	if commandName == "" {
		logService.Fatal("set command name")
	}

	logService.SetJobName("cron-" + strings.ReplaceAll(strings.Join(os.Args[1:], "_"), "-", ""))

	command.ArgumentUtil{}.ExceptArgument(&os.Args, 1)

	err := command.Factory{
		CommandName:  commandName,
		Config:       config.GetConfig(),
		MainDatabase: mainDatabase.GetInstance(0),
	}.CreateCommand().Exec()

	if err != nil {
		logService.Error(err)
	}
}
