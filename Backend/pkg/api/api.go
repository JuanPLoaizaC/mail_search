package api

import (
	"fmt"
	"net/http"

	"github.com/JuanPLoaizaC/mail_search/tree/main/Backend/pkg/database"
	"github.com/JuanPLoaizaC/mail_search/tree/main/Backend/pkg/handler"
	"github.com/JuanPLoaizaC/mail_search/tree/main/Backend/pkg/service"
	"github.com/go-chi/chi/v5"
)

const (
	serverPort = ":8000"
)

// Api contains the structure of the server
type Api struct {
	Router               *chi.Mux
	httpClient           *http.Client
	indexedSearchHandler *handler.IndexedSearchHandler
}

// NewApi works as the constructor of the Api struc
func NewApi() *Api {
	server := &Api{}
	server.configureHttpClient()
	server.configureDependencies()
	server.configureRouter()

	return server
}

// configureHttpClient setups the http client which needs the sling client
func (a *Api) configureHttpClient() {
	a.httpClient = &http.Client{}
}

// configureDependencies setups dependency injection
func (a *Api) configureDependencies() {
	datasourceZincSearch := database.NewZincsearchClient(a.httpClient)

	indexedSearchService := service.NewIndexedSearchService(datasourceZincSearch)
	indexedSearchHandler := handler.NewIndexedSearchHandler(indexedSearchService)
	a.indexedSearchHandler = indexedSearchHandler
}

func (a *Api) Run() {
	fmt.Println("Application is running on host http://localhost:8000/")

	err := http.ListenAndServe(serverPort, a.Router)
	if err != nil {
		panic(err)
	}
}
