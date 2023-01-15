package resi

import (
	"net/http"
	"errors"
	"io"
	"encoding/json"
	"bytes"
	"fmt"
)

type StreamProfile struct {
	UUID string `json:"uuid"`
}

type Encoder struct {
	UUID string `json:"uuid,omitempty"`
	Name string `json:"name,omitempty"`
	CustomerId string `json:"customerId,omitempty"`
	SerialNumber string `json:"serialNumber,omitempty"`
	Status string `json:"status,omitempty"`
	OperationalState string `json:"operationalState,omitempty"`
	StreamProfile StreamProfile `json:"streamProfile,omitempty"`
	EncoderProfile EncoderProfile `json:"encoderProfile,omitempty"`
	RequestedStatus string `json:"requestedStatus,omitempty"`
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

func StopEncoder(token string, encoderUuid string) (bool, error) {
	var oEncoder Encoder
	oEncoder.RequestedStatus = "stop"

	jsonPayload, _ := json.Marshal(oEncoder)
	b := bytes.NewReader(jsonPayload)

	req, err := http.NewRequest("PATCH", "https://central.resi.io/api_v2.svc/encoders/"+encoderUuid, b)
	if err != nil {
		return false, errors.New("ERROR: failed to setup new HTTP request")
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization","X-Bearer "+token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return false, errors.New("ERROR: failed to connect to central.resi.io to stop encoder")
	}
	if resp.Body != nil {
		defer resp.Body.Close()
	}

	if resp.StatusCode != 200 {
		fmt.Println(resp.Status)
		return false, errors.New("ERROR: failed to stop encoder")
	}

	return true, nil
}

func StartEncoder(token string, encoderUuid string, eventProfileUuid string, presetUuid string) (bool, error) {
	var oEncoder Encoder
	oEncoder.StreamProfile.UUID = eventProfileUuid
	oEncoder.EncoderProfile.UUID = presetUuid
	oEncoder.RequestedStatus = "start"

	jsonPayload, _ := json.Marshal(oEncoder)
	b := bytes.NewReader(jsonPayload)

	req, err := http.NewRequest("PATCH", "https://central.resi.io/api_v2.svc/encoders/"+encoderUuid, b)
	if err != nil {
		return false, errors.New("ERROR: failed to setup new HTTP request")
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization","X-Bearer "+token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return false, errors.New("ERROR: failed to connect to central.resi.io to start encoder")
	}
	if resp.Body != nil {
		defer resp.Body.Close()
	}

	if resp.StatusCode != 200 {
		return false, errors.New("ERROR: failed to start encoder")
	}

	return true, nil
}
