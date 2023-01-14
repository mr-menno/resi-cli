package resi

import (
	"net/http"
	"errors"
	"io"
	"encoding/json"
)

type EncoderProfile struct {
	UUID string `json:"uuid"`
	Name string `json:"name"`
	CustomerId string `json:"customerId"`
	SerialNumber string `json:"serialNumber"`
	Status string `json:"status"`
	OperationalState string `json:"operationalState"`
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

//https://central.resi.io/api/v3/customers/60e5326c-791f-47f8-aaef-00c26338c880/encoderprofiles
//Authorization: X-Bearer <token>
/*
[{
	"uuid": "d50e1bad-0afb-43a6-a0ee-00ff75f40114",
	"name": "1-1080p30 4Mb H264 2-Ch - Default",
	"type": "DASH",
	"inputOffset": null,
	"audio": {
		"channels": 2.0,
		"bitratePerChan": 64
	},
	"video": [{
		"input": [{
			"format": "Hp30"
		}],
		"format": "Hp30",
		"bitRate": 4.0,
		"codec": "H264",
		"bitDepth": 8,
		"hardwareAcceleration": true
	}],
	"customName": "Default",
	"interlaced": false,
	"shortSegments": false,
	"customProfile": false,
	"lcevc": false,
	"colorspace": null
}]*/