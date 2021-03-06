package ivy

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"testing"
)

func TestFileCacheSave(t *testing.T) {
	fs := NewFileSystemCache("testcache")
	bucket := "bucket"
	pathFile := "subfolder/test.txt"
	paramsStr := ""
	dir, filePath := filepath.Split(pathFile)
	fullPath := path.Join(fs.root, bucket, dir, paramsStr+filePath)
	defer os.RemoveAll("testcache")

	err := fs.Save(bucket, pathFile, NewParams(), []byte("TESTCACHE"))
	assert.NoError(t, err)

	byteString, err := ioutil.ReadFile(fullPath)
	assert.NoError(t, err)
	assert.Equal(t, "TESTCACHE", string(byteString))
}

func TestFileCacheSaveExist(t *testing.T) {
	bucket := "bucket"
	fs := NewFileSystemCache("testcache")
	defer os.RemoveAll("testcache")

	err := fs.Save(bucket, "test.txt", NewParams(), []byte("TESTCACHE"))
	err = fs.Save(bucket, "test.txt", NewParams(), []byte("TESTCACHE"))
	assert.Error(t, err)
}

func TestFileCacheLoad(t *testing.T) {
	bucket := "bucket"
	fs := NewFileSystemCache("testcache")
	pathFile := "subfolder/test.txt"
	defer os.RemoveAll("testcache")

	err := fs.Save(bucket, pathFile, NewParams(), []byte("TESTCACHE"))
	assert.NoError(t, err)

	reader, err := fs.Load(bucket, pathFile, NewParams())

	assert.NoError(t, err)

	data := bytes.NewBuffer(nil)
	data.ReadFrom(reader)
	assert.Equal(t, "TESTCACHE", data.String())
}

func TestFileCacheLoadNotExist(t *testing.T) {
	bucket := "bucket"
	fs := NewFileSystemCache("testcache")
	reader, err := fs.Load(bucket, "test.txt", NewParams())
	assert.Error(t, err)
	assert.Nil(t, reader)
}

func TestFileCacheDelete(t *testing.T) {
	bucket := "bucket"
	fs := NewFileSystemCache("testcache")
	defer os.RemoveAll("testcache")

	params, err := ParseParams("r_100x200")
	assert.NoError(t, err)

	err = fs.Save(bucket, "test.txt", params, []byte("TESTCACHE"))
	assert.NoError(t, err)

	err = fs.Delete(bucket, "test.txt", params)
	assert.NoError(t, err)

	_, err = os.Open(path.Join(fs.root, bucket, "100_200_0_0_-1test.txt"))
	assert.True(t, os.IsNotExist(err))
}

func TestFileCacheFlush(t *testing.T) {
	bucket := "bucket"
	fs := NewFileSystemCache("testcache")
	defer os.RemoveAll("testcache")

	err := fs.Save(bucket, "test.txt", NewParams(), []byte("TESTCACHE"))
	assert.NoError(t, err)

	err = fs.Save(bucket, "test2.txt", NewParams(), []byte("TESTCACHE"))
	assert.NoError(t, err)

	err = fs.Flush(bucket)
	assert.NoError(t, err)

	_, err = os.Open(path.Join(fs.root, bucket))
	assert.True(t, os.IsNotExist(err))
}
