// controllers/location.go
package controllers

import (
	"context"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	square "github.com/square/square-go-sdk"
	client "github.com/square/square-go-sdk/client"
	"github.com/square/square-go-sdk/loyalty"
	option "github.com/square/square-go-sdk/option"
)

// GetLocations handles GET /api/v1/locations
func GetLocations(c *gin.Context) {
	// Initialize Square client
	client := client.NewClient(
		option.WithToken(
			os.Getenv("SQUARE_ACCESS_TOKEN"),
		),
		option.WithBaseURL(square.Environments.Sandbox),
	)

	// Call the List locations endpoint
	resp, err := client.Locations.List(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Transform into a simple JSON payload
	var out []gin.H
	for _, loc := range resp.Locations {
		out = append(out, gin.H{
			"id":      *loc.ID,
			"name":    *loc.Name,
			"address": *loc.Address.AddressLine1,
			"city":    *loc.Address.Locality,
		})
	}

	c.JSON(http.StatusOK, gin.H{"locations": out})
}

// / SearchLoyalty handles POST /api/v1/searchloyalty
func SearchLoyalty(c *gin.Context) {
	// 1) Bind incoming JSON into the SDK request struct.
	//    If the user POSTs `{}` or omits `query`, this will work too.
	var reqBody loyalty.SearchLoyaltyAccountsRequest
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 2) Initialize the Square client
	client := client.NewClient(
		option.WithToken(os.Getenv("SQUARE_ACCESS_TOKEN")),
		option.WithBaseURL(square.Environments.Sandbox),
	)

	// 3) Call the Search endpoint
	resp, err := client.Loyalty.Accounts.Search(context.Background(), &reqBody)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 4) Transform the response into a simple JSON payload
	var accounts []gin.H
	for _, acct := range resp.LoyaltyAccounts {
		accounts = append(accounts, gin.H{
			"id":         *acct.ID,
			"program_id": acct.ProgramID,
			"points":     acct.LifetimePoints,
			"mappings":   acct.Mapping, // phone/customer mappings
			"created_at": *acct.CreatedAt,
		})
	}

	// If there’s more data, Square returns a cursor for pagination
	c.JSON(http.StatusOK, gin.H{
		"accounts": accounts,
		"cursor":   resp.Cursor,
	})
}
