package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage: %v <filename>\n", os.Args[0])
		os.Exit(1)
	}

	err := run(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	} else {
		fmt.Fprintln(os.Stderr, "succeeded")
	}
}

func run(fileName string) error {
	f, err := os.Open(fileName)
	if err != nil {
		return fmt.Errorf("could not open %v : %v", fileName, err)
	}
	defer f.Close()
	return processFile(f, os.Stdout)
}

func processFile(r io.Reader, w io.Writer) error {
	sc := bufio.NewScanner(r)
	prevTabCnt := 0
	isFirst := true
	for sc.Scan() {
		line := sc.Text()
		tabCnt := countPrefixTabs(line)
		text := removePrefixNumbers(removePrefixTabs(line))

		if tabCnt > prevTabCnt {
			fmt.Fprintf(w, "\t%v", text)
		} else if isFirst {
			fmt.Fprintf(w, "%v%v", tabs(tabCnt), text)
			isFirst = false
		} else {
			fmt.Fprintf(w, "\n%v%v", tabs(tabCnt), text)
		}

		prevTabCnt = tabCnt
	}
	fmt.Fprintf(w, "\n")
	return sc.Err()
}

func tabs(cnt int) string {
	tabs := ""
	for i := 0; i < cnt; i++ {
		tabs = tabs + "\t"
	}
	return tabs
}

func countPrefixTabs(line string) int {
	tabs := 0
	for _, ch := range []rune(line) {
		if ch != '\t' {
			return tabs
		} else {
			tabs += 1
		}
	}
	return tabs
}

func removePrefixTabs(text string) string {
	trimmed := text
	for strings.HasPrefix(trimmed, "\t") {
		trimmed = strings.TrimPrefix(trimmed, "\t")
	}
	return trimmed
}

func removePrefixNumbers(text string) string {
	parts := strings.Split(text, " ")
	if len(parts) == 1 {
		return text
	}
	prefix := parts[0]
	content := strings.Join(parts[1:], " ")

	re := regexp.MustCompile("^[\\d\\.]+$")
	if re.Match([]byte(prefix)) {
		return content
	} else {
		return text
	}
}
