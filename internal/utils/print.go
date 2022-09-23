package utils

import (
	"encoding/json"
	"fmt"
)


func PrettyPrint(v interface{}) error {
	return printPrettyJson(v)
}

func printPrettyJson(v interface{}) error {
	b, err := json.MarshalIndent(v, "", "  ")
	if err == nil {
		fmt.Println(string(b))
	}
	return err
}
