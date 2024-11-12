package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/MicahParks/keyfunc/v3"
	"github.com/golang-jwt/jwt/v5"
)

// JWTValidator validates a JWT token.in the Authorization header.
func JWTValidator(issuer string, jwksURL string) gin.HandlerFunc {
	jwks, _ := keyfunc.NewDefault([]string{jwksURL})

	return func(c *gin.Context) {
		authorization := c.GetHeader("Authorization")
		if authorization == "" {
			c.JSON(401, gin.H{"error": "Authorization header is missing"})
			c.Abort()
			return
		}

		if !strings.HasPrefix(authorization, "Bearer ") {
			c.JSON(401, gin.H{"error": "Authorization header is not a bearer token"})
			c.Abort()
			return
		}

		token := strings.TrimPrefix(authorization, "Bearer ")

		t, err := jwt.Parse(token, jwks.Keyfunc)
		if err != nil {
			c.JSON(401, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		if !t.Valid {
			c.JSON(401, gin.H{"error": "Token is invalid"})
			c.Abort()
			return
		}

		claims := t.Claims.(jwt.MapClaims)
		if claims["iss"] != issuer {
			c.JSON(401, gin.H{"error": "Invalid issuer"})
			c.Abort()
			return
		}

		c.Next()
	}
}
