package main

import (
	"errors"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

func main() {
	debug := false
	autorun := true
	modified := false
	for _, arg := range os.Args {
		if arg == "-debug" {
			debug = true
		} else if arg == "-noautorun" {
			autorun = false
		}
	}

	config, err := getConfigFile()
	if err != nil {
		panic(err)
	}

	index, line, err := findDebugLine(config)
	if err != nil {
		panic(err)
	}

	if debug {
		if strings.Contains(line, "false") {
			config[index] = "debug: 'true'"
			modified = true
		}
	} else {
		if strings.Contains(line, "true") {
			config[index] = "debug: 'false'"
			modified = true
		}
	}

	if modified {
		err := updateConfigFile(config)
		if err != nil {
			panic(err)
		}
	}

	if autorun {
		exec.Command("./altv.exe").Start()
	}
}

func getConfigFile() ([]string, error) {
	file, err := ioutil.ReadFile("altv.cfg")
	if err != nil {
		return []string{}, err
	}
	return strings.Split(string(file), "\n"), nil
}

func findDebugLine(config []string) (int, string, error) {
	index := -1
	line := ""
	for i, l := range config {
		if len(l) > 6 && l[0:6] == "debug:" {
			if index != -1 || len(line) > 0 {
				return 0, "", errors.New("multiple debug lines found in altv.cfg")
			}
			index = i
			line = l
		}
	}

	if index == -1 || len(line) > 0 {
		return 0, "", errors.New("no debug line found")
	}

	return index, line, nil
}

func updateConfigFile(config []string) error {
	var configString string
	for i, l := range config {
		configString += l
		if i != len(config)-1 {
			configString += "\n"
		}
	}
	err := ioutil.WriteFile("altv.cfg", []byte(configString), 0644)
	return err
}
