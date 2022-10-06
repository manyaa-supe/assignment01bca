package main

import (
	"crypto/sha256"
	"fmt"
	"strconv"
	"strings"
)

var hashes [6]string
var last_block_hash string

type block struct {
	block_id int
	transaction_string       string
	prev_hash    string
	hash_block string
}

func CalculateHash(hashh string) string {

	return fmt.Sprintf("%x", sha256.Sum256([]byte(hashh)))
}

func NewBlock(nonce int, transaction_s string, prev_h string) *block {
	b := new(block)
	b.block_id = nonce
	b.transaction_string = transaction_s
	b.prev_hash = prev_h
	b.hash_block = CalculateHash(strconv.Itoa(nonce) + transaction_s + prev_h)
	if nonce != 5{
		hashes[nonce]= b.hash_block
	} else{
		last_block_hash = b.hash_block
		hashes[5] = last_block_hash
	}
	


	return b
}

type blockChain struct {
	list []*block
}

func (obj *blockChain) createBlock(nonce int, transaction_s string, prev_hash string)*block {

	block1 := NewBlock(nonce, transaction_s,prev_hash)
	obj.list = append(obj.list, block1)
	return block1
}

func ListBlocks(obj *blockChain, size int) {

	for i := 0; i < size; i++ {
		fmt.Printf("%s Block  %d %s\n", strings.Repeat("=", 25), i+1, strings.Repeat("=", 25))

		fmt.Println("Nonce: ", obj.list[i].block_id)
		fmt.Println("Transaction string: ", obj.list[i].transaction_string)
		fmt.Println("Previous Block Hash: ", obj.list[i].prev_hash)
		fmt.Println("Hash of Block: ", obj.list[i].hash_block)
		
	}

}

func VerifyChain(obj *blockChain, size int) {
	//hash_block = CalculateHash(strconv.Itoa(nonce) + transaction_s + prev_h)
	var hashBlock [5]string 
	 var count int
	 count = 0

	for i := 0; i < size; i++ {
		hashBlock[i] = CalculateHash(strconv.Itoa(obj.list[i].block_id) + obj.list[i].transaction_string + obj.list[i].prev_hash)
		

		fmt.Printf("Hash of Block %d:  %s \n",i+1, obj.list[i].hash_block)

		fmt.Println("Calculated Hash: ", hashBlock[i])
		if(obj.list[i].hash_block == hashBlock[i]){
			count++
		}
			
		
	}
	if(count == 5){
		fmt.Println("\nBlockchain not modified!")
	}else{
		fmt.Println("\nBlockchain modified!")
	}
}
/*func VerifyChain(nonce int, transaction_s string, prev_h string ) {
	hash_block := CalculateHash(strconv.Itoa(nonce) + transaction_s + prev_h)
	if(hash_block == last_block_hash){
		fmt.Println("\nBlockchain not modified!")
	}else
	{
		fmt.Println("\nBlockchain modified!")
	}
}*/

func ChangeBlock(nonce int, transaction string, obj *blockChain, size int ){
	//update all?
	var newHash string
	for i := 0; i < size; i++ {
		if(i == nonce-1){
			obj.list[i].transaction_string = transaction
			newHash = CalculateHash(strconv.Itoa(obj.list[i].block_id) + obj.list[i].transaction_string + obj.list[i].prev_hash)
			hashes[nonce] = newHash

		}
	}

}

func main() {

	hashes[0]="xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
	BChain := new(blockChain)

	BChain.createBlock(1, "Sheldon to Penny", hashes[0])
	BChain.createBlock(2, "Penny to Bernadette", hashes[1])
	BChain.createBlock(3, "Bernadette to Howard", hashes[2])
	BChain.createBlock(4, "Howard to Rajesh", hashes[3])
	BChain.createBlock(5, "Rajesh to Amy", hashes[4])

	ListBlocks(BChain, 5)
	fmt.Printf("%s Verify 1 %s\n", strings.Repeat("=", 25), strings.Repeat("=", 25))
	VerifyChain(BChain, 5)


	ChangeBlock(2,"Bernadette to Denise", BChain,5)

	fmt.Printf("%s Verify 2 %s\n", strings.Repeat("=", 25), strings.Repeat("=", 25))
	VerifyChain(BChain, 5)
}