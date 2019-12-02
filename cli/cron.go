package main

import (
	"currencyParser/command"
	"currencyParser/service/config"
	"currencyParser/service/mainDatabase"
	"os"
)

func main() {
	defer mainDatabase.Close()

	args := os.Args[1:]
	if len(args) == 0 {
		panic("set command name")
	}

	commandName := args[0]

	command.ArgumentUtil{}.ExceptArgument(&os.Args, 1)

	err := command.Factory{
		CommandName:  commandName,
		Config:       config.GetConfig(),
		MainDatabase: mainDatabase.GetInstance(0),
	}.CreateCommand().Exec()

	if err != nil {
		panic(err)
	}
}
