package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type TempMailClient struct {
	baseURL    string
	domains    []string
	httpClient *http.Client
}

type TempMailResult struct {
	Email string `json:"email"`
	Token string `json:"token"`
}

type TempMailMessage struct {
	Subject  string `json:"subject"`
	BodyText string `json:"body_text"`
	BodyHTML string `json:"body_html"`
}

func NewTempMailClient() *TempMailClient {
	return &TempMailClient{
		baseURL:    "https://api.internal.temp-mail.io/api/v3",
		httpClient: &http.Client{},
	}
}

func (c *TempMailClient) headers() http.Header {
	h := http.Header{}
	h.Set("Accept", "*/*")
	h.Set("Origin", "https://temp-mail.io")
	h.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36")
	h.Set("Referer", "https://temp-mail.io/")
	h.Set("Accept-Language", "en-US,en;q=0.9")
	return h
}

func (c *TempMailClient) fetchDomains() ([]string, error) {
	if len(c.domains) > 0 {
		return c.domains, nil
	}

	req, _ := http.NewRequest("GET", c.baseURL+"/domains", nil)
	req.Header = c.headers()

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	var result struct {
		Domains []string `json:"domains"`
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	c.domains = result.Domains
	return c.domains, nil
}

func (c *TempMailClient) GetRandomDomain(exclude []string) (string, error) {
	domains, err := c.fetchDomains()
	if err != nil {
		return "", err
	}
	for _, d := range domains {
		blocked := false
		for _, ex := range exclude {
			if d == ex {
				blocked = true
				break
			}
		}
		if !blocked {
			return d, nil
		}
	}
	if len(domains) > 0 {
		return domains[0], nil
	}
	return "", fmt.Errorf("no domains available")
}

func (c *TempMailClient) CreateNewMail(maxAttempts int, specificDomain string) (*TempMailResult, error) {
	for i := 0; i < maxAttempts; i++ {
		body := map[string]interface{}{
			"min_name_length": 10,
			"max_name_length": 10,
		}
		if specificDomain != "" {
			body["domain"] = specificDomain
		}

		jsonBody, _ := json.Marshal(body)
		req, _ := http.NewRequest("POST", c.baseURL+"/email/new", bytes.NewReader(jsonBody))
		req.Header = c.headers()
		req.Header.Set("Content-Type", "application/json")

		resp, err := c.httpClient.Do(req)
		if err != nil {
			continue
		}

		respBody, _ := io.ReadAll(resp.Body)
		resp.Body.Close()

		var result TempMailResult
		if err := json.Unmarshal(respBody, &result); err != nil {
			continue
		}

		if result.Email != "" && !strings.Contains(result.Email, "gmeenramy") {
			return &result, nil
		}
	}
	return nil, fmt.Errorf("failed to create email after %d attempts", maxAttempts)
}

func (c *TempMailClient) CheckMessages(email string) ([]TempMailMessage, error) {
	req, _ := http.NewRequest("GET", c.baseURL+"/email/"+email+"/messages", nil)
	req.Header = c.headers()

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	var messages []TempMailMessage
	if err := json.Unmarshal(body, &messages); err != nil {
		return nil, err
	}
	return messages, nil
}
