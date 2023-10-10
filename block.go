package assignment01bca

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
)

var b bool
var verifyChain bool

type Node struct {
	transaction  string
	nonce        int
	previousHash string
	next         *Node
}

type List struct {
	head *Node
}

func (l *List) NewBlock_(transaction string, nonce int) {
	newNode := &Node{transaction: transaction, nonce: nonce}

	if l.head == nil {
		l.head = newNode
		newNode.previousHash = ""
		return
	}

	curr := l.head
	for curr.next != nil {
		curr = curr.next
	}

	hashString := l.CalculateHash(curr.transaction, curr.nonce, curr.previousHash)
	curr.next = newNode
	newNode.previousHash = hashString
}

func (l *List) NewBlock(transaction string, nonce int, previousHash string) {

	newNode := &Node{transaction: transaction, nonce: nonce, previousHash: previousHash}

	if l.head == nil {
		l.head = newNode
		return
	}

	curr := l.head
	for curr.next != nil {
		curr = curr.next
	}
	curr.next = newNode
}

func (l *List) DisplayBlock() {
	if l.head == nil {
		fmt.Printf("Block-chain is Empty")
	}
	curr := l.head

	for curr != nil {
		fmt.Println("Transaction: " + curr.transaction)
		fmt.Println("Nonce: " + strconv.Itoa(curr.nonce))
		fmt.Println("Previous Hash: " + curr.previousHash)

		hashString := l.CalculateHash(curr.transaction, curr.nonce, curr.previousHash)
		fmt.Println("Current Block Hash: " + hashString)
		curr = curr.next
	}
}
func (l *List) ChangeBlock(currentBlockHash string, newTransaction string) {

	b = false

	if l.head == nil {
		fmt.Printf("Block-chain is Empty")
	}
	curr := l.head

	for curr != nil {
		hashString := l.CalculateHash(curr.transaction, curr.nonce, curr.previousHash)
		if currentBlockHash == hashString {
			curr.transaction = newTransaction
			b = true
			break
		}
		curr = curr.next
	}

	if b {
		curr = l.head
		for curr.next != nil {
			hashString := l.CalculateHash(curr.transaction, curr.nonce, curr.previousHash)
			prev := hashString
			curr = curr.next
			curr.previousHash = prev
		}
	}
}
func (l *List) VerifyChain() {
	verifyChain = false
	if l.head == nil {
		fmt.Printf("Block-chain is Empty")
	}
	curr := l.head
	for curr.next != nil {
		hashString := l.CalculateHash(curr.transaction, curr.nonce, curr.previousHash)
		curr = curr.next
		if curr.previousHash == hashString {
			verifyChain = true
			continue
		} else {
			verifyChain = false
			fmt.Printf("Block-chain is Tempered")
			break
		}
	}
	if b {
		fmt.Printf("Block-chain is Verified and Not Tempered")
	}
}
func (l *List) CalculateHash(transaction string, nonce int, previousHash string) string {
	s := transaction + previousHash + strconv.Itoa(nonce)
	h := sha256.New()
	h.Write([]byte(s))
	bs := h.Sum(nil)
	hashString := hex.EncodeToString(bs)
	return hashString
}
