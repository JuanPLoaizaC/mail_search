package domain

type IEmail interface {
	IndexEmails(emails interface{}) (*IndexEmailResponse, error)
	IndexedSearch(bodyrequest IndexedSearchRequest) (*IndexedSearchResponse, error)
}

// IndexerEmailResponse is the response for the IndexEmailResponse function
type IndexEmailResponse struct {
	Code   int                        `json:"code"`
	Status []IndexEmailStatusResponse `json:"status"`
}

type IndexEmailStatusResponse struct {
	Name       string `json:"name"`
	Successful int    `json:"successful"`
	Failed     int    `json:"failed"`
}

// IndexedSearchRequest is the request for the IndexedSearch function
type IndexedSearchRequest struct {
	SearchType string                    `json:"search_type"`
	SortFields []string                  `json:"sort_fields"`
	From       int                       `json:"from"`
	MaxResults int                       `json:"max_results"`
	Query      IndexedSearchRequestQuery `json:"query"`
	// Source     []string                  `json:"_source"`
}

// IndexedSearchRequestQuery is the query for the IndexedSearch function
type IndexedSearchRequestQuery struct {
	Term string `json:"term"`
	// StartTime string `json:"start_time"`
	// EndTime   string `json:"end_time"`
}

// IndexedSearchResponse is the response for the IndexedSearch function
type IndexedSearchResponse struct {
	Took     float64 `json:"took"`
	TimedOut bool    `json:"timed_out"`
	MaxScore float64 `json:"max_score"`
	Hits     struct {
		Total struct {
			Value int `json:"value"`
		} `json:"total"`
		Hits []struct {
			Index     string                 `json:"_index"`
			Type      string                 `json:"_type"`
			ID        string                 `json:"_id"`
			Score     float64                `json:"_score"`
			Timestamp string                 `json:"@timestamp"`
			Source    map[string]interface{} `json:"_source"`
		} `json:"hits"`
	} `json:"hits"`
}

// ErrorReponse is the response of datasource when an error occurs
type ErrorReponse struct {
	Code         int    `json:"code"`
	ErrorDetail  string `json:"error-detail"`
	ErrorMessage string `json:"message"`
}
