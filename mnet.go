package main

import (
	"bytes"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type MnetPlusClient struct {
	baseURL       string
	cookies       string
	sessionID     string
	transactionID string
	httpClient    *http.Client
}

func NewMnetPlusClient() *MnetPlusClient {
	return &MnetPlusClient{
		baseURL:       "https://www.mnetplus.world",
		sessionID:     uuid(),
		transactionID: strings.ToLower(uuid()),
		httpClient:    &http.Client{},
	}
}

func uuid() string {
	b := make([]byte, 16)
	rand.Read(b)
	return fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
}

func (c *MnetPlusClient) baseHeaders() http.Header {
	h := http.Header{}
	h.Set("Host", "www.mnetplus.world")
	h.Set("Accept", "application/json, text/plain, */*")
	h.Set("Accept-Language", "en-GB,en-US;q=0.9,en;q=0.8")
	h.Set("User-Agent", "Mnet%20Plus/2080 CFNetwork/1404.0.5 Darwin/22.3.0")
	h.Set("x-transaction-id", c.transactionID)
	h.Set("x-user-agent", "en:US:3.48.0.3471:iOS:16.6.1:iPhone 14:UTC:00000000-0000-0000-0000-000000000000")
	h.Set("Origin", "https://id.mnetplus.world")
	h.Set("Referer", "https://id.mnetplus.world/")
	h.Set("Sec-Fetch-Dest", "empty")
	h.Set("Sec-Fetch-Mode", "cors")
	h.Set("Sec-Fetch-Site", "same-site")
	if c.cookies != "" {
		h.Set("Cookie", c.cookies)
	}
	return h
}

func (c *MnetPlusClient) InitSession() error {
	req, _ := http.NewRequest("GET", "https://id.mnetplus.world/", nil)
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
	req.Header.Set("Accept-Encoding", "gzip, deflate")
	req.Header.Set("Accept-Language", "en-GB,en-US;q=0.9,en;q=0.8")
	req.Header.Set("User-Agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 16_6_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/16.6 Mobile/15E148 Safari/604.1")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	io.Copy(io.Discard, resp.Body)

	cookies := parseCookies(resp.Header["Set-Cookie"])
	if cookies != "" {
		c.cookies = cookies
	}
	return nil
}

func parseCookies(setCookie []string) string {
	var parts []string
	for _, raw := range setCookie {
		if idx := strings.Index(raw, ";"); idx >= 0 {
			if eq := strings.Index(raw[:idx], "="); eq >= 0 {
				parts = append(parts, raw[:idx])
			}
		}
	}
	return strings.Join(parts, "; ")
}

func (c *MnetPlusClient) request(method, path string, body interface{}) (map[string]interface{}, error) {
	var reqBody io.Reader
	if body != nil {
		b, _ := json.Marshal(body)
		reqBody = bytes.NewReader(b)
	}

	req, _ := http.NewRequest(method, c.baseURL+path, reqBody)
	req.Header = c.baseHeaders()
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	newCookies := parseCookies(resp.Header["Set-Cookie"])
	if newCookies != "" {
		if c.cookies != "" {
			c.cookies = c.cookies + "; " + newCookies
		} else {
			c.cookies = newCookies
		}
	}

	respBody, _ := io.ReadAll(resp.Body)
	var result map[string]interface{}
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("parse error: %s", string(respBody[:min(len(respBody), 200)]))
	}
	return result, nil
}

func (c *MnetPlusClient) VerifyEmail(token string) (map[string]interface{}, error) {
	return c.request("POST", "/api/account-service/v3/user/signup", map[string]string{
		"authToken": token,
	})
}

func (c *MnetPlusClient) CheckMailAvailability(email string) (int, error) {
	res, err := c.request("GET", "/api/account-service/v1/user/signup/email/status?email="+email, nil)
	if err != nil {
		return -1, err
	}

	msg, _ := res["message"].(string)
	if strings.Contains(msg, "abuse.block") || strings.Contains(msg, "block") {
		return -1, fmt.Errorf("msg:popup.signup.abuse.block")
	}

	status := getInt(res, "status")
	return status, nil
}

func (c *MnetPlusClient) GetTokenForSignup(email string) (string, error) {
	res, err := c.request("POST", "/api/account-service/v1/user/email/authToken", map[string]string{
		"email":    email,
		"purpose":  "signup",
		"locale":   "en",
		"deviceName": "Apple iPhone 16.6.1",
	})
	if err != nil {
		return "", err
	}

	token, _ := res["token"].(string)
	if token == "" {
		if data, ok := res["data"].(map[string]interface{}); ok {
			token, _ = data["token"].(string)
		}
	}
	if token == "" {
		return "", fmt.Errorf("no token in response: %v", truncate(fmt.Sprintf("%v", res), 200))
	}
	return token, nil
}

func (c *MnetPlusClient) Signup(email, password, authToken, gender string, birthYear int) (string, error) {
	res, err := c.request("POST", "/api/account-service/v1/user/signup/save-tmp", map[string]interface{}{
		"email":         email,
		"password":      password,
		"gender":        gender,
		"birthDate":     fmt.Sprintf("%d", birthYear),
		"optionalTerms": []interface{}{},
		"authToken":     authToken,
	})
	if err != nil {
		return "", err
	}

	msg, _ := res["message"].(string)
	if msg == "" {
		if data, ok := res["data"].(map[string]interface{}); ok {
			msg, _ = data["message"].(string)
		}
	}
	if msg == "" {
		msg = fmt.Sprintf("%v", res)
	}
	return msg, nil
}

func (c *MnetPlusClient) CheckSignupStatus(email string) (int, error) {
	res, err := c.request("GET", "/api/account-service/v1/user/signup/email/status?email="+email, nil)
	if err != nil {
		return -1, err
	}
	return getInt(res, "status"), nil
}

func getInt(m map[string]interface{}, key string) int {
	switch v := m[key].(type) {
	case float64:
		return int(v)
	case int:
		return v
	default:
		if data, ok := m["data"].(map[string]interface{}); ok {
			return getInt(data, key)
		}
		return -1
	}
}

func truncate(s string, n int) string {
	if len(s) > n {
		return s[:n]
	}
	return s
}
