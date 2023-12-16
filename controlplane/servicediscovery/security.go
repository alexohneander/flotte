package servicediscovery

import (
	"log"
	"os"
)

func InitializedToken() string {
	// Check if token.yaml exists
	if _, err := os.Stat("/tmp/token.yaml"); err == nil {
		log.Println("ServiceDiscovery: Token file exists")
		// If yes, read token.yaml
		tokenByte, err := os.ReadFile("/tmp/token.yaml")
		if err != nil {
			log.Fatal(err)
		}

		token := string(tokenByte)
		log.Println("ServiceDiscovery: Token is: ", token)
		return token
	} else if os.IsNotExist(err) {
		log.Println("ServiceDiscovery: Token file does not exist")

		// If not, create token.yaml
	}

	return ""
}
