package resi

import (
	"net/http"
	"errors"
	"io"
	"encoding/json"
)

type Encoder struct {
	UUID string `json:"uuid"`
	Name string `json:"name"`
	CustomerId string `json:"customerId"`
	SerialNumber string `json:"serialNumber"`
	Status string `json:"status"`
	OperationalState string `json:"operationalState"`
}

func Encoders(token string) ([]Encoder, error) {
	httpRequest, httpRequestErr := http.NewRequest("GET", "https://central.resi.io/api_v2.svc/encoders?wide=true", nil)
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

	var encoders []Encoder
	encodersErr := json.Unmarshal(respBody, &encoders)
	if encodersErr != nil {
		return nil, errors.New("ERROR: failed to read JSON response from resi.io for encoders")
	}
	return encoders, nil
}

//https://central.resi.io/api_v2.svc/encoders?wide=true
//Authorization: X-Bearer <token>
