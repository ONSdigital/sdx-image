package schema

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

func Read() {

	bytes := readFile()

	m := toCompleteMap(bytes)
	s := toSurvey(m)
	fmt.Print(s)

}

func readFile() []byte {
	jsonFile, err := os.Open("schemas/mbs_0106.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	bytes, _ := io.ReadAll(jsonFile)
	return bytes
}

func toCompleteMap(bytes []byte) map[string]any {
	m := map[string]any{}
	err := json.Unmarshal(bytes, &m)
	if err != nil {
		fmt.Println(err)
	}
	return m
}

func Print(schema map[string]any) {
	data, err := json.MarshalIndent(schema, "", "    ")
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}
	fmt.Println("")
	fmt.Println("--------------")
	fmt.Printf("%s\n", data)
}
