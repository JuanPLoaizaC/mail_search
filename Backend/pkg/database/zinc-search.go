package database

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	customerror "github.com/JuanPLoaizaC/mail_search/tree/main/Backend/pkg/custom-error"
	"github.com/JuanPLoaizaC/mail_search/tree/main/Backend/pkg/domain"
)

var headers = map[string]string{
	"Content-Type": "application/json",
}

const (
	defaultZincSearchHost = "http://localhost:5080"
	organizationName      = "enron"
	streamName            = "enron_emails"
)

// ZincSearchClient is the client that will communicate with the zincsearch api
type ZincSearchClient struct {
	client *http.Client
}

// NewZincsearchClient  works as the constructor of the ZincsearchClient struc
func NewZincsearchClient(client *http.Client) *ZincSearchClient {
	return &ZincSearchClient{
		client: client,
	}
}

// IndexEmails uses the zincsearch API to create indexed documents
func (zc *ZincSearchClient) IndexEmails(emailsToinsert interface{}) (*domain.IndexEmailResponse, error) {
	successResponse := &domain.IndexEmailResponse{}
	errorResponse := &domain.ErrorReponse{}
	url := fmt.Sprintf("%s/api/%s/%s/_json", defaultZincSearchHost, organizationName, streamName)

	req, err := makeRequest(url, emailsToinsert)
	if err != nil {
		return nil, err
	}

	resp, err := zc.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		err := json.NewDecoder(resp.Body).Decode(errorResponse)
		if err != nil {
			return nil, err
		}
		return nil, &customerror.CustomError{ErrorText: errorResponse.ErrorMessage}
	}

	err = json.NewDecoder(resp.Body).Decode(successResponse)
	if err != nil {
		return nil, err
	}

	return successResponse, nil
}

// IndexedSearch uses the zincsearch API to perform an indexed search for a term within the content of documents
func (zc *ZincSearchClient) IndexedSearch(termToSearch string) (*domain.IndexedSearchResponse, error) {
	successResponse := &domain.IndexedSearchResponse{}
	errorResponse := &domain.ErrorReponse{}
	url := fmt.Sprintf("%s/api/%s/_search", defaultZincSearchHost, organizationName)

	bodyRequest := domain.IndexedSearchRequest{
		Query: domain.QueryIndexedSearchRequest{
			Sql:  makeSqlQuery(termToSearch),
			From: 0,
			Size: 105,
		},
	}
	req, err := makeRequest(url, bodyRequest)
	if err != nil {
		return nil, err
	}

	resp, err := zc.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		err := json.NewDecoder(resp.Body).Decode(errorResponse)
		if err != nil {
			return nil, err
		}
		return nil, &customerror.CustomError{ErrorText: errorResponse.ErrorMessage}
	}

	err = json.NewDecoder(resp.Body).Decode(successResponse)
	if err != nil {
		return nil, err
	}

	return successResponse, nil
}

// makeRequest makes a request to the provided url with the provided POST method and body
func makeRequest(url string, body interface{}) (*http.Request, error) {
	bodyRequest := makeBodyRequest(body)

	req, err := http.NewRequest("POST", url, bodyRequest)
	if err != nil {
		return nil, err
	}
	setHeaders(req)

	return req, nil
}

func makeBodyRequest(bodyRequest interface{}) io.Reader {
	if bodyRequest == nil {
		return nil
	}
	body, err := json.Marshal(bodyRequest)
	if err != nil {
		panic(err)
	}
	return bytes.NewReader(body)
}

// setHeaders configures the headers for the request to the zincsearch api
func setHeaders(req *http.Request) {
	req.SetBasicAuth(ZO_ROOT_USER_EMAIL, ZO_ROOT_USER_PASSWORD)

	for k, v := range headers {
		req.Header.Set(k, v)
	}
}

func makeSqlQuery(term string) string {
	columns := "*"
	return fmt.Sprintf(`SELECT %s FROM %s WHERE match_all('%s')`, columns, streamName, term)
}
