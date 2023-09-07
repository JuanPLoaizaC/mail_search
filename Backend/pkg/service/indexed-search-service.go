package service

import (
	"github.com/JuanPLoaizaC/mail_search_truora/tree/main/Backend/pkg/domain"
)

type IndexedSearchService struct {
	datasource domain.IEmail
}

func NewIndexedSearchService(ds domain.IEmail) *IndexedSearchService {
	return &IndexedSearchService{
		datasource: ds,
	}
}

func (iss *IndexedSearchService) IndexedSearch(term string) ([]domain.Email, error) {
	return nil, nil
}

func mapResponseToEmails(response *domain.IndexedSearchResponse) []domain.Email {
	var emails []domain.Email

	return emails
}
