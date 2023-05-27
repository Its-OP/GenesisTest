package infrastructure

import (
	"btcRate/domain"
	"fmt"
	"github.com/emirpasic/gods/sets/hashset"
	"os"
)

const storageFile = "data/emails.json"

type FileEmailRepository struct {
	Emails hashset.Set
}

func NewFileEmailRepository() *FileEmailRepository {
	emails := *hashset.New()

	if fileExists() {
		data, _ := os.ReadFile(storageFile)
		emails.FromJSON(data)
	}

	repo := FileEmailRepository{Emails: emails}
	return &repo
}

func (repo *FileEmailRepository) AddEmail(email string) error {
	if repo.Emails.Contains(email) {
		return &domain.DataConsistencyError{Message: fmt.Sprintf("Email address '%s' is already present in the database", email)}
	}

	repo.Emails.Add(email)
	return nil
}

func (repo *FileEmailRepository) GetAll() []string {
	if !fileExists() {
		return []string{}
	}

	values := repo.Emails.Values()

	emailsArray := make([]string, len(values))
	for i, value := range values {
		emailsArray[i] = value.(string)
	}

	return emailsArray
}

func (repo *FileEmailRepository) Save() {
	data, _ := repo.Emails.MarshalJSON()

	if !fileExists() {
		os.Create(storageFile)
	}

	os.WriteFile(storageFile, data, 0644)
}

func fileExists() bool {
	info, err := os.Stat(storageFile)
	if os.IsNotExist(err) {
		return false
	}

	return !info.IsDir()
}
