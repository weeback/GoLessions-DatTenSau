package main

import (
	"fmt"
	"strings"
)

func stdReadStr(str string, def ...string) string {
	var val string
	print(str) // Print the string
	_, err := fmt.Scanf("%s\n", &val)
	if err != nil || val == "" {
		return strings.Join(def, "")
	}
	return val
}
