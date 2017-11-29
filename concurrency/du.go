package main

import (
	"flag"
	"fmt"
	"os"
)

func walk(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	files, err := f.Readdir(0)
	if err != nil {
		return err
	}
	for _, file := range files {
		if file.IsDir() {
			walk(path + file.Name())
			continue
		}
		fmt.Printf("name: %s, size: %d\n", file.Name(), file.Size())
	}
	return nil
}

func main() {
	flag.Parse()
	var dirs []string
	if len(flag.Args()) == 0 {
		dirs = []string{"."}
	}
	for _, dir := range dirs {
		walk(dir)
	}
}
