package command

import (
	"errors"
	"os"
	"regexp"
)

type ArgumentUtil struct {}

func (util ArgumentUtil) ExceptArgument(arguments *[]string, argumentIndex int) {
	var cleanedArgs []string
	if argumentIndex == 1 {
		cleanedArgs = append(cleanedArgs, (*arguments)[0])
	} else if argumentIndex > 0 {
		cleanedArgs = append(cleanedArgs, (*arguments)[:argumentIndex]...)
	}
	if argumentIndex < len(*arguments) - 1 {
		cleanedArgs = append(cleanedArgs, (*arguments)[argumentIndex+1:]...)
	}

	*arguments = nil
	*arguments = cleanedArgs
}

func (util ArgumentUtil) GetFlag(flagKey string) (string, error) {
	for key, arg := range os.Args {
		if found, _ := regexp.MatchString(`^[-]{0,}` + flagKey + `[=\- ]{1,}["]{0,}(\w)["]{0,}`, arg); !found {
			continue
		}

		reg := regexp.MustCompile(`^[-]{1,}` + flagKey + `[=\- ]{1,}["]{0,}(\w)["]{0,}`)
		value := reg.FindStringSubmatch(arg)

		util.ExceptArgument(&os.Args, key)

		return value[1], nil
	}

	return ``, errors.New("Not found")
}
