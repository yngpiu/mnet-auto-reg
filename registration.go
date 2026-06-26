package main

import (
	"fmt"
	"regexp"
	"strings"
	"time"
)

type Account struct {
	Email     string `json:"email"`
	Password  string `json:"password"`
	CreatedAt string `json:"createdAt"`
}

func extractVerificationToken(body string) string {
	re := regexp.MustCompile(`https://id\.mnetplus\.world/signup-certify\?token=([^"'\s)]+)`)
	matches := re.FindStringSubmatch(body)
	if len(matches) > 1 {
		return matches[1]
	}
	return ""
}

func createAccountApi(email, password string, mailClient *TempMailClient, mnetClient *MnetPlusClient, birthdate Birthdate, gender string) (*Account, error) {
	printBlue("  Registering on Mnet...")

	if err := mnetClient.InitSession(); err != nil {
		return nil, fmt.Errorf("session init: %w", err)
	}
	sleep(randomInt(2000, 4000))

	avail, err := mnetClient.CheckMailAvailability(email)
	if err != nil {
		return nil, err
	}
	if avail != 0 {
		printYellow("  Email not available for signup")
		return nil, nil
	}

	sleep(randomInt(1500, 3000))

	authToken, err := mnetClient.GetTokenForSignup(email)
	if err != nil {
		return nil, err
	}
	printCyan(fmt.Sprintf("  Auth token received (%d chars)", len(authToken)))

	sleep(randomInt(2000, 4000))

	msg, err := mnetClient.Signup(email, password, authToken, gender, birthdate.Year)
	if err != nil {
		return nil, err
	}
	printGreen("  Registration successful!")
	printGreen(fmt.Sprintf("  Signup response: %s", msg))

	printBlue("  Waiting for verification email from Mnet...")
	startTime := time.Now()
	timeout := 90 * time.Second

	for time.Since(startTime) < timeout {
		time.Sleep(5 * time.Second)

		messages, err := mailClient.CheckMessages(email)
		if err != nil {
			continue
		}

		for _, msg := range messages {
			if strings.Contains(msg.Subject, "Mnet") {
				body := msg.BodyText + msg.BodyHTML
				verifyToken := extractVerificationToken(body)
				if verifyToken != "" {
					printGreen(fmt.Sprintf("  Verification token found: %s...", truncate(verifyToken, 20)))

					sleep(randomInt(2000, 5000))

					printBlue("  Verifying account...")
					res, err := mnetClient.VerifyEmail(verifyToken)
					if err != nil {
						printRed(fmt.Sprintf("  Verification error: %s", err.Error()))
						return nil, fmt.Errorf("verify error: %w", err)
					}

					printGray(fmt.Sprintf("  Verify response: %v", truncate(fmt.Sprintf("%v", res), 200)))

					var accessToken string
					if d, ok := res["data"].(map[string]interface{}); ok {
						if t, _ := d["accessToken"].(string); t != "" {
							accessToken = t
						}
					}

					if accessToken != "" {
						printGreen("  Account verified!")
						status, err := mnetClient.CheckSignupStatus(email)
						if err != nil {
							printYellow(fmt.Sprintf("  Status check error: %s", err.Error()))
						}
						if status == 1 {
							printGreen("  Signed up successfully via API!")
							return &Account{
								Email:     email,
								Password:  password,
								CreatedAt: time.Now().Format(time.RFC3339),
							}, nil
						}
						printYellow(fmt.Sprintf("  Signup status: %d (expected 1)", status))
					} else {
						printYellow("  Verify response missing accessToken")
					}

					return nil, nil
				}
			}
		}
	}

	printRed("  Timeout: No verification email received")
	return nil, nil
}
