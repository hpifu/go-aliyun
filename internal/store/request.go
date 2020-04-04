package store

import (
	"path/filepath"
)

type RequestStore struct {
	store *FileStore
}

func NewRequestStore(root string) (*RequestStore, error) {
	store, err := NewFileStore(root)
	if err != nil {
		return nil, err
	}
	return &RequestStore{
		store: store,
	}, nil
}

func (cs *RequestStore) Put(category, subCategory, filename string, request map[string]string) error {
	return cs.store.Put(filepath.Join(category, subCategory, filename), request)
}

func (cs *RequestStore) Get(category, subCategory, filename string) (map[string]string, error) {
	r := map[string]string{}
	if err := cs.store.Get(filepath.Join(category, subCategory, filename), r); err != nil {
		return nil, err
	}
	return r, nil
}

func (cs *RequestStore) Del(category, subCategory, filename string) error {
	return cs.store.Del(filepath.Join(category, subCategory, filename))
}

func (cs *RequestStore) List(category, subCategory string) ([]string, error) {
	return cs.store.List(filepath.Join(category, subCategory))
}
