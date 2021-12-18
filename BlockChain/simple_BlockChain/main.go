package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	badger "github.com/dgraph-io/badger/v3"
	logic "github.com/fodedoumbouya/GO-Guide-Doumbouya/BlockChain/simple_BlockChain/logic"
	"github.com/gorilla/mux"
)

var blockChain *logic.BlockChain
var db *badger.DB

func ErrorHandle(err error) {
	if err != nil {
		fmt.Println("Fail to get Data: ", err)
	}
}
func HandleAddBlock(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var vars = mux.Vars(r)

	data := vars["block"]
	if len(data) > 0 {
		block := addBlock(data)

		json.NewEncoder(w).Encode(block)
	} else {
		json.NewEncoder(w).Encode("Fail to create Block")

	}

}

func HandleGetBlock(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// var blockData map[string]interface{}
	// json.NewDecoder(r.Body).Decode(&blockData)
	// data := fmt.Sprintf("%s", blockData["data"])
	json.NewEncoder(w).Encode(blockChain.Blocks)

}

func addBlock(data string) logic.Block {
	block := blockChain.AddBlock(data, len(blockChain.Blocks)+1)
	logic.AddBlockInDB(db, *block)

	return *block
}

func login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var vars = mux.Vars(r)
	typeLogin := vars["type"]

	json.NewEncoder(w).Encode(typeLogin)
}

func HandleFuncRequest() {
	start := time.Now()
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/addBlock/block/{block}", HandleAddBlock)
	myRouter.HandleFunc("/getBlock", HandleGetBlock)

	port := "8080"
	fmt.Println("Server running at port", port)
	fmt.Println("----------------RunTime Duration---------------------", time.Since(start).Seconds())
	log.Fatal(http.ListenAndServe(":"+port, myRouter))

}

func main() {
	db1, err := badger.Open(badger.DefaultOptions("/data/cache"))
	db = db1

	if err != nil {
		log.Fatal("Open DB error: ", err)
	}
	defer db.Close()

	res, data := logic.GetInitFromDB(db)
	if res {
		blockChain = data
		fmt.Println("Had Data before")
	} else {
		fmt.Println("No Data before")
		blockChain = logic.InitBlockChain()
		logic.AddBlockInDB(db, *blockChain.Blocks[0])
	}

	// list := []string{"Block 1"} // "Block 2", "Block 3", "Block 4", "Block 5", "Block 6", "Block 7"
	// for block := range list {
	// 	b := blockChain.AddBlock(list[block], len(blockChain.Blocks)+1)
	// 	logic.AddBlockInDB(db, *b)
	// }

	for i := range blockChain.Blocks {
		fmt.Printf("\nID       : %v \n", blockChain.Blocks[i].ID)
		fmt.Printf("Hash     : %v \n", blockChain.Blocks[i].Hash)
		fmt.Printf("prevHash : %v \n", blockChain.Blocks[i].PrevHash)
		fmt.Printf("Data     : %v \n", blockChain.Blocks[i].Data)
	}

	HandleFuncRequest()

}
