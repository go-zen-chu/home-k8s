package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

var (
	placeholderRegex = regexp.MustCompile(`<([\w-]+?)>`)
)

func convertToEnvVarStr(s string) string {
	upper := strings.ToUpper(s)
	ev := strings.ReplaceAll(upper, "-", "_")
	return ev
}

func replaceAVPPlaceholder(filePathStr string) error {
	var fi os.FileInfo
	var err error
	if fi, err = os.Stat(filePathStr); err != nil {
		return fmt.Errorf("could not find path %s: %w", filePathStr, err)
	}
	if fi.IsDir() {
		return fmt.Errorf("given path is dir: %s", filePathStr)
	}
	var fp *os.File
	if fp, err = os.Open(filePathStr); err != nil {
		return fmt.Errorf("open file: %w", err)
	}
	defer fp.Close()

	sc := bufio.NewScanner(fp)
	for sc.Scan() {
		line := sc.Text()
		fmt.Println(line)
		matches := placeholderRegex.FindAllStringSubmatch(line, -1)
		if len(matches) == 0 {
			continue
		}
		for _, submatches := range matches {
			if len(submatches) == 0 {
				continue
			}
			submatch := submatches[1]
			evs := convertToEnvVarStr(submatch)
			fmt.Println(evs)
			ev := os.Getenv(evs)
			fmt.Println(ev)
		}
	}
	return nil
}

func main() {
	if len(os.Args) == 1 {
		panic("usage: replace <any_file_path>")
	}
	if err := replaceAVPPlaceholder(os.Args[1]); err != nil {
		panic(fmt.Errorf("replace argo vault plugin placeholder failed: %w", err))
	}
}
