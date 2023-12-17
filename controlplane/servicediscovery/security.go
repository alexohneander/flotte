package servicediscovery

import (
	b64 "encoding/base64"
	"log"
	"math/rand"
	"os"
)

var tokenPath string = "/tmp/token.yaml"

func InitializedToken() string {
	// Check if token.yaml exists
	if _, err := os.Stat(tokenPath); err == nil {
		log.Println("Service Discovery: Token file exists")
		// If yes, read token.yaml
		token := GetToken()
		return token
	} else if os.IsNotExist(err) {
		log.Println("Service Discovery: Token file does not exist")

		// If not, create token.yaml
		token := createToken()
		log.Println("Service Discovery: Token was created")

		// Write token to token.yaml
		err := os.WriteFile(tokenPath, []byte(token), 0644)
		if err != nil {
			log.Fatal(err)
		}

		return token
	}

	return ""
}

func createToken() string {
	// Create Secure Token and encode it to base64

	// create 64bit random string
	tokenPlain := randSeq(64)

	// Encode token with base64
	token := b64.StdEncoding.EncodeToString([]byte(tokenPlain))
	return token
}

func GetToken() string {
	// Check if token.yaml exists
	if _, err := os.Stat(tokenPath); err == nil {
		tokenByte, err := os.ReadFile(tokenPath)
		if err != nil {
			log.Fatal(err)
		}
		token := string(tokenByte)
		return token
	} else if os.IsNotExist(err) {
		log.Println("Service Discovery: Token file does not exist")
		return ""
	}
	return ""
}

func randSeq(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
