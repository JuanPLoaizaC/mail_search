package handler

import (
	"net/http"

	customerror "github.com/JuanPLoaizaC/mail_search_truora/tree/main/Backend/pkg/custom-error"
	"github.com/JuanPLoaizaC/mail_search_truora/tree/main/Backend/pkg/domain"
	"github.com/go-chi/render"
)

type IndexedSearchHandler struct {
	indexedSearchService IIndexedSearch
}

func NewIndexedSearchHandler(iss IIndexedSearch) *IndexedSearchHandler {
	return &IndexedSearchHandler{
		indexedSearchService: iss,
	}
}

type SearchTermInEmailsResponse struct {
	Emails []domain.Email `json:"emails"`
}

func (ish *IndexedSearchHandler) SearchTermInEmails(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	term := query.Get("term")

	if len(term) > 150 || len(term) < 3 {
		errMessage := "term invalid. Length must be between 3 and 150"
		customerror.NewCustomHttpError(http.StatusBadRequest, errMessage).ErrorResponseHandling(w, r)
		return
	}

	emails, err := ish.indexedSearchService.IndexedSearch(term)
	if err != nil {
		customerror.NewCustomHttpError(http.StatusInternalServerError, err.Error()).ErrorResponseHandling(w, r)
		return
	}
	response := &SearchTermInEmailsResponse{
		Emails: emails,
	}
	render.JSON(w, r, response)
}
