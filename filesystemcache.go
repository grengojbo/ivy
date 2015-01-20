package fileproxy

import (
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"
)

type FileSystemCache struct {
	root string
}

func NewFileSystemCache(root string) *FileSystemCache {
	if err := os.MkdirAll(root, 0755); err != nil {
		panic(err)
	}

	return &FileSystemCache{root}
}

func (fs *FileSystemCache) Save(filename, paramsStr string, data []byte) error {
	dir, filePath := filepath.Split(filename)
	filename = path.Join(fs.root, dir, paramsStr+filePath)
	dir = path.Join(fs.root, dir)

	_, err := os.Open(dir)
	if os.IsNotExist(err) {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return err
		}
	}

	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_EXCL|os.O_SYNC, 0755)
	defer file.Close()
	if os.IsNotExist(err) || err == nil {
		if _, err := file.Write(data); err != nil {
			return err
		}
	}

	return err
}

func (fs *FileSystemCache) Load(filename, paramsStr string) (io.Reader, error) {
	dir, filePath := filepath.Split(filename)
	filename = path.Join(fs.root, dir, paramsStr+filePath)

	file, err := os.Open(filename)
	if os.IsNotExist(err) {
		return nil, ErrNotFound
	}
	if err != nil {
		return nil, err
	}

	return file, nil
}

func (fs *FileSystemCache) Delete(filename string) error {
	filename = path.Join(fs.root, filename)
	ext := filepath.Ext(filename)
	if ext == "" {
		return fmt.Errorf("this is not file")
	}
	return os.Remove(filename)
}

func (fs *FileSystemCache) Flush() error {
	return os.RemoveAll(fs.root)
}
