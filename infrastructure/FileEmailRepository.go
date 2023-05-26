package infrastructure

import (
	"github.com/emirpasic/gods/sets/hashset"
	"os"
)

const storageFile = "emails.json"

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

func (repo *FileEmailRepository) AddEmail(email string) {
	if repo.Emails.Contains(email) {
		// TODO add error handling
		return
	}

	repo.Emails.Add(email)
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
