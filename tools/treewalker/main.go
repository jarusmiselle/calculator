package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	walk("/workspaces/calculator/")
}

func walk(name string) {
	name = strings.TrimRight(name, "/") + "/"

	files, err := os.ReadDir(name)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		path := name + file.Name()

		if file.IsDir() {
			fmt.Println(path + "/")
			walk(path)
		} else {
			fmt.Println(path)
		}
	}
}
