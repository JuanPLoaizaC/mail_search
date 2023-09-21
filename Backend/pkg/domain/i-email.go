package domain

// IEmail is the contract that must implement the data source.
type IEmail interface {
	IndexEmails(emails interface{}) (*IndexEmailResponse, error)
	IndexedSearch(term string) (*IndexedSearchResponse, error)
}

// IndexEmailResponse is the response for the IndexEmailResponse function
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
	// Aggs  struct{}               `json:"aggs"`
	Query QueryIndexedSearchRequest `json:"query"`
}

type QueryIndexedSearchRequest struct {
	// EndTime   int    `json:"end_time"`
	From int    `json:"from"`
	Size int    `json:"size"`
	Sql  string `json:"sql"`
	// StartTime int    `json:"start_time"`
}

// IndexedSearchResponse is the response for the IndexedSearch function
type IndexedSearchResponse struct {
	Took       int `json:"took"`
	TookDetail struct {
		Total            int `json:"total"`
		WaitQueue        int `json:"wait_queue"`
		ClusterTotal     int `json:"cluster_total"`
		ClusterWaitQueue int `json:"cluster_wait_queue"`
	} `json:"took_detail"`
	Hits     []HitIndexedSearchResponse `json:"hits"`
	Total    int                        `json:"total"`
	From     int                        `json:"from"`
	Size     int                        `json:"size"`
	ScanSize int                        `json:"scan_size"`
}

type HitIndexedSearchResponse struct {
	Timestamp int64  `json:"_timestamp"`
	Content   string `json:"content"`
	Date      string `json:"date"`
	From      string `json:"from"`
	MessageID string `json:"message_id"`
	Subject   string `json:"subject"`
	To        string `json:"to"`
}

// ErrorReponse is the response of datasource when an error occurs
type ErrorReponse struct {
	Code         int    `json:"code"`
	ErrorDetail  string `json:"error-detail"`
	ErrorMessage string `json:"message"`
}
