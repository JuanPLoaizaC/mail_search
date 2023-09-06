package database

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	customerror "github.com/JuanPLoaizaC/mail_search_truora/tree/main/Backend/pkg/database/custom-error"
	"github.com/JuanPLoaizaC/mail_search_truora/tree/main/Backend/pkg/domain"
)

var headers = map[string]string{
	"Content-Type": "application/json",
	// "Authorization": "Basic",
}

const (
	defaultZincSearchHost = "http://localhost:5080"
	organizationName      = "enron"
	streamName            = "enron_emails"
)

type ZincSearchClient struct {
	client *http.Client
}

func NewZincsearchClient(client *http.Client) *ZincSearchClient {
	return &ZincSearchClient{
		client: client,
	}
}

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
		return nil, &customerror.CustomDbError{ErrorText: errorResponse.ErrorMessage}
	}

	err = json.NewDecoder(resp.Body).Decode(successResponse)
	if err != nil {
		return nil, err
	}

	return successResponse, nil
}

func (zc *ZincSearchClient) IndexedSearch(bodyrequest domain.IndexedSearchRequest) (*domain.IndexedSearchResponse, error) {
	// TODO
	return nil, nil
}

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

func setHeaders(req *http.Request) {
	user, pass := getAuthentication()
	req.SetBasicAuth(user, pass)

	for k, v := range headers {
		req.Header.Set(k, v)
	}
}

func getAuthentication() (string, string) {
	// user := os.Getenv("ZO_ROOT_USER_EMAIL")
	// pass := os.Getenv("ZINC_FIRST_ADMIN_PASSWORD")

	// if user == "" || pass == "" {
	// 	panic("ZO_ROOT_USER_EMAIL and  ZO_ROOT_USER_PASSWORD must be set")
	// }
	return "juanpablo9730@gmail.com", "Bur11Caldas"
	// return base64.StdEncoding.EncodeToString([]byte(user + ":" + pass))
}
