package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/MrM21632/snowball/snowball"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func prometheusHandler() gin.HandlerFunc {
	h := promhttp.Handler()
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func main() {
	// Change to false if you want to provide pre-assigned IDs for servers
	node, err := snowball.InitNode(true)
	if err != nil {
		fmt.Printf("Encountered error while initializing node: %s", err)
		return
	}

	r := gin.Default() // Request routing
	p := gin.New()     // Metrics routing

	p.Use(prometheusHandler())
	r.SetTrustedProxies(nil)

	r.POST("/generate", func(c *gin.Context) {
		id := node.GenerateID()
		c.JSON(http.StatusOK, gin.H{"id": strconv.FormatUint(uint64(id), 10)})
	})

	go func() {
		p.Run(":9100")
	}()
	r.Run(":8080")
}
