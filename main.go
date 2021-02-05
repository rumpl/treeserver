package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if err := getDirectoryContents(scanner.Text()); err != nil {
			return err
		}
	}

	return nil
}
