package server

import (
	"fmt"
	"jwt-test/internal/user"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// Add a new global variable for the secret key
var secretKey = []byte("your-secret-key")

// Function to create JWT tokens with claims
func CreateToken(username string) (string, error) {
	// Create a new JWT token with claims
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": username,                         // Subject (user identifier)
		"iss": "jwt-test",                       // Issuer
		"aud": user.GetRole(username),           // Audience (user role)
		"exp": time.Now().Add(time.Hour).Unix(), // Expiration time
		"iat": time.Now().Unix(),                // Issued at
	})

	// Print information about the created token
	fmt.Printf("Token claims added: %+v\n", claims)

	tokenString, err := claims.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func HandleLogin(c *gin.Context) {
	username := c.PostForm("email")
	password := c.PostForm("password")

	// Dummy credential check
	if user.ValidateUser(username, password) {
		tokenString, err := CreateToken(username)
		if err != nil {
			c.String(http.StatusInternalServerError, "Error creating token")
			return
		}

		fmt.Printf("Token created: %s\n", tokenString)
		c.SetCookie("token", tokenString, 3600, "/", "localhost", false, true)
		c.Redirect(http.StatusFound, "/")
	} else {
		c.String(http.StatusUnauthorized, "Invalid credentials")
	}
}

// Function to verify JWT tokens
func AuthenticateMiddleware(c *gin.Context) {
	// Retrieve the token from the cookie
	tokenString, err := c.Cookie("token")
	if err != nil {
		fmt.Println("Token missing in cookie")
		c.Redirect(http.StatusSeeOther, "/login")
		c.Abort()
		return
	}

	// Verify the token
	token, err := verifyToken(tokenString)
	if err != nil {
		fmt.Printf("Token verification failed: %v\\n", err)
		c.Redirect(http.StatusSeeOther, "/login")
		c.Abort()
		return
	}

	// Print information about the verified token
	fmt.Printf("Token verified successfully. Claims: %+v\\n", token.Claims)

	// Continue with the next middleware or route handler
	c.Next()
}

// Function to verify JWT tokens
func verifyToken(tokenString string) (*jwt.Token, error) {
	// Parse the token with the secret key
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	// Check for verification errors
	if err != nil {
		return nil, err
	}

	// Check if the token is valid
	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	// Return the verified token
	return token, nil
}
