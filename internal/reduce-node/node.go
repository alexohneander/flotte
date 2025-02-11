package reducenode

import "github.com/gin-gonic/gin"

func Start() {
	err := registerAsReduceNode()
	if err != nil {
		panic(err)
	}

	router := gin.Default()
	router = registerRoutes(router)

	router.Run("localhost:4001") // listen and serve on 0.0.0.0:4001
}

func registerAsReduceNode() error {
	return nil
}
