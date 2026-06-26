package main

import (
	"fmt"
	"time"
)

func handleCreateAccounts(cfg Config) {
	count, ok := promptAccountCount()
	if !ok {
		return
	}

	startTime := time.Now()
	var sessionAccounts []Account

	onCreated := func(a Account) {
		sessionAccounts = append(sessionAccounts, a)
	}

	accounts := runAccountCreation(cfg.Mode, count, cfg.Password, promptEmail, onCreated)

	if len(sessionAccounts) > 0 {
		saveSession(sessionAccounts)
	}

	elapsed := int(time.Since(startTime).Seconds())
	mins := elapsed / 60
	secs := elapsed % 60
	timeStr := fmt.Sprintf("%ds", secs)
	if mins > 0 {
		timeStr = fmt.Sprintf("%dm %ds", mins, secs)
	}

	fmt.Println()
	printBold(fmt.Sprintf("  Done! %d/%d accounts created successfully", len(accounts), count))
	printGray(fmt.Sprintf("  Total time: %s", timeStr))
	printGray(fmt.Sprintf("  Total accounts saved: %d", getEmailCount()))
	fmt.Println()
}

func handleChangePassword(cfg Config) Config {
	printGray(fmt.Sprintf("  Current password: %s", cfg.Password))
	printBlue("  Requirements: 8-20 chars, 1+ uppercase, 1+ lowercase, 1+ number, 1+ special char")
	fmt.Println()

	newPwd := askQuestion("Enter new password")

	valid, errors := validatePassword(newPwd)
	if !valid {
		printRed("  Password does not meet requirements")
		printRed(fmt.Sprintf("  Missing: %s", joinStrings(errors, ", ")))
		return cfg
	}

	updateConfig(map[string]string{"password": newPwd})
	printGreen("  Password changed successfully")
	return loadConfig()
}

func handleToggleMode(cfg Config) Config {
	selected := promptModeSelection(cfg.Mode)
	if selected == "" {
		return cfg
	}

	updateConfig(map[string]string{"mode": selected})
	modeLabel := "Auto"
	for _, m := range modes {
		if m.Key == selected {
			modeLabel = m.Label
			break
		}
	}
	printGreen(fmt.Sprintf("  Mode changed to %s", modeLabel))
	return loadConfig()
}

func joinStrings(strs []string, sep string) string {
	result := ""
	for i, s := range strs {
		if i > 0 {
			result += sep
		}
		result += s
	}
	return result
}
