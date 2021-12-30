package main

import (
	"fmt"
	"com.github/jatin837/blockchain"
)

func main(){
	blockchain := CreateBlockchain(2)
	blockchain.addBlock("Jon Jones", "Daniel Cormier", 5)
	blockchain.addBlock("Anthony Jhonson", "Anderson Silva", 2)

	fmt.Println(blockchain.isValid())
	fmt.Println(blockchain.chain)
}
