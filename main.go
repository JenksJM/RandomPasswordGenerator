package main

import (
    "net/http"

    "github.com/gin-gonic/gin"

	"math/rand"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

type Response struct {
	Responsecode int `json:"responsecode"`
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
	response.Responsecode = http.StatusOK
	response.Passwords = buildRandomPassword(50)
    c.IndentedJSON(http.StatusOK, response)
}

func main() {

    router := gin.Default()
    router.GET("/generate", getRandomPassword)

    router.Run("localhost:8080")
}
