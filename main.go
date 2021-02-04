package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
)

const volume = "/volume"

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

type Info struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type Contents struct {
	Contents []Info `json:"contents"`
}

func getDirectoryContents(dir string) error {
	wanted := path.Join(volume, dir)
	files, err := ioutil.ReadDir(wanted)
	if err != nil {
		return err
	}

	a := Contents{
		Contents: []Info{},
	}
	for _, f := range files {
		t := "file"
		if f.IsDir() {
			t = "directory"
		}
		a.Contents = append(a.Contents, Info{
			Name: f.Name(),
			Type: t,
		})
	}

	b, err := json.Marshal(a)
	if err != nil {
		return err
	}

	fmt.Println(string(b))

	return nil
}
