package models

type Transaction struct {
	Amount   int
	Sender   string
	Receiver string
}

type Block struct {
	PrevHash    string
	Transaction Transaction
	Timestamp   int
}
