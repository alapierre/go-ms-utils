package oauth

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

type TokenInfo struct {
	UserName         string   `json:"user_name,omitempty"`
	Active           bool     `json:"active,omitempty"`
	Exp              int      `json:"exp,omitempty"`
	ClientId         string   `json:"client_id,omitempty"`
	Scope            []string `json:"scope,omitempty"`
	Authorities      []string `json:"authorities,omitempty"`
	Error            string   `json:"error,omitempty"`
	ErrorDescription string   `json:"error_description,omitempty"`
	StatusCode       int
}

func CheckToken(url, token string) (TokenInfo, error) {

	url = trimTailingSlash(url)
	url = fmt.Sprintf("%s/check_token?token=%s", url, token)

	fmt.Println(url)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return TokenInfo{}, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return TokenInfo{}, err
	}

	log.Printf("response status %d", resp.StatusCode)

	var tokenInfo TokenInfo
	err = json.NewDecoder(resp.Body).Decode(&tokenInfo)
	if err != nil {
		return TokenInfo{}, err
	}

	tokenInfo.StatusCode = resp.StatusCode
	return tokenInfo, nil
}

func trimTailingSlash(url string) string {
	if strings.HasSuffix(url, "/") {
		return url[:len(url)-1]
	}
	return url
}
