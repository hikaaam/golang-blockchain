package helpers

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"myproject/src/models"
	"time"
)

type Block = models.Block
type Transaction = models.Transaction

func Hash(s string) string {
	h := sha256.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

func JsonToString(j interface{}) string {
	b, _ := json.Marshal(j)
	return string(b)
}

func GetTS() int {
	return int(time.Now().Unix())
}

func AddNewBlock(blockchain []Block, transaction models.Transaction) []Block {
	var lastHash string

	if len(blockchain) > 0 {
		lastHash = Hash(JsonToString(blockchain[len(blockchain)-1]))
	} else {
		lastHash = Hash("genesis")
	}
	var newblock Block = Block{PrevHash: lastHash, Transaction: transaction, Timestamp: GetTS()}
	blockchain = append(blockchain, newblock)
	fmt.Println("")
	fmt.Println("🔨 New block created ->", Hash(JsonToString(newblock)))
	fmt.Println("")
	return blockchain
}
