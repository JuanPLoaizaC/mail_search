package main

import (
	"fmt"

	"github.com/JuanPLoaizaC/mail_search/tree/main/Backend/cmd/profiling"
	indexer "github.com/JuanPLoaizaC/mail_search/tree/main/Backend/pkg/Indexer"
	"github.com/JuanPLoaizaC/mail_search/tree/main/Backend/pkg/api"
)

func main() {
	var mode int

	fmt.Println("modes: 1 -> api. 2 -> indexer. 3 -> profiling.")
	fmt.Scanln(&mode)

	if mode == 1 {
		api.NewApi().Run()
	} else if mode == 2 {
		indexer.NewIndexer().Run()
	} else if mode == 3 {
		profiling.RunProfiling()
	} else {
		fmt.Println("Incorrect data, the options are 1 -> api. 2 -> indexer.")
	}
}
