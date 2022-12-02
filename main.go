package main

import (
    "net/http"

    "github.com/gin-gonic/gin"

	"math/rand"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!£$%^&*()_+-=¬[]{}~#,.<>?/*"

type Response struct {
	Responsecode int `json:"responsecode"`
	NumberOfChars	int	`json:"NumberOfChars"`
	Passwords	string	`json:"password"`
}

func buildRandomPassword(numberOfChars int) string{
	password := make([]byte, numberOfChars)
	for i := range password {
        password[i] = letterBytes[rand.Intn(len(letterBytes))]
    }

	return string(password)
}

// getRandomPassword
func getRandomPassword(c *gin.Context) {
	var response Response
	var numberOfChars int = 50

	response.Responsecode = http.StatusOK
	response.Passwords = buildRandomPassword(numberOfChars)
	response.NumberOfChars = numberOfChars
    c.IndentedJSON(http.StatusOK, response)
}

func main() {

    router := gin.Default()
    router.GET("/generate", getRandomPassword)

    router.Run("localhost:8080")
}
