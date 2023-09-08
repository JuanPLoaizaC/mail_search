package service

import (
	"encoding/json"

	"github.com/JuanPLoaizaC/mail_search/tree/main/Backend/pkg/domain"
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
	response, err := iss.datasource.IndexedSearch(term)
	if err != nil {
		return nil, err
	}

	return mapResponseToEmails(response), nil
}

func mapResponseToEmails(response *domain.IndexedSearchResponse) []domain.Email {
	var emails []domain.Email

	for _, hit := range response.Hits {
		var email domain.Email
		contetEmialBytes, _ := json.Marshal(hit)
		err := json.Unmarshal(contetEmialBytes, &email)
		if err != nil {
			continue
		}
		emails = append(emails, email)
	}

	return emails
}
