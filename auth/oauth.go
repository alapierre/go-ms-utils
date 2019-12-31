package oauth

import (
	"encoding/base64"
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

type Token struct {
	AccessToken      string `json:"access_token"`
	TokenType        string `json:"token_type"`
	RefreshToken     string `json:"refresh_token"`
	ExpiresIn        int64  `json:"expires_in"`
	Error            string `json:"error,omitempty"`
	ErrorDescription string `json:"error_description,omitempty"`
	StatusCode       int
}

func CheckToken(url, token string) (TokenInfo, error) {

	url = trimTailingSlash(url)
	url = fmt.Sprintf("%s/check_token?token=%s", url, token)

	log.Println(url)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return TokenInfo{}, err
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Cache-Control", "no-cache")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return TokenInfo{}, err
	}
	defer resp.Body.Close()

	log.Printf("response status %d", resp.StatusCode)

	var tokenInfo TokenInfo
	err = json.NewDecoder(resp.Body).Decode(&tokenInfo)
	if err != nil {
		return TokenInfo{}, err
	}

	tokenInfo.StatusCode = resp.StatusCode
	return tokenInfo, nil
}

func GetToken(url, user, password, client, secret string) (Token, error) {

	url = trimTailingSlash(url)
	url = fmt.Sprintf("%s/token?grant_type=password&username=%s&password=%s", url, user, password)

	log.Println(url)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Token{}, err
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Cache-Control", "no-cache")

	authValue := fmt.Sprintf("Basic %s", base64.StdEncoding.EncodeToString([]byte(client+":"+secret)))
	req.Header.Add("Authorization", authValue)

	server := &http.Client{}
	resp, err := server.Do(req)
	if err != nil {
		return Token{}, err
	}
	defer resp.Body.Close()

	log.Printf("response status %d", resp.StatusCode)

	var token Token
	err = json.NewDecoder(resp.Body).Decode(&token)
	if err != nil {
		return Token{}, err
	}

	token.StatusCode = resp.StatusCode
	return token, nil
}

func trimTailingSlash(url string) string {
	if strings.HasSuffix(url, "/") {
		return url[:len(url)-1]
	}
	return url
}
