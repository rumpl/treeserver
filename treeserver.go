package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path"
	"syscall"
	"time"
)

const volume = "/volume"

type Info struct {
	Name  string    `json:"name"`
	Type  string    `json:"type"`
	Size  int64     `json:"size"`
	ATime time.Time `json:"atime"`
	MTime time.Time `json:"mtime"`
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
		stat_t := f.Sys().(*syscall.Stat_t)
		a.Contents = append(a.Contents, Info{
			Name:  f.Name(),
			Size:  f.Size(),
			ATime: timespecToTime(stat_t.Atim),
			MTime: timespecToTime(stat_t.Mtim),
			Type:  t,
		})
	}

	b, err := json.Marshal(a)
	if err != nil {
		return err
	}

	fmt.Println(string(b))

	return nil
}

func timespecToTime(ts syscall.Timespec) time.Time {
	return time.Unix(int64(ts.Sec), int64(ts.Nsec))
}
