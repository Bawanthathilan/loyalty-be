// controllers/location.go
package controllers

import (
	"context"
	"net/http"
	"os"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	square "github.com/square/square-go-sdk"
	client "github.com/square/square-go-sdk/client"
	loyalty "github.com/square/square-go-sdk/loyalty"
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
			"mappings":   acct.Mapping, 
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

	sess := sessions.Default(c)
	sess.Set("account_id", resp.LoyaltyAccount.ID)
	sess.Set("program_id", resp.LoyaltyAccount.ProgramID)
	if err := sess.Save(); err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"error": "could not save session"})
    return
  }

	// 4) Return the account details
	acct := resp.LoyaltyAccount
	c.JSON(http.StatusOK, gin.H{
		"response": acct,
	})
}

func Earn(c *gin.Context) {
	// 1) Read the account ID from the path
	account_id := c.Param("account_id")

	if account_id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "account_id is required"})
		return
	}

	// 2) Initialize Square client
	client := client.NewClient(
		option.WithToken(os.Getenv("SQUARE_ACCESS_TOKEN")),
		option.WithBaseURL(square.Environments.Sandbox),
	)

	var reqBody loyalty.AccumulateLoyaltyPointsRequest
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := client.Loyalty.Accounts.AccumulatePoints(context.TODO(), &loyalty.AccumulateLoyaltyPointsRequest{
			AccountID: account_id,
            IdempotencyKey: uuid.New().String(),
            LocationID: reqBody.LocationID,
			AccumulatePoints: &square.LoyaltyEventAccumulatePoints{
                Points: square.Int(
                    *reqBody.AccumulatePoints.Points,
                ),
            },
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"response": resp,
	})

}

func Redeem(c *gin.Context) {

	// 1) Read the account ID from the path
	reward_id := c.Param("reward_id")

	if reward_id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "reward_id is required"})
		return
	}

	// 2) Initialize Square client
	client := client.NewClient(
		option.WithToken(os.Getenv("SQUARE_ACCESS_TOKEN")),
		option.WithBaseURL(square.Environments.Sandbox),
	)

	var reqBody loyalty.RedeemLoyaltyRewardRequest
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := client.Loyalty.Rewards.Redeem(context.TODO(), &loyalty.RedeemLoyaltyRewardRequest{
		RewardID:       reward_id,
		LocationID:     reqBody.LocationID,
		IdempotencyKey: uuid.New().String(),
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"response": resp,
	})

}

func Balance(c *gin.Context) {

}

func History(c *gin.Context) {

}
