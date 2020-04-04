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

func (rs *RequestStore) Put(category, subCategory, filename string, request map[string]string) error {
	return rs.store.Put(filepath.Join(category, subCategory, filename), request)
}

func (rs *RequestStore) Get(category, subCategory, filename string) (map[string]string, error) {
	r := map[string]string{}
	if err := rs.store.Get(filepath.Join(category, subCategory, filename), r); err != nil {
		return nil, err
	}
	return r, nil
}

func (rs *RequestStore) Del(category, subCategory, filename string) error {
	return rs.store.Del(filepath.Join(category, subCategory, filename))
}

func (rs *RequestStore) List(category, subCategory string) ([]string, error) {
	return rs.store.List(filepath.Join(category, subCategory))
}
