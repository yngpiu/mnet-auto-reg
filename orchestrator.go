package main

import (
	"fmt"
	"time"
)

func tryApiAuto(index, total int, password string, blockedDomains *[]string, onCreated func(Account)) *Account {
	maxRetries := 3

	for retry := 1; retry <= maxRetries; retry++ {
		mailClient := NewTempMailClient()
		mnetClient := NewMnetPlusClient()

		var domain string
		if retry > 1 {
			d, err := mailClient.GetRandomDomain(*blockedDomains)
			if err == nil && d != "" {
				domain = d
				printCyan(fmt.Sprintf("  Trying domain: %s", domain))
			}
		}

		emailResult, err := mailClient.CreateNewMail(5, domain)
		if err != nil || emailResult == nil {
			break
		}

		email := emailResult.Email
		if retry > 1 {
			printCyan(fmt.Sprintf("  Retrying with new email: %s", email))
		} else {
			printGreen(fmt.Sprintf("  Email: %s", email))
		}

		birthdate := randomBirthdate()
		gender := randomGender()

		account, err := createAccountApi(email, password, mailClient, mnetClient, birthdate, gender)
		if err != nil {
			if retry < maxRetries {
				printYellow(fmt.Sprintf("  Account %d/%d API error, retrying (%d/%d): %s", index, total, retry, maxRetries, err.Error()))
				continue
			}
			printRed(fmt.Sprintf("  Registration error: %s", err.Error()))
			break
		}

		if account != nil {
			createdAt := time.Now().Format(time.RFC3339)
			saved := Account{Email: email, Password: password, CreatedAt: createdAt}
			printGreen(fmt.Sprintf("  Account %d/%d completed: %s", index, total, email))
			printGray(fmt.Sprintf("     Created at: %s", createdAt))
			onCreated(saved)
			return &saved
		}

		parts := splitEmail(email)
		if len(parts) > 1 {
			blocked := false
			for _, b := range *blockedDomains {
				if b == parts[1] {
					blocked = true
					break
				}
			}
			if !blocked {
				*blockedDomains = append(*blockedDomains, parts[1])
			}
		}

		if retry < maxRetries {
			printYellow(fmt.Sprintf("  Account %d/%d failed, retrying (%d/%d)...", index, total, retry, maxRetries))
		}
	}

	return nil
}

func runAuto(count int, password string, onCreated func(Account)) []Account {
	blockedDomains := []string{}
	var accounts []Account

	for i := 1; i <= count; i++ {
		printCyan(fmt.Sprintf("\n  --- Account %d/%d ---", i, count))

		saved := tryApiAuto(i, count, password, &blockedDomains, onCreated)

		if saved != nil {
			accounts = append(accounts, *saved)
		} else {
			printRed(fmt.Sprintf("  Account %d/%d failed", i, count))
		}

		if i < count {
			d := delay("betweenAccounts")
			printGray(fmt.Sprintf("  Wait %ds...", d/1000))
			sleep(d)
		}
	}

	return accounts
}

func runManual(count int, password string, askEmail func() (string, bool), onCreated func(Account)) []Account {
	var accounts []Account

	for i := 1; i <= count; i++ {
		email, ok := askEmail()
		if !ok {
			continue
		}

		printCyan(fmt.Sprintf("\n  --- Account %d/%d ---", i, count))
		printBlue(fmt.Sprintf("  Manual email mode: %s", email))

		mailClient := NewTempMailClient()
		mnetClient := NewMnetPlusClient()
		birthdate := randomBirthdate()
		gender := randomGender()

		account, err := createAccountApi(email, password, mailClient, mnetClient, birthdate, gender)
		if err != nil {
			printRed(fmt.Sprintf("  Registration error: %s", err.Error()))
			printRed(fmt.Sprintf("  Account %d/%d failed", i, count))
			if i < count {
				d := delay("betweenAccounts")
				printGray(fmt.Sprintf("  Wait %ds...", d/1000))
				sleep(d)
			}
			continue
		}

		if account != nil {
			createdAt := time.Now().Format(time.RFC3339)
			saved := Account{Email: email, Password: password, CreatedAt: createdAt}
			printGreen(fmt.Sprintf("  Account %d/%d completed: %s", i, count, email))
			printGray(fmt.Sprintf("     Created at: %s", createdAt))
			accounts = append(accounts, saved)
			onCreated(saved)
		} else {
			printRed(fmt.Sprintf("  Account %d/%d failed", i, count))
		}

		if i < count {
			d := delay("betweenAccounts")
			printGray(fmt.Sprintf("  Wait %ds...", d/1000))
			sleep(d)
		}
	}

	return accounts
}

func runAccountCreation(mode string, count int, password string, askEmail func() (string, bool), onCreated func(Account)) []Account {
	if isManualMode(mode) {
		return runManual(count, password, askEmail, onCreated)
	}
	return runAuto(count, password, onCreated)
}

func splitEmail(email string) []string {
	for i := 0; i < len(email); i++ {
		if email[i] == '@' {
			return []string{email[:i], email[i+1:]}
		}
	}
	return []string{email}
}
