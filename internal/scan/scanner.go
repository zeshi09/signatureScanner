package scan

import (
	"os"
	"path/filepath"
	"regexp"
	"sync"

	"github.com/zeshi09/signatureScanner/internal/rules"
)

type Finding struct {
	File  string
	Line  int
	Match string
}

func Run(root string) []Finding {
	var findings []Finding
	var mu sync.Mutex
	var wg sync.WaitGroup

	fileChan := make(chan string)

	go func() {
		_ = filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
			if err == nil && !info.IsDir() {
				fileChan <- path
			}
			return nil
		})
		close(fileChan)
	}()

	for i := 0; i < 8; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for file := range fileChan {
				localFindings := scanFile(file, rules.GetRegexes())
				mu.Lock()
				findings = append(findings, localFindings...)
				mu.Unlock()
			}
		}()
	}

	wg.Wait()
	return findings
}

func scanFile(path string, patterns []*regexp.Regexp) []Finding {
	var results []Finding
	content, err := os.ReadFile(path)
	if err != nil {
		return nil
	}

	for _, re := range patterns {
		matches := re.FindAllString(string(content), -1)
		for _, m := range matches {
			results = append(results, Finding{
				File:  path,
				Match: m,
			})
		}
	}
	return results
}
