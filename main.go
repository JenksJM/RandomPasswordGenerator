package main

import (
    "net/http"

    "github.com/gin-gonic/gin"

	"math/rand"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

type Password struct {
	Password	string	`json:"password"`
}

func buildRandomPassword(numberOfChars int) Password{
	var randomPassword Password

	password := make([]byte, numberOfChars)
	for i := range password {
        password[i] = letterBytes[rand.Intn(len(letterBytes))]
    }

	randomPassword.Password = string(password)
	return randomPassword
}

// getRandomPassword
func getRandomPassword(c *gin.Context) {
	var response = []Password{
		buildRandomPassword(50),
	}
    c.IndentedJSON(http.StatusOK, response)
}

func main() {

    router := gin.Default()
    router.GET("/generate", getRandomPassword)

    router.Run("localhost:8080")
}
