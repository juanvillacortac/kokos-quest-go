package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

var output string

func init() {
	if len(os.Args) > 1 {
		output = os.Args[1]
	} else {
		fmt.Printf("USAGE: \n\t%s [FILE]\n", os.Args[0])
		os.Exit(-1)
	}

}

func LoadLevelHandler() ([]byte, error) {
	_, err := os.Stat(output)
	if os.IsNotExist(err) {
		return nil, fmt.Errorf("The path/file doesn't exist")
	}
	return ioutil.ReadFile(output)
}

func SaveLevelHandler(buff []byte) error {
	return ioutil.WriteFile(output, buff, 0644)
}
