package resi

import (
	"errors"
	"net/http"
	"encoding/json"
	"bytes"
	"io"
)

type Token_Req_Password struct {
	Username string `json:"username"`
	Password string `json:"password"`
	GrantType string `json:"grant_type" default:"password"`
}
type Token_Resp_Error struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
type Token_Resp struct {
	AccessToken string `json:"access_token,omitempty"`
	ExpiresIn int `json:"expires_in,omitempty"`
	Errors []Token_Resp_Error `json:"errors,omitempty"`
}

func Authenticate(username string, password string) (string, error) {
	if username == "" {
		return "", errors.New("ERROR: Username is missing")
	}
	if password == "" {
		return "", errors.New("ERROR: Password is missing")
	}

	var body Token_Req_Password
	body.Username = username
	body.Password = password
	body.GrantType = "password"

	jsonValue, _ := json.Marshal(body)
	b := bytes.NewReader(jsonValue)

	req, err := http.NewRequest("POST", "https://central.resi.io/api/v3/auth/token", b)
	if err != nil {
		return "", errors.New("ERROR: failed to setup new HTTP request")
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", errors.New("ERROR: failed to connect to central.resi.io for authentiation")
	}
	if resp.Body != nil {
		defer resp.Body.Close()
	}

	respBody, readErr := io.ReadAll(resp.Body)
	if readErr != nil {
		return "", errors.New("ERROR: failed to read response from resi.io")
	}

	var respToken Token_Resp
  respTokenErr := json.Unmarshal(respBody, &respToken)
	if respTokenErr != nil {
		return "", errors.New("ERROR: failed to read JSON response from resi.io")
	}

	if respToken.Errors != nil {
		return "", errors.New("ERROR: resi.io authentication failure = "+respToken.Errors[0].Message)
	}
	return respToken.AccessToken, nil
}