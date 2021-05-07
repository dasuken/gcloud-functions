package gcloud_functions

import (
	"cloud.google.com/go/storage"
	"context"
	"fmt"
	"net/http"
	"os"
	"time"
)

func TriggerHTTPToBucket(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		names, err := r.URL.Query()["name"]
		if !err || len(names[0]) < 1 {
			fmt.Fprint(w, "パラメータに\"name\"がありません。\r\n")
			return
		}

		WriteBucket(w, names[0])
	default:
		http.Error(w, "405 - Method Not Allowd", http.StatusMethodNotAllowed)
	}
}

func WriteBucket(w http.ResponseWriter, name string) {
	bucketName := os.Getenv("BUCKET_NAME")

	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		fmt.Fprintln(w, "Storage 接続エラー", err)
	}
	defer client.Close()

	objectName := time.Now().Format("20060102150405")

	fw := client.Bucket(bucketName).Object(objectName).NewWriter(ctx)
	if _, err := fw.Write([]byte(name + "\r\n")); err != nil {
		fmt.Fprint(w, "オブジェクト書き込みエラー　エラー: %s", err)
		return
	}

	if err := fw.Close(); err != nil {
		fmt.Fprint(w, "オブジェクト切断エラー  エラー: %s", err)
		return
	}
}