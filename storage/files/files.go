package files

import (
	"bot/lib/e"
	"bot/storage"
	"os"
	"path/filepath"
)

type Storage struct {
	basePath string
}

func New(basePath string) Storage {
	return Storage{basePath: basePath}
}

func (s Storage) Save(page *storage.Page) (err error) {
	defer func() {err = e.Wrap("can't save", err) }()

	filePath := filepath.Join(s.basePath, page.UserName)
	if err := os.MkdirAll(filePath, 0774); err!= nil {
		return err
	}
	fileName, err := filName(page)
	if err != nil {
		return err
	}
	filePath = filepath.Join(filePath, fileName)
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()
}

func filName(p *storage.Page) (string, error) {
	return p.Hash()
}