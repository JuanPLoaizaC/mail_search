package service

import (
	"encoding/json"

	"github.com/JuanPLoaizaC/mail_search/tree/main/Backend/pkg/domain"
)

// IndexedSearchService is the struc that will communicate with the datasource
type IndexedSearchService struct {
	datasource domain.IEmail
}

// NewIndexedSearchService works as the constructor of the IndexedSearchService struc
func NewIndexedSearchService(ds domain.IEmail) *IndexedSearchService {
	return &IndexedSearchService{
		datasource: ds,
	}
}

// IndexedSearch performs the search for the term in the configured datasource
func (iss *IndexedSearchService) IndexedSearch(term string) ([]domain.Email, error) {
	response, err := iss.datasource.IndexedSearch(term)
	if err != nil {
		return nil, err
	}

	return mapResponseToEmails(response), nil
}

// mapResponseToEmails maps the IndexedSearchResponseResponse response structure to Emails
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
