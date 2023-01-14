package resi

import (
	"net/http"
	"errors"
	"io"
	"encoding/json"
)

type EventProfile struct {
	UUID string `json:"uuid"`
	Name string `json:"name"`
	Description string `json:"description"`
}

func EventProfiles(token string, customerId string) ([]EventProfile, error) {
	httpRequest, httpRequestErr := http.NewRequest("GET", "https://central.resi.io/api/v3/customers/"+customerId+"/eventprofiles", nil)
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

	var eventProfiles []EventProfile
	eventProfilesErr := json.Unmarshal(respBody, &eventProfiles)
	if eventProfilesErr != nil {
		return nil, errors.New("ERROR: failed to read JSON response from resi.io for encoders")
	}
	return eventProfiles, nil
}

//https://central.resi.io/api/v3/customers/60e5326c-791f-47f8-aaef-00c26338c880/eventprofiles
//Authorization: X-Bearer <token>
/*
[{
	"uuid": "3c184517-0d81-4813-90fc-cc29964d9e17",
	"name": "Default Encoder Event Profile",
	"description": "Default Encoder Event Profile",
	"deleteAfter": 10,
	"lanOnly": false,
	"bucket": "lao-default_encoder_event_profile_imf",
	"regionId": "1290ae83-b319-4cf4-9069-41a7296c8c2c"
}]*/