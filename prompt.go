package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func askQuestion(prompt string) string {
	fmt.Print(colorCyan + "  " + prompt + ": " + colorReset)
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text)
}
