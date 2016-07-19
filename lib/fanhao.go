package fanhao

import (
	"fmt"
	"log"
	"path/filepath"
	"regexp"
	"strings"
	"sync"
)

var (
	fh         *regexp.Regexp
	threadSize int = 10
)

func NormalizeAll(folder string) {

	wg := new(sync.WaitGroup)

	fileList := Files(folder)

	queue := make(chan string)

	for i := 0; i < threadSize; i++ {
		go renameFileWorker(queue, wg, folder)
	}

	for _, file := range fileList {
		wg.Add(1)
		queue <- file
	}

	close(queue)
	wg.Wait()
	fmt.Printf("Processed %d files.\n", len(fileList))
}

func renameFileWorker(queue chan string, wg *sync.WaitGroup, folder string) {
	for filename := range queue {
		renameFile(filename, wg, folder)
	}
}

func renameFile(name string, wg *sync.WaitGroup, folder string) {
	defer wg.Done()
	normalized := Normalize(name)

	oldPath := filepath.Join(folder, name)
	newPath := filepath.Join(folder, normalized)

	renamed := false

	if name == normalized {
		// skip
	} else if strings.ToUpper(name) == strings.ToUpper(normalized) {
		// you cannot rename directly in windows
		// because windows treat them as the same file
		tempPath := filepath.Join(folder, fmt.Sprintf("__%v", name))
		renamed = Rename(oldPath, tempPath)
		renamed = Rename(tempPath, newPath)
	} else {
		renamed = Rename(oldPath, newPath)
	}

	if renamed {
		log.Printf("%-10v -> %v", name, normalized)
	}
}

func getExp() *regexp.Regexp {
	if fh == nil {
		var err error
		fh, err = regexp.Compile(`(^[_A-Z0-9]+?)[-_]?(\d+)[-_]?([A-C]?)R?P?L?[D-Z,]?\s*\.(\w+$)`)
		if err != nil {
			log.Fatal(err)
		}
	}
	return fh
}

func preProcess(old string) string {
	ret := strings.ToUpper(old)
	ret = strings.TrimPrefix(ret, "HD-")
	squareRight := strings.LastIndex(ret, "]")
	squareLeft := strings.Index(ret, "[")
	if squareRight > squareLeft {
		ret = ret[0:squareLeft] + ret[squareRight+1:len(ret)]
	}
	return ret
}

func Normalize(old string) string {
	ret := removeDuplicatedSuffix(old)
	exp := getExp()
	matched := exp.FindAllStringSubmatch(preProcess(ret), -1)
	length := len(matched)
	if length > 0 {
		groups := matched[0]
		product := strings.ToUpper(groups[1])
		number := strings.TrimLeft(groups[2], " 0")
		series := strings.ToUpper(groups[3])
		ext := strings.ToLower(groups[4])
		if len(groups[3]) == 0 {
			ret = fmt.Sprintf("%v-%03v.%v", product, number, ext)
		} else {
			ret = fmt.Sprintf("%v-%03v_%v.%v", product, number, series, ext)
		}
	} else {
		log.Printf("%v doesn't match to a FH.", ret)
	}

	return ret
}

func removeDuplicatedSuffix(filename string) string {
	currentName := filename
	ext := filepath.Ext(currentName)
	for filepath.Ext(strings.TrimSuffix(currentName, ext)) == ext {
		currentName = strings.TrimSuffix(currentName, ext)
	}
	return currentName
}
