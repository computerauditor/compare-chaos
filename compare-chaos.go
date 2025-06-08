package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

// Helper to read lines from file into map
func readLines(filePath string) (map[string]struct{}, error) {
	lines := make(map[string]struct{})
	file, err := os.Open(filePath)
	if err != nil {
		return lines, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" {
			lines[line] = struct{}{}
		}
	}
	return lines, scanner.Err()
}

// Helper to write lines to file
func writeLines(lines []string, filePath string) error {
	if err := os.MkdirAll(filepath.Dir(filePath), 0755); err != nil {
		return err
	}
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	for _, line := range lines {
		_, err := file.WriteString(line + "\n")
		if err != nil {
			return err
		}
	}
	return nil
}

// Compare two files and output added/removed/unchanged lines
func compareFiles(oldPath, newPath, outputDir string, verbose, noUnchanged bool) {
	oldData, oldErr := readLines(oldPath)
	newData, newErr := readLines(newPath)

	programFile := filepath.Base(newPath)
	programName := strings.TrimSuffix(programFile, ".txt")
	resultDir := filepath.Join(outputDir, programName)

	oldSet := oldData
	if oldErr != nil {
		oldSet = make(map[string]struct{})
	}
	newSet := newData
	if newErr != nil {
		newSet = make(map[string]struct{})
	}

	var added, removed, unchanged []string

	for line := range newSet {
		if _, exists := oldSet[line]; exists {
			unchanged = append(unchanged, line)
		} else {
			added = append(added, line)
		}
	}

	for line := range oldSet {
		if _, exists := newSet[line]; !exists {
			removed = append(removed, line)
		}
	}

	sort.Strings(added)
	sort.Strings(removed)
	sort.Strings(unchanged)

	_ = writeLines(added, filepath.Join(resultDir, "added.txt"))
	_ = writeLines(removed, filepath.Join(resultDir, "removed.txt"))
	if !noUnchanged {
		_ = writeLines(unchanged, filepath.Join(resultDir, "unchanged.txt"))
	}

	if verbose {
		log.Printf("[INFO] %-25s → Added: %3d | Removed: %3d | Unchanged: %3d", programName, len(added), len(removed), len(unchanged))
	}
}

// Show user-friendly help
func printUsage() {
	fmt.Println(`Usage: compare-chaos [options]

Compare two chaos-output folders and identify added, removed, and unchanged URLs per program.

Flags:
  -n, --new, --today        Path to today’s chaos-output folder
  -p, --old, --yesterday    Path to yesterday’s chaos-output folder
  -o, --output              Output directory for comparison results (default: results)
  -v, --verbose             Enable verbose output
  --nu, --no-unchanged      Skip writing 'unchanged.txt' files

Example:
  go run compare-chaos.go \
    -n chaos-output-2025-06-08 \
    -p chaos-output-2025-06-07 \
    -o results \
    -v \
    --nu`)
}

func main() {
	newDir := flag.String("n", "", "")
	oldDir := flag.String("p", "", "")
	outputDir := flag.String("o", "results", "")
	verbose := flag.Bool("v", false, "")
	noUnchanged := flag.Bool("nu", false, "")

	flag.StringVar(newDir, "new", "", "")
	flag.StringVar(newDir, "today", "", "")
	flag.StringVar(oldDir, "old", "", "")
	flag.StringVar(oldDir, "yesterday", "", "")
	flag.StringVar(outputDir, "output", "results", "")
	flag.BoolVar(verbose, "verbose", false, "")
	flag.BoolVar(noUnchanged, "no-unchanged", false, "")

	flag.Usage = printUsage
	flag.Parse()

	if *newDir == "" || *oldDir == "" {
		fmt.Println("[ERROR] Both --new (today) and --old (yesterday) directories must be provided.\n")
		printUsage()
		os.Exit(1)
	}

	files, err := os.ReadDir(*newDir)
	if err != nil {
		log.Fatalf("[ERROR] Failed to read today's directory: %v\n", err)
	}

	log.Printf("[INFO] Starting comparison:\n  OLD: %s\n  NEW: %s\n", *oldDir, *newDir)

	for _, entry := range files {
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".txt") {
			continue
		}
		newPath := filepath.Join(*newDir, entry.Name())
		oldPath := filepath.Join(*oldDir, entry.Name())
		compareFiles(oldPath, newPath, *outputDir, *verbose, *noUnchanged)
	}

	log.Println("[INFO] ✅ Comparison completed. Results stored in:", *outputDir)
}
