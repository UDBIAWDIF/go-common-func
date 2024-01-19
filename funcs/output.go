package funcs

import (
	"encoding/json"
	"fmt"
	"log"
)

func LogAsJson(v any) {
	jsonStr, err := json.Marshal(v)
	if err == nil {
		log.Println(string(jsonStr))
	}
}

func LogIfError(err error) {
	if err != nil {
		log.Println(err.Error())
	}
}

func LogListAsJson(vList []any) {
	for _, v := range vList {
		jsonStr, err := json.Marshal(v)
		if err == nil {
			log.Println(jsonStr)
		}
	}
}

func PrintAsJson(v any) {
	jsonStr, err := json.Marshal(v)
	if err == nil {
		fmt.Println(string(jsonStr))
	}
}

func PrintIfError(err error) {
	if err != nil {
		fmt.Println(err.Error())
	}
}

func PrintListAsJson(vList []any) {
	for _, v := range vList {
		jsonStr, err := json.Marshal(v)
		if err == nil {
			fmt.Println(jsonStr)
		}
	}
}
