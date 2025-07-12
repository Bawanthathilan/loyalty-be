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

func Search(c *gin.Context) {

    var reqBody loyalty.SearchLoyaltyAccountsRequest
    if err := c.ShouldBindJSON(&reqBody); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    client := client.NewClient(
        option.WithToken(os.Getenv("SQUARE_ACCESS_TOKEN")),
        option.WithBaseURL(square.Environments.Sandbox),
    )

    resp, err := client.Loyalty.Accounts.Search(context.Background(), &reqBody)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

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

    c.JSON(http.StatusOK, gin.H{
        "accounts": accounts,
        "cursor":   resp.Cursor,
    })
}

func Login(c *gin.Context) {

    // 1) Read the account ID from the path
    accountID := c.Param("account_id")

    if accountID == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "account_id is required"})
        return
    }

    // 2) Initialize Square client
    client := client.NewClient(
        option.WithToken(os.Getenv("SQUARE_ACCESS_TOKEN")),
        option.WithBaseURL(square.Environments.Sandbox),
    )

    // 3) Call the Get endpoint
    req := &loyalty.GetAccountsRequest{AccountID: accountID}
    resp, err := client.Loyalty.Accounts.Get(context.Background(), req)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    // 4) Return the account details
    acct := resp.LoyaltyAccount
    c.JSON(http.StatusOK, gin.H{
       "response": acct,
    })
}

func Earn(c *gin.Context) {

}

func Redeem(c *gin.Context) {

}

func Balance(c *gin.Context) {

}

func History(c *gin.Context) {

}