package handler

import "github.com/JuanPLoaizaC/mail_search/tree/main/Backend/pkg/domain"

type IIndexedSearch interface {
	IndexedSearch(term string) ([]domain.Email, error)
}
