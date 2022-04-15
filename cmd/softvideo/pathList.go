package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

type pathList struct {
	paths []string
}

func newPathList() *pathList {
	return new(pathList)
}

func (p *pathList) addPath(path string) {
	if p.pathExists(path) {
		return
	}

	p.paths = append(p.paths, path)
}

func (p *pathList) pathExists(path string) bool {
	for k := range p.paths {
		if p.paths[k] == path {
			return true
		}
	}
	return false
}

func (p *pathList) save() {
	file, err := os.OpenFile("paths.txt", os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}

	datawriter := bufio.NewWriter(file)

	for _, data := range p.paths {
		_, _ = datawriter.WriteString(data + "\n")
	}

	datawriter.Flush()
	file.Close()
}

func (p *pathList) load() {
	for _, line := range linesInFile(`paths.txt`) {
		p.paths = append(p.paths, line)
	}
}

func linesInFile(fileName string) []string {
	f, err := os.Open(fileName)
	if err != nil {
		return []string{}
	}

	// Create new Scanner.
	scanner := bufio.NewScanner(f)
	var result []string
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
