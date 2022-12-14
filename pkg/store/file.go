package store

import (
	"encoding/csv"
	"os"
)

type Store interface {
	Read(data interface{}) error
}

type fileStore struct {
	FilePath string
}

func New(filename string) Store {
	return &fileStore{filename}
}

func (f *fileStore) Read(data interface{}) error {
	file, err := os.Open(f.FilePath)
	if err != nil {
		return err
	}
	defer file.Close()

	rdr := csv.NewReader(file)
	data, err = rdr.ReadAll()
	if err != nil {
		return err
	}
	return nil
}
