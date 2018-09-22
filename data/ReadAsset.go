package data

import (
	"fmt"
	"io"
)

func ReadAsset(path string) (string, error) {
	file, err := Assets.Open(path)

	if err != nil {
		fmt.Println("Error opening file", err)
		return "", err
	}

	fileinfo, err := file.Stat()
	if err != nil {
		fmt.Println("Error gettting State of file", err)
		return "", err
	}

	filesize := fileinfo.Size()

	buffer := make([]byte, filesize)

	readbytes, err := file.Read(buffer)
	if err != nil && err != io.EOF {
		fmt.Println("Error Reading file. readbytes=", readbytes)
		fmt.Println(err)
		return "", err
	}

	return string(buffer), nil
}
