package gcloud_functions

import (
	"cloud.google.com/go/functions/metadata"
	"context"
	"fmt"
	"log"
	"time"
)

type GCSEvent struct {
	Bucket         string    `json:"bucket"`
	Name           string    `json:"name"`
	Metageneration string    `json:"metageneration"`
	ResourceState  string    `json:"resource_state"`
	TimeCreated    time.Time `json:"time_created"`
	Updated        time.Time `json:"updated"`
}

func TriggerStorage(ctx context.Context, e GCSEvent) error {
	meta, err := metadata.FromContext(ctx)
	if err != nil {
		return fmt.Errorf("metadata.FromContext: %v", err)
	}

	log.Printf("Event ID: %v\n", meta.EventID)
	log.Printf("Event type: %v\n", meta.EventType)
	log.Printf(" バケット名 : %v\n", e.Bucket)
	log.Printf(" ファイル名 : %v\n", e.Name)
	log.Printf(" リソース状態 : %v\n", e.ResourceState)
	log.Printf(" 作成日時 : %v\n", e.TimeCreated)
	log.Printf(" 更新日時 : %v\n", e.Updated)
	return nil
}