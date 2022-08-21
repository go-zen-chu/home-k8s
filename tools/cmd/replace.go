package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
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

func replaceAVPPlaceholder(filePathStr string) (string, error) {
	var fi os.FileInfo
	var err error
	if fi, err = os.Stat(filePathStr); err != nil {
		return "", fmt.Errorf("could not find path %s: %w", filePathStr, err)
	}
	if fi.IsDir() {
		return "", fmt.Errorf("given path is dir: %s", filePathStr)
	}
	var fp *os.File
	if fp, err = os.Open(filePathStr); err != nil {
		return "", fmt.Errorf("open file: %w", err)
	}
	defer fp.Close()

	sc := bufio.NewScanner(fp)
	var sb strings.Builder
	for sc.Scan() {
		line := sc.Text()
		matches := placeholderRegex.FindAllStringSubmatch(line, -1)
		if len(matches) == 0 {
			sb.WriteString(line + "\n")
			continue
		}
		replacedLine := line
		for _, submatches := range matches {
			if len(submatches) == 0 {
				return "", fmt.Errorf("could not find submatch: %v", matches)
			}
			submatch := submatches[1]
			evs := convertToEnvVarStr(submatch)
			ev := os.Getenv(evs)
			if ev == "" {
				return "", fmt.Errorf("failed to replace %s, could not find environment valiable %s", submatch, evs)
			}
			if _, err := strconv.ParseFloat(ev, 64); err == nil {
				// is numeric
				replacedLine = placeholderRegex.ReplaceAllString(replacedLine, ev)
			} else {
				// double quote string value
				replacedLine = placeholderRegex.ReplaceAllString(replacedLine, `"`+ev+`"`)
			}
		}
		sb.WriteString(replacedLine + "\n")
	}
	return sb.String(), nil
}

func main() {
	if len(os.Args) == 1 {
		panic("usage: go run ./tools/cmd/replace <any_file_path>")
	}
	if val, err := replaceAVPPlaceholder(os.Args[1]); err != nil {
		panic(fmt.Errorf("replace argo vault plugin placeholder failed: %w", err))
	} else {
		fmt.Println(val)
	}
}
