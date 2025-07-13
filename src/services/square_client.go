// src/services/square_client.go
package services

import (
	"os"
	"sync"

	square "github.com/square/square-go-sdk"
	client "github.com/square/square-go-sdk/client"
	option "github.com/square/square-go-sdk/option"
)

var (
	squareClient *client.Client
	once         sync.Once
)

func GetSquareClient() *client.Client {
	once.Do(func() {
		squareClient = client.NewClient(
			option.WithToken(os.Getenv("SQUARE_ACCESS_TOKEN")),
			option.WithBaseURL(square.Environments.Sandbox),
		)
	})
	return squareClient
}
