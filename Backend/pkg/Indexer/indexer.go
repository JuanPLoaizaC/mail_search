package indexer

import (
	"fmt"
	"net/http"

	"github.com/JuanPLoaizaC/mail_search/tree/main/Backend/pkg/database"
	"github.com/JuanPLoaizaC/mail_search/tree/main/Backend/pkg/service"
)

type Indexer struct {
	httpClient     *http.Client
	indexerService *service.IndexerEmailService
}

func NewIndexer() *Indexer {
	indexer := &Indexer{}
	indexer.configureHttpClient()
	indexer.configureDependencies()

	return indexer
}

func (i *Indexer) configureHttpClient() {
	i.httpClient = &http.Client{}
}

func (i *Indexer) configureDependencies() {
	zincSearch := database.NewZincsearchClient(i.httpClient)

	indexerEmailService := service.NewIndexerService(zincSearch)
	i.indexerService = indexerEmailService
}

func (i *Indexer) Run() {
	fmt.Println("Indexer is running")
	i.indexerService.IndexEmails()
}
