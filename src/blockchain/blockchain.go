package blockchain

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type Block struct {
	data	map[string]interface{}
	hash	string
	prevHash	string
	timeStamp	time.Time
	pow	int
}

type Blockchain struct {
	genesisBlock	Block
	chain					[]Block
	difficulty		int
}

func (b Block) calcHash() string {
	data, _ := json.Marshal(b.data)
	blockData := b.prevHash + string(data) + b.timeStamp.String() + strconv.Itoa(b.pow)
	blockHash	:= sha256.Sum256([]byte(blockData))
	return fmt.Sprintf("%x", blockHash)
}

func (b *Block) mine(difficulty int) {
	for !strings.HasPrefix(b.hash, strings.Repeat("0", difficulty)){
		b.pow++
		b.hash = b.calcHash()
	}
}

func CreateBlockchain(difficulty int) Blockchain {
	genesisBlock := Block {
		hash:	"0",
		timeStamp: time.Now(),
	}
	return Blockchain{
		genesisBlock,
		[]Block{genesisBlock},
		difficulty,
	}
}

func (b *Blockchain) addBlock(from string, to string, amt float64){
	blockData := map[string]interface{}{
		"from":from,
		"to":to,
		"amt":amt,
	}
	lastBlock := b.chain[len(b.chain)-1]
	newBlock := Block{
		data:	blockData,
		prevHash: lastBlock.hash,
		timeStamp: time.Now(),
	}
	newBlock.mine(b.difficulty)
	b.chain = append(b.chain, newBlock)
}

func (b Blockchain) isValid() bool {
	for i:= range b.chain[1:] {
		prevBlock := b.chain[i]
		currBlock := b.chain[i+1]
		if currBlock.hash != currBlock.calcHash() || currBlock.prevHash != prevBlock.hash {
			return false
		}
	}
	return true
}

