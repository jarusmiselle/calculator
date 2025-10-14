package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	a()
	b()
}

func a() {
	num := 5
	bad := true
	blue := "red"

	out, _ := json.Marshal(num)
	fmt.Println(string(out))

	out, _ = json.Marshal(bad)
	fmt.Println(string(out))

	out, _ = json.Marshal(blue)
	fmt.Println(string(out))
}

func b() {
	num := "27"
	bad := "false"
	blue := "\"grey\""

	var numV int
	json.Unmarshal([]byte(num), &numV)
	fmt.Println(numV)

	var badV bool
	json.Unmarshal([]byte(bad), &badV)
	fmt.Println(badV)

	var blueV string
	json.Unmarshal([]byte(blue), &blueV)
	fmt.Println(blueV)
}
