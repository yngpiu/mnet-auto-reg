package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

const sessionFile = "accounts.txt"

func saveSession(accounts []Account) {
	f, err := os.OpenFile(sessionFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		printRed(fmt.Sprintf("  Failed to save accounts: %s", err.Error()))
		return
	}
	defer f.Close()

	now := time.Now()
	fmt.Fprintf(f, "--- %s ---\n", now.Format("2006.01.02 15:04:05"))
	for _, a := range accounts {
		fmt.Fprintf(f, "%s\n", a.Email)
	}

	printGray(fmt.Sprintf("  Total accounts saved: %d", getEmailCount()))
}

func getEmailCount() int {
	data, err := os.ReadFile(sessionFile)
	if err != nil {
		return 0
	}
	count := 0
	for _, line := range strings.Split(string(data), "\n") {
		line = strings.TrimSpace(line)
		if line != "" && !strings.HasPrefix(line, "---") {
			count++
		}
	}
	return count
}
