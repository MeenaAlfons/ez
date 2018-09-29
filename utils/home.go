package utils

import (
	"os"

	homedir "github.com/mitchellh/go-homedir"
)

var HomeDir string
var BinDir string
var PrivateHomeDir string

func init() {
	HomeDir, err := homedir.Dir()
	if err != nil {
		panic(err)
	}
	BinDir = HomeDir + "/.ez/bin/"
	PrivateHomeDir = HomeDir + "/.ez/home/"

	err = os.MkdirAll(BinDir, os.ModePerm)
	if err != nil {
		panic(err)
	}

	err = os.MkdirAll(PrivateHomeDir, os.ModePerm)
	if err != nil {
		panic(err)
	}
}
