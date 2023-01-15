package resi

import (
	"net/http"
	"errors"
	"io"
	"encoding/json"
)

type EncoderProfile struct {
	UUID string `json:"uuid,omitempty"`
	Name string `json:"name,omitempty"`
	CustomerId string `json:"customerId,omitempty"`
	SerialNumber string `json:"serialNumber,omitempty"`
	Status string `json:"status,omitempty"`
	OperationalState string `json:"operationalState,omitempty"`
}

func EncoderProfiles(token string, customerId string) ([]EncoderProfile, error) {
	httpRequest, httpRequestErr := http.NewRequest("GET", "https://central.resi.io/api/v3/customers/"+customerId+"/encoderprofiles", nil)
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

	var encoderProfiles []EncoderProfile
	encoderProfilesErr := json.Unmarshal(respBody, &encoderProfiles)
	if encoderProfilesErr != nil {
		return nil, errors.New("ERROR: failed to read JSON response from resi.io for encoder profiles")
	}
	return encoderProfiles, nil
}

