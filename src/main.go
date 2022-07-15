package main

import (
	"encoding/base64"
	"fmt"
	"os"
	"strings"
)

func main() {
	var whatIsIt string

	secret := "aWFuZ25vVzpOQU06RU5JTDp0YTpzdTpuaW9K"

	decoded, err := base64.StdEncoding.DecodeString(secret)
	if err != nil {
		fmt.Println("[Failed to decode string]", err)
		os.Exit(1)
	}

	var sb strings.Builder

	for i := len(decoded) - 1; i >= 0; i-- {
		sb.WriteByte(decoded[i])
	}

	whatIsIt = sb.String()

	fmt.Println(whatIsIt)
}
