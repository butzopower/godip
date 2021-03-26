package main

import (
	"example/dip/core"
	"example/dip/db"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

func main() {
	bank, cleanUp, err := db.NewDbBank()

	if err != nil {
		log.Fatal(err)
	}

	defer cleanUp()

	r := gin.Default()
	r.GET("/balance", func(c *gin.Context) {
		accountNumber := c.Query("account")
		balanceAmount := core.Balance(bank)(accountNumber)
		c.JSON(200, gin.H{
			"balance": balanceAmount,
		})
	})

	r.GET("/deposit", func(c *gin.Context) {
		accountNumber := c.Query("account")
		amount, _ := strconv.ParseInt(c.Query("amount"), 10, 0)

		core.Deposit(bank)(accountNumber, int(amount))

		c.JSON(201, gin.H{})
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
