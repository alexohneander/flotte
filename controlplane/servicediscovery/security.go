package servicediscovery

import (
	b64 "encoding/base64"
	"log"
	"math/rand"
	"os"
)

func InitializedToken() string {
	// Check if token.yaml exists
	if _, err := os.Stat("/tmp/token.yaml"); err == nil {
		log.Println("ServiceDiscovery: Token file exists")
		// If yes, read token.yaml
		token := GetToken()
		return token
	} else if os.IsNotExist(err) {
		log.Println("ServiceDiscovery: Token file does not exist")

		// If not, create token.yaml
		token := createToken()
		log.Println("ServiceDiscovery: Token was created")

		// Write token to token.yaml
		err := os.WriteFile("/tmp/token.yaml", []byte(token), 0644)
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
	if _, err := os.Stat("/tmp/token.yaml"); err == nil {
		tokenByte, err := os.ReadFile("/tmp/token.yaml")
		if err != nil {
			log.Fatal(err)
		}
		token := string(tokenByte)
		return token
	} else if os.IsNotExist(err) {
		log.Println("ServiceDiscovery: Token file does not exist")
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
