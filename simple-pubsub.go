package gcloud_functions

import (
	"context"
	"encoding/json"
	"log"
)

type PubSubMessage struct {
	Data []byte `json:"data"`
}

type Info struct {
	Name string `json:"name"`
	Place string `json:"place"`
}

func TriggerPubSub(ctx context.Context, m PubSubMessage) error {
	var i Info

	err := json.Unmarshal(m.Data, &i)

	if err != nil {
		log.Printf("Error: %T, message: %v", err, err)
		return nil
	}

	log.Printf("こんにちはName: %sさん！, Place: %sへPubSubからvia Functions", i.Name, i.Place)
	return nil
}