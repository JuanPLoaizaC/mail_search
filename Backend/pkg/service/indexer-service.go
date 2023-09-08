package service

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"

	"github.com/JuanPLoaizaC/mail_search/tree/main/Backend/pkg/domain"
)

const emailFolderPath = "../../Files/enron_mail_20110402/maildir"

type IndexerEmailService struct {
	datasource domain.IEmail
}

func NewIndexerService(ds domain.IEmail) *IndexerEmailService {
	return &IndexerEmailService{
		datasource: ds,
	}
}

func (ies *IndexerEmailService) IndexEmails() {
	emailUsers, err := ies.getMailUsers()
	if err != nil {
		// customerror.NewCustomError(http.StatusInternalServerError, err.Error()).ErrorResponseHandling(w, r)
		return
	}

	var wg sync.WaitGroup
	for _, emailUser := range emailUsers {
		wg.Add(1)
		go ies.indexEmailByUser(emailUser, &wg)
	}
	wg.Wait()
}

func (ies *IndexerEmailService) getMailUsers() ([]string, error) {
	var mailUsers []string

	dirs, err := os.ReadDir(emailFolderPath)
	if err != nil {
		return nil, err
	}

	for _, dir := range dirs {
		mailUsers = append(mailUsers, dir.Name())
	}

	return mailUsers, nil
}

func (ies *IndexerEmailService) indexEmailByUser(userEmail string, wg *sync.WaitGroup) {
	defer wg.Done()
	emails, err := ies.processMailsByUser(userEmail)
	if err != nil {
		return
	}

	err = ies.indexEmails(emails)
	if err != nil {
		return
	}
}

func (ies *IndexerEmailService) processMailsByUser(user string) ([]domain.Email, error) {
	var emails []domain.Email
	path := emailFolderPath + "/" + user
	err := filepath.Walk(path, readEmails(&emails))
	if err != nil {
		return nil, err
	}

	return emails, nil
}

func readEmails(emails *[]domain.Email) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			email, err := processEmailFile(path)
			if err != nil {
				return nil
			}
			*emails = append(*emails, *email)
		}
		return nil
	}

}

func processEmailFile(emailFilepath string) (*domain.Email, error) {
	emailContent, err := os.ReadFile(emailFilepath)
	if err != nil {
		return nil, err
	}
	return mapStringToEmail(string(emailContent)), nil
}

func (ies *IndexerEmailService) indexEmails(records []domain.Email) error {
	response, err := ies.datasource.IndexEmails(records)
	if err != nil {
		return err
	}
	fmt.Printf("Name: " + response.Status[0].Name + ", Inserted: " + strconv.Itoa(response.Status[0].Successful))
	fmt.Println()

	return nil
}

func mapStringToEmail(emailString string) *domain.Email {
	detailsAndContent := strings.SplitN(string(emailString), "\r\n\r\n", 2)
	details := strings.Split(detailsAndContent[0], "\r\n")

	newEmail := &domain.Email{}
	for _, detail := range details {
		detailValue := strings.SplitN(detail, ": ", 2)
		switch detailValue[0] {
		case "Message-ID":
			newEmail.MessageID = detailValue[1]
		case "Date":
			newEmail.Date = detailValue[1]
		case "From":
			newEmail.From = detailValue[1]
		case "To":
			newEmail.To = detailValue[1]
		case "Subject":
			newEmail.Subject = detailValue[1]
		default:
			continue
		}
	}
	newEmail.Content = detailsAndContent[1]

	return newEmail
}
