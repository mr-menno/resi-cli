package resi

import (
	"net/http"
	"errors"
	"io"
	"encoding/json"
)

type Me struct {
	CustomerId string `json:"customerId"`
	CustomerName string `json:"customerName"`
	UserName string `json:"userName"`
	UserId string `json:"userId"`
}

func Whoami(token string) (*Me, error) {
	httpRequest, httpRequestErr := http.NewRequest("GET", "https://central.resi.io/api_v2.svc/users/me", nil)
	if httpRequestErr != nil {
		return nil, errors.New("ERROR: failed to setup new HTTP request")
	}
	httpRequest.Header.Set("Authorization","X-Bearer "+token)
	httpRequest.Header.Set("Accept", "application/json")
	
	client := &http.Client{}
	response, err := client.Do(httpRequest)
	if err != nil {
		return nil, errors.New("ERROR: failed to connect to central.resi.io for encoders")
	}
	if response.Body != nil {
		defer response.Body.Close()
	}

	respBody, readErr := io.ReadAll(response.Body)
	if readErr != nil {
		return nil, errors.New("ERROR: failed to read response from resi.io for encoders")
	}

	var me Me
	meErr := json.Unmarshal(respBody, &me)
	if meErr != nil {
		return nil, errors.New("ERROR: failed to read JSON response from resi.io for /api_v2.svc/users/me")
	}
	return &me, nil
}

//https://central.resi.io/api_v2.svc/encoders?wide=true
//Authorization: X-Bearer <token>
