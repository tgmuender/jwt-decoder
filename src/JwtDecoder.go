package main

import (
	"bufio"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func main() {

	args := os.Args[1:]

	var token string
	if len(args) == 0 {
		// No args passed, read from stdin
		token = readInput()
	} else {
		token = os.Args[1]
	}

	var parts []string = strings.Split(token, ".")

	for i, v := range parts {
		if i == 2 {
			continue // skip signature part
		}
		v = addPadding(v)

		decoded, error := base64.URLEncoding.DecodeString(string(v))
		if error != nil {
			fmt.Println("Decoding error ", error)
			return
		}

		var objmap interface{}
		error = json.Unmarshal(decoded, &objmap)

		if error != nil {
			fmt.Println("Unmarshaling error: ", error)
			return
		}

		indentedJson, error := json.MarshalIndent(objmap, "", "\t")
		fmt.Println(string(indentedJson))
	}
}

func readInput() string {
	reader := bufio.NewReader(os.Stdin)
	token, error := reader.ReadString('\n')
	if error != nil {
		fmt.Println("Error reading from input")
	}
	return token
}

func addPadding(v string) string {
	if remainder := len(v) % 4; remainder > 0 {
		for j := 0; j < 4-remainder; j++ {
			v = v + "="
		}
	}
	return v
}
