package main

import (
	// "fmt"
	"context"
	"log"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

func main() {
	// Use a service account
	ctx := context.Background()
	sa := option.WithCredentialsFile("./grandcolline-dev-87202938d0c8.json")
	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		log.Fatalln(err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	// データ追加
	_, _, err = client.Collection("task").Add(ctx, map[string]interface{}{
		"name":   "Ada",
		"status": "Lovelace",
	})
	if err != nil {
		log.Fatalf("Failed adding alovelace: %v", err)
	}

	// 切断
	defer client.Close()
}
