package ivy

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
)

//FileSystemCache is file system cache
type FileSystemCache struct {
	root string
}

//NewFileSystemCache create file system cache
func NewFileSystemCache(root string) *FileSystemCache {
	if err := os.MkdirAll(root, 0755); err != nil {
		panic(err)
	}

	return &FileSystemCache{root}
}

//Save file into file system
func (fs *FileSystemCache) Save(bucket, filename string, params *Params, data []byte) error {
	dir, filePath := filepath.Split(filename)
	filename = path.Join(fs.root, bucket, dir, params.String()+filePath)
	log.Printf("[Ivy] Save %s ", filename)
	dir = path.Join(fs.root, bucket, dir)

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

//Load file from file system return not found if file doesn't exist
func (fs *FileSystemCache) Load(bucket, filename string, params *Params) (*bytes.Buffer, error) {
	dir, filePath := filepath.Split(filename)
	filename = path.Join(fs.root, bucket, dir, params.String()+filePath)
	log.Printf("[Ivy] Load %s ", filename)
	file, err := os.Open(filename)
	if os.IsNotExist(err) {
		return nil, ErrNotFound
	}
	if err != nil {
		return nil, err
	}
	defer file.Close()

	buffer := bytes.NewBuffer(nil)
	buffer.ReadFrom(file)

	return buffer, nil
}

//Delete delete individual ache file
func (fs *FileSystemCache) Delete(bucket, filename string, params *Params) error {
	dir, filePath := filepath.Split(filename)
	filename = path.Join(fs.root, bucket, dir, params.String()+filePath)
	ext := filepath.Ext(filename)
	if ext == "" {
		return fmt.Errorf("this is not file")
	}
	return os.Remove(filename)
}

//Flush remove all cache file specific bucket
func (fs *FileSystemCache) Flush(bucket string) error {
	return os.RemoveAll(path.Join(fs.root, bucket))
}
