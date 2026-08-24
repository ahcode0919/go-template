package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func replaceLine(path string, line int, newLine string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	lines := strings.Split(string(data), "\n")
	if line < 0 || line >= len(lines) {
		return fmt.Errorf("line %d out of range", line)
	}
	lines[line] = newLine
	return os.WriteFile(path, []byte(strings.Join(lines, "\n")), 0644)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter module name (github.com/user/modulename):")
	scanner.Scan()
	input := scanner.Text()
	
	err := replaceLine("./go.mod", 0, "module " + input)

	if err != nil {
		fmt.Fprintln(os.Stderr, "error modifying go.mod:", err)
		os.Exit(1)
	}

	err = os.Remove("./migrate.go")

	if err != nil {
		fmt.Fprintln(os.Stderr, "error removing migration file:", err)
		os.Exit(1)
	}
}