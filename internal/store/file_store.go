package store

import (
	"bufio"
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
)

type FileStore struct {
	root string
}

func NewFileStore(root string) (*FileStore, error) {
	if err := os.MkdirAll(root, 0755); err != nil {
		return nil, err
	}
	return &FileStore{
		root: root,
	}, nil
}

func (cs *FileStore) Put(filename string, v interface{}) error {
	fp, err := os.OpenFile(filepath.Join(cs.root, filename), os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
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

func (cs *FileStore) Get(filename string, v interface{}) error {
	buf, err := ioutil.ReadFile(filepath.Join(cs.root, filename))
	if err != nil {
		return err
	}
	if err = json.Unmarshal(buf, v); err != nil {
		return err
	}
	return nil
}

func (cs *FileStore) Del(filename string) error {
	return os.RemoveAll(filepath.Join(cs.root, filename))
}

func (cs *FileStore) List() ([]string, error) {
	var fns []string
	infos, err := ioutil.ReadDir(cs.root)
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
