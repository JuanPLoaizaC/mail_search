package main

import (
	"fmt"

	indexer "github.com/JuanPLoaizaC/mail_search_truora/tree/main/Backend/pkg/Indexer"
)

func main() {
	var mode int

	fmt.Println("modes: 1 -> api. 2 -> indexer.")
	fmt.Scanln(&mode)

	if mode == 1 {
		fmt.Println("TODO")
	} else if mode == 2 {
		indexer.NewIndexer().RunIndexer()
	} else {
		fmt.Println("Incorrect data, the options are 1 -> api. 2 -> indexer.")
	}
}
