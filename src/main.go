package main

import (
	"myproject/src/helpers"
	"myproject/src/models"

	"github.com/gin-gonic/gin"
)

type Block = models.Block
type Transaction = models.Transaction

var blockchain []Block

func getBlockChain(c *gin.Context) {
	c.JSON(200, blockchain)
}

func createTransaction(c *gin.Context) {
	var transaction Transaction
	c.BindJSON(&transaction)
	blockchain = helpers.AddNewBlock(blockchain, transaction)
	c.JSON(200, gin.H{"message": "Transaction added", "data": blockchain[len(blockchain)-1]})
}

func main() {
	// Initialize our blockchain with the genesis block
	blockchain = helpers.AddNewBlock(blockchain, Transaction{Amount: 100, Sender: "genesis", Receiver: "dev"})
	r := gin.Default()
	r.GET("/", getBlockChain)
	r.POST("/add", createTransaction)
	r.Run() // listen and serve on 0.0.0.0:8080
}
