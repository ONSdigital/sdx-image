package controller

import (
	"fmt"
	"io"
	"os"
)

func readFile() []byte {
	jsonFile, err := os.Open("example/mbs_0106.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	bytes, _ := io.ReadAll(jsonFile)
	return bytes
}
