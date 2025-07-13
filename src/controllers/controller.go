// controllers/location.go
package controllers

import (
	"context"
	"net/http"
	"os"

	"loyalty-be/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	square "github.com/square/square-go-sdk"
	client "github.com/square/square-go-sdk/client"
	loyalty "github.com/square/square-go-sdk/loyalty"
	option "github.com/square/square-go-sdk/option"

	"loyalty-be/models"
)



type loginRequest struct {
    PhoneNumber string `json:"phone_number" binding:"required"`
    Password    string `json:"password"    binding:"required"`
}

func Login(c *gin.Context) {
    // Bind JSON body
    var req loginRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "phone_number and password are required"})
        return
    }

    // Find matching account in our hard-coded slice
    var acctID string
    for _, acct := range models.LoyaltyAccounts {
        if acct.Mapping.PhoneNumber == req.PhoneNumber {
            if acct.Mapping.Password != req.Password {
                c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
                return
            }
            acctID = acct.ID
            break
        }
    }
    if acctID == "" {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
        return
    }

    // Initialize Square client
    sq := client.NewClient(
        option.WithToken(os.Getenv("SQUARE_ACCESS_TOKEN")),
        option.WithBaseURL(square.Environments.Sandbox),
    )

    // Call the Square Loyalty API
    sdkReq := &loyalty.GetAccountsRequest{AccountID: acctID}
    sdkResp, err := sq.Loyalty.Accounts.Get(context.Background(), sdkReq)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    // (Optional) save session
    utils.SaveUserSession(c, acctID)

    // Return the SDK’s response object
    c.JSON(http.StatusOK, gin.H{"response": sdkResp.LoyaltyAccount})
}


func Earn(c *gin.Context) {
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
			AccountID: utils.GetSessionData(c),
            IdempotencyKey: uuid.New().String(),
            LocationID: os.Getenv("SQUARE_LOCATION_ID"),
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

	resp, err := client.Loyalty.Rewards.Redeem(context.TODO(), &loyalty.RedeemLoyaltyRewardRequest{
		RewardID:       reward_id,
		LocationID:    os.Getenv("SQUARE_LOCATION_ID"),
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
	// 2) Initialize Square client
	client := client.NewClient(
		option.WithToken(os.Getenv("SQUARE_ACCESS_TOKEN")),
		option.WithBaseURL(square.Environments.Sandbox),
	)

	resp, err := client.Loyalty.Accounts.Get(context.TODO(), &loyalty.GetAccountsRequest{
		AccountID: utils.GetSessionData(c),
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"response": resp.LoyaltyAccount,
	})
}

func History(c *gin.Context) {

	var reqBody square.SearchLoyaltyEventsRequest

	client := client.NewClient(
		option.WithToken(os.Getenv("SQUARE_ACCESS_TOKEN")),
		option.WithBaseURL(square.Environments.Sandbox),
	)

	resp, err := client.Loyalty.SearchEvents(context.Background(), &reqBody)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var accounts []gin.H
	for _, acct := range resp.Events {
		accounts = append(accounts, gin.H{
			"response": acct,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"accounts": accounts,
		"cursor":   resp.Cursor,
	})
}