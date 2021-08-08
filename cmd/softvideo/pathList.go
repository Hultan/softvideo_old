package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

type PathList struct {
	Paths []string
}

func NewPathList() *PathList {
	return new(PathList)
}

func (p *PathList) AddPath(path string) {
	if p.pathExists(path) {
		return
	}

	p.Paths = append(p.Paths, path)
}

func (p *PathList) pathExists(path string) bool {
	for k := range p.Paths {
		if p.Paths[k] == path {
			return true
		}
	}
	return false
}

func (p *PathList) save() {
	file, err := os.OpenFile("paths.txt", os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}

	datawriter := bufio.NewWriter(file)

	for _, data := range p.Paths {
		_, _ = datawriter.WriteString(data + "\n")
	}

	datawriter.Flush()
	file.Close()
}

func (p *PathList) load() {
	for _, line := range LinesInFile(`paths.txt`) {
		p.Paths = append(p.Paths, line)
	}
}

func LinesInFile(fileName string) []string {
	f, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	// Create new Scanner.
	scanner := bufio.NewScanner(f)
	result := []string{}
	// Use Scan.
	for scanner.Scan() {
		line := scanner.Text()
		// Append line to result.
		if strings.Trim(line, " ") != "" {
			result = append(result, line)
		}
	}
	return result
}