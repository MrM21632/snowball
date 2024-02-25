package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/MrM21632/snowball/snowball"
	"github.com/gin-gonic/gin"
)

func main() {
	// Change to false if you want to provide pre-assigned IDs for servers
	node, err := snowball.InitNode(true)
	if err != nil {
		fmt.Printf("Encountered error while initializing node: %s", err)
		return
	}

	r := gin.Default()
	r.SetTrustedProxies(nil)

	r.POST("/generate", func(c *gin.Context) {
		id := node.GenerateID()
		c.JSON(http.StatusOK, gin.H{"id": strconv.FormatUint(uint64(id), 10)})
	})

	r.Run(":8080")
}
