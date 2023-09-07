package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func closeFile(file *os.File) {
	err := file.Close()
	if err != nil {
		log.Fatal(err)
	}
}

func ReadFile(file string) (lines []string) {
	f, err := os.Open(file)
	if err != nil {
		if strings.Contains(err.Error(), "no such file or directory") {
			// print warning and return without error; we might be adding the first do
			fmt.Printf("[file '%s' not found]\n", file)
			return lines
		}
		log.Fatal(err)
	}
	defer closeFile(f)

	input := bufio.NewScanner(f)
	for input.Scan() {
		lines = append(lines, input.Text())
	}
	return lines
}

func WriteFile(file string, lines []string) {
	f, err := os.Create(file)
	if err != nil {
		log.Fatal(err)
	}
	defer closeFile(f)
	w := bufio.NewWriter(f)
	for _, line := range lines {
		_, err := fmt.Fprintf(w, "%s\n", line)
		if err != nil {
			log.Fatal(err)
		}
	}
	err = w.Flush()
	if err != nil {
		log.Println(err)
		return
	}
}
