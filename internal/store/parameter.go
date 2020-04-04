package store

import (
	"path/filepath"
)

type ParameterStore struct {
	store *FileStore
}

func NewParameterStore(root string) (*ParameterStore, error) {
	store, err := NewFileStore(root)
	if err != nil {
		return nil, err
	}
	return &ParameterStore{
		store: store,
	}, nil
}

func (ps *ParameterStore) Put(category, subCategory, filename string, request map[string]string) error {
	return ps.store.Put(filepath.Join(category, subCategory, filename), request)
}

func (ps *ParameterStore) Get(category, subCategory, filename string) (map[string]string, error) {
	r := map[string]string{}
	if err := ps.store.Get(filepath.Join(category, subCategory, filename), r); err != nil {
		return nil, err
	}
	return r, nil
}

func (ps *ParameterStore) Del(category, subCategory, filename string) error {
	return ps.store.Del(filepath.Join(category, subCategory, filename))
}

func (ps *ParameterStore) List(category, subCategory string) ([]string, error) {
	return ps.store.List(filepath.Join(category, subCategory))
}
