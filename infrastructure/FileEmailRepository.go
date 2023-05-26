package infrastructure

import (
	"github.com/emirpasic/gods/sets/hashset"
	"github.com/goccy/go-json"
	"os"
)

const storageFile = "emails.json"

type FileEmailRepository struct {
	Emails hashset.Set
}

func NewFileEmailRepository() *FileEmailRepository {
	emails := *hashset.New()

	if _, err := os.Stat(storageFile); !os.IsNotExist(err) {
		data, _ := os.ReadFile(storageFile)

		emails.FromJSON(data)
	}

	repo := FileEmailRepository{Emails: emails}

	return &repo
}

func (repo *FileEmailRepository) AddEmail(email string) {
	if repo.Emails.Contains(email) {
		// TODO add error handling
		return
	}

	repo.Emails.Add(email)
}

func (repo *FileEmailRepository) Save() {
	data, _ := json.Marshal(repo.Emails)

	os.WriteFile(storageFile, data, 0644)
}
