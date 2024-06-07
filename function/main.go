package main

import (
	"fmt"
	"strconv"
)

func getInfo() string {
	m := map[string]string{"foo": "1", "bar": "2"}

	if s, ok := m["baz"]; !ok {
		return "not found"
	} else {
		return strconv.Itoa(len(s))
	}
}

func main() {
	fmt.Println(getInfo())
}
