package store

import (
	"bufio"
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type FileStore struct {
	root string
}

func NewFileStore(root string) (*FileStore, error) {
	if err := os.MkdirAll(root, 0755); err != nil {
		return nil, err
	}
	abs, err := filepath.Abs(root)
	if err != nil {
		return nil, err
	}
	return &FileStore{
		root: abs,
	}, nil
}

func (f *FileStore) validPath(filename string) (string, error) {
	path, err := filepath.Abs(filepath.Join(f.root, filename))
	if err != nil {
		return "", err
	}
	if !strings.HasPrefix(path, f.root) {
		return "", errors.New("operation forbidden")
	}
	return path, nil
}

func (f *FileStore) Put(filename string, v interface{}) error {
	path, err := f.validPath(filename)
	if err != nil {
		return err
	}

	if _, err := os.Stat(filepath.Dir(path)); err != nil && os.IsNotExist(err) {
		err = os.MkdirAll(filepath.Dir(path), 0755)
		if err != nil {
			return err
		}
	}

	fp, err := os.OpenFile(path, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	defer fp.Close()
	if err != nil {
		return err
	}
	w := bufio.NewWriter(fp)
	buf, err := json.Marshal(v)
	if err != nil {
		return err
	}
	_, err = w.Write(buf)
	if err != nil {
		return err
	}
	if err = w.Flush(); err != nil {
		return err
	}
	return nil
}

func (f *FileStore) Get(filename string, v interface{}) error {
	path, err := f.validPath(filename)
	if err != nil {
		return err
	}
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	if err = json.Unmarshal(buf, v); err != nil {
		return err
	}
	return nil
}

func (f *FileStore) Del(filename string) error {
	path, err := f.validPath(filename)
	if err != nil {
		return err
	}
	return os.Remove(path)
}

func (f *FileStore) List(subDir string) ([]string, error) {
	path, err := f.validPath(subDir)
	if err != nil {
		return nil, err
	}

	var fns []string
	infos, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}
	for _, info := range infos {
		if !info.IsDir() {
			fns = append(fns, info.Name())
		}
	}
	return fns, nil
}
