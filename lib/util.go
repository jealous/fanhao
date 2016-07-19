package fanhao

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func CurrentFolder() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	return dir
}

func Files(folder string) []string {
	files := make([]string, 0)
	fileInfoList, err := ioutil.ReadDir(folder)

	if err != nil {
		log.Fatal(err)
	} else {
		for _, fileInfo := range fileInfoList {
			if !fileInfo.IsDir() {
				files = append(files, fileInfo.Name())
			}
		}
	}
	return files
}

func Exists(filename string) bool {
	ret := false
	if _, err := os.Stat(filename); err == nil {
		ret = true
	}
	return ret
}

func Rename(oldName, newName string) bool {
	renamed := false
	if Exists(newName) {
		log.Printf("%v already exists, skip.", newName)
	} else {
		os.Rename(oldName, newName)
		renamed = true
	}
	return renamed
}
