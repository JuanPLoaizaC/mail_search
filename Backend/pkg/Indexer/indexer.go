package indexer

import (
	"fmt"
	"net/http"

	"github.com/JuanPLoaizaC/mail_search/tree/main/Backend/pkg/database"
	"github.com/JuanPLoaizaC/mail_search/tree/main/Backend/pkg/service"
)

// Indexer contains the structure of the indexer service
type Indexer struct {
	httpClient     *http.Client
	indexerService *service.IndexerEmailService
}

// NewIndexer works as the constructor of the Indexer struc
func NewIndexer() *Indexer {
	indexer := &Indexer{}
	indexer.configureHttpClient()
	indexer.configureDependencies()

	return indexer
}

// configureHttpClient to communicate with zincsearch
func (i *Indexer) configureHttpClient() {
	i.httpClient = &http.Client{}
}

// configureDependencies setups dependency injection
func (i *Indexer) configureDependencies() {
	zincSearch := database.NewZincsearchClient(i.httpClient)

	indexerEmailService := service.NewIndexerService(zincSearch)
	i.indexerService = indexerEmailService
}

func (i *Indexer) Run() {
	fmt.Println("Indexer is running")
	i.indexerService.IndexEmails()
}
