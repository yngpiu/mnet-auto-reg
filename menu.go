package main

import "fmt"

func showMenu(cfg Config) {
	fmt.Println()
	printBold("  Mnet Plus Account Auto Registration")
	fmt.Println()
	printGray(fmt.Sprintf("  Current password: %s", cfg.Password))

	modeLabel := "Auto"
	for _, m := range modes {
		if m.Key == cfg.Mode {
			modeLabel = m.Label
			break
		}
	}
	printGray(fmt.Sprintf("  Current mode: %s", modeLabel))
	fmt.Println()
	printCyan("  1. Create Account")
	printCyan("  2. Switch Mode")
	printCyan("  3. Change Password")
	printCyan("  4. Exit")
	fmt.Println()
}

func promptAccountCount() (int, bool) {
	input := askQuestion("Number of accounts to create")
	count := 0
	_, err := fmt.Sscanf(input, "%d", &count)
	if err != nil || count <= 0 {
		printRed("  Invalid number!")
		return 0, false
	}
	return count, true
}

func promptEmail() (string, bool) {
	email := askQuestion("Enter email address")
	if email == "" {
		printRed("  Invalid number!")
		return "", false
	}
	return email, true
}

func promptModeSelection(current string) string {
	fmt.Println()
	printBold("  Switch Mode")
	fmt.Println()
	for i, m := range modes {
		marker := ""
		if m.Key == current {
			marker = " (current)"
		}
		printCyan(fmt.Sprintf("  %d. %s%s", i+1, m.Label, marker))
	}
	printCyan(fmt.Sprintf("  %d. Exit", len(modes)+1))
	fmt.Println()

	input := askQuestion("Select")
	var idx int
	_, err := fmt.Sscanf(input, "%d", &idx)
	if err != nil || idx < 1 || idx > len(modes)+1 {
		printRed("  Invalid number!")
		return ""
	}
	if idx == len(modes)+1 {
		return ""
	}
	return modes[idx-1].Key
}
