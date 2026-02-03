package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	err := walk("/workspaces/calculator/")
	if err != nil {
		log.Fatal(err)
	}
}

func walk(name string) error {
	name = strings.TrimRight(name, "/") + "/"

	files, err := os.ReadDir(name)
	if err != nil {
		return fmt.Errorf("walk: %w", err)
	}

	for _, file := range files {
		path := name + file.Name()

		if file.IsDir() {
			err := walk(path)
			if err != nil {
				return fmt.Errorf("walk: %w", err)
			}
		}
	}

	return nil
}
