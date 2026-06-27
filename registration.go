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
	printBlue("  Registering...")

	if err := mnetClient.InitSession(); err != nil {
		return nil, fmt.Errorf("session init failed: %w", err)
	}
	sleep(randomInt(2000, 4000))

	avail, err := mnetClient.CheckMailAvailability(email)
	if err != nil {
		return nil, fmt.Errorf("email check failed: %w", err)
	}
	if avail != 0 {
		printYellow("  Email not available for signup")
		return nil, nil
	}

	sleep(randomInt(1500, 3000))

	authToken, err := mnetClient.GetTokenForSignup(email)
	if err != nil {
		return nil, fmt.Errorf("auth token failed: %w", err)
	}

	sleep(randomInt(2000, 4000))

	_, err = mnetClient.Signup(email, password, authToken, gender, birthdate.Year)
	if err != nil {
		return nil, fmt.Errorf("signup failed: %w", err)
	}

	printBlue("  Verifying...")
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
					sleep(randomInt(2000, 5000))

					res, err := mnetClient.VerifyEmail(verifyToken)
					if err != nil {
						printRed(fmt.Sprintf("  Verification error: %s", err.Error()))
						return nil, fmt.Errorf("verify error: %w", err)
					}

					var accessToken string
					if d, ok := res["data"].(map[string]interface{}); ok {
						if t, _ := d["accessToken"].(string); t != "" {
							accessToken = t
						}
					}

					if accessToken != "" {
						status, err := mnetClient.CheckSignupStatus(email)
						if err != nil {
							printRed(fmt.Sprintf("  Status check failed: %s", err.Error()))
							return nil, nil
						}
						if status == 1 {
							printGreen("  Account verified!")
							return &Account{
								Email:     email,
								Password:  password,
								CreatedAt: time.Now().Format(time.RFC3339),
							}, nil
						}
						printYellow(fmt.Sprintf("  Signup status: %d (expected 1) — server not ready", status))
						return nil, nil
					}

					printRed("  Verification failed: no access token in response")
					printGray(fmt.Sprintf("  Server response: %v", truncate(fmt.Sprintf("%v", res), 300)))
					return nil, nil
				}
			}
		}
	}

	printRed("  Timeout: No verification email received")
	return nil, nil
}
