package fanhao

import (
	"fmt"
	"log"
	"path/filepath"
	"regexp"
	"strings"
	"sync"
	"unicode"
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
		fh, err = regexp.Compile(`(^[_A-Z0-9]+?)[-_]?(\d+)\s*[-_~]?([A-C1-9]?)R?P?L?@?[A-Z\d.,+]*\s*\.(\w+$)`)
		if err != nil {
			log.Fatal(err)
		}
	}
	return fh
}

func preProcess(old string) string {
	ret := strings.ToUpper(old)
	ret = strings.TrimPrefix(ret, "HD")
	ret = strings.Trim(ret, `_-`)
	ret = removePair(ret, "[]")
	ret = removePair(ret, "()")
	ret = removeInvalidFields(ret, "_")
	return ret
}

func removeInvalidFields(input string, sep string) string {
	ext := filepath.Ext(input)
	s := strings.TrimSuffix(input, ext)
	var invalidSet = []string{"@", "."}
	o := strings.Split(s, sep)
	if len(o) == 1 {
		return input
	}
	n := make([]string, 0)
	for _, f := range o {
		isInvalid := false
		for _, c := range invalidSet {
			if strings.Index(f, c) != -1 {
				isInvalid = true
				break
			}
		}
		if !isInvalid {
			n = append(n, f)
		}
	}
	return strings.Join(n, sep) + ext
}

func removePair(s, pair string) string {
	if len(pair) != 2 {
		panic("pair should have 2 characters")
	}
	left := string(pair[0])
	right := string(pair[1])
	for {
		iRight := strings.LastIndex(s, right)
		if iRight == -1 {
			break
		}
		iLeft := strings.LastIndex(s, left)
		if iLeft == -1 {
			break
		}
		s = s[0:iLeft] + s[iRight+1:]
	}
	return s
}

func Normalize(old string) string {
	ret := removeDuplicatedSuffix(old)
	exp := getExp()
	preProcessed := preProcess(ret)
	if !first2LetterHasNumber(preProcessed) {
		matched := exp.FindAllStringSubmatch(preProcessed, -1)
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
	} else {
		ret = removeExt(preProcessed) + strings.ToLower(filepath.Ext(preProcessed))
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

func removeExt(fn string) string {
	return strings.TrimSuffix(fn, filepath.Ext(fn))
}

func first2LetterHasNumber(s string) bool {
	name := removeExt(s)
	if len(name) <= 2 {
		return false
	}
	s = name[:2]
	for _, r := range s {
		if unicode.IsNumber(r) {
			return true
		}
	}
	return false
}
