package utils

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// Session key constants
const (
	SessionKeyAccountID = "account_id"
)

// SaveUserSession sets the account and location IDs into the session and saves it.
func SaveUserSession(c *gin.Context, accountID string) error {
	sess := sessions.Default(c)
	sess.Set(SessionKeyAccountID, accountID)
	return sess.Save()
}

// GetSessionData retrieves the account and location IDs from the session.
func GetSessionData(c *gin.Context) (accountID string) {
	sess := sessions.Default(c)
	if v := sess.Get(SessionKeyAccountID); v != nil {
		accountID, _ = v.(string)
	}
	return
}

// RequireSession is a middleware that ensures a user is logged in.
func RequireSession() gin.HandlerFunc {
	return func(c *gin.Context) {
		accountID := GetSessionData(c)
		if accountID == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			return
		}
		c.Next()
	}
}
