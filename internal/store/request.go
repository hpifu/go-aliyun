package store

import (
	"bufio"
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
)

type RequestStore struct {
	root string
}

func NewRequestStore(root string) (*RequestStore, error) {
	if err := os.MkdirAll(root, 0755); err != nil {
		return nil, err
	}
	return &RequestStore{
		root: root,
	}, nil
}

func (cs *RequestStore) Put(category, subCategory, filename string, request map[string]string) error {
	fp, err := os.OpenFile(filepath.Join(cs.root, category, subCategory, filename), os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	defer fp.Close()
	if err != nil {
		return err
	}
	w := bufio.NewWriter(fp)
	buf, err := json.Marshal(request)
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

func (cs *RequestStore) Get(category, subCategory, filename string) (map[string]string, error) {
	buf, err := ioutil.ReadFile(filepath.Join(cs.root, category, subCategory, filename))
	if err != nil {
		return nil, err
	}
	c := map[string]string{}
	if err = json.Unmarshal(buf, c); err != nil {
		return nil, err
	}
	return c, nil
}

func (cs *RequestStore) Del(category, subCategory, filename string) error {
	return os.RemoveAll(filepath.Join(cs.root, category, subCategory, filename))
}

func (cs *RequestStore) List(category, subCategory string) ([]string, error) {
	var fns []string
	infos, err := ioutil.ReadDir(filepath.Join(cs.root, category, subCategory))
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
