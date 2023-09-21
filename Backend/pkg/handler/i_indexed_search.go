package handler

import "github.com/JuanPLoaizaC/mail_search/tree/main/Backend/pkg/domain"

//  IIndexedSearch is the contract to be implemented by the service that performs the indexed search
type IIndexedSearch interface {
	IndexedSearch(term string) ([]domain.Email, error)
}
