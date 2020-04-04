package store

import (
	"bufio"
	"errors"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

type EndpointStore struct {
	root string
}

func NewEndpointStore(root string) (*EndpointStore, error) {
	if err := os.MkdirAll(root, 0755); err != nil {
		return nil, err
	}
	abs, err := filepath.Abs(root)
	if err != nil {
		return nil, err
	}
	return &EndpointStore{
		root: abs,
	}, nil
}

func (es *EndpointStore) Put(category string, endpoint string) error {
	path, err := es.validPath(category)
	if err != nil {
		return err
	}

	endpoints, err := es.List(category)
	if err != nil {
		return err
	}
	endpoint = strings.TrimSpace(endpoint)
	for _, e := range endpoints {
		if e == endpoint {
			return nil
		}
	}
	endpoints = append(endpoints, endpoint)
	sort.Strings(endpoints)

	fp, err := os.OpenFile(path, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer fp.Close()
	w := bufio.NewWriter(fp)
	for _, e := range endpoints {
		_, _ = w.WriteString(e + "\n")
	}
	if err := w.Flush(); err != nil {
		return err
	}
	return nil
}

func (es *EndpointStore) List(category string) ([]string, error) {
	path, err := es.validPath(category)
	if err != nil {
		return nil, err
	}
	if _, err := os.Stat(path); err != nil {
		return nil, nil
	}

	fp, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer fp.Close()
	r := bufio.NewReader(fp)
	var endpoints []string
	for {
		line, err := r.ReadString('\n')
		line = strings.TrimSpace(line)
		if line != "" {
			endpoints = append(endpoints, line)
		}
		if err != nil {
			break
		}
	}

	return endpoints, nil
}

func (es *EndpointStore) Del(category string, endpoint string) error {
	path, err := es.validPath(category)
	if err != nil {
		return err
	}

	endpoints, err := es.List(category)
	if err != nil {
		return err
	}
	endpoint = strings.TrimSpace(endpoint)
	var endpointsCopy []string
	for _, e := range endpoints {
		if e != endpoint {
			endpointsCopy = append(endpointsCopy, e)
		}
	}
	if len(endpointsCopy) == len(endpoint) {
		return nil
	}

	fp, err := os.OpenFile(path, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer fp.Close()
	w := bufio.NewWriter(fp)
	for _, e := range endpointsCopy {
		_, _ = w.WriteString(e + "\n")
	}
	if err := w.Flush(); err != nil {
		return err
	}
	return nil
}

func (f *EndpointStore) validPath(filename string) (string, error) {
	path, err := filepath.Abs(filepath.Join(f.root, filename))
	if err != nil {
		return "", err
	}
	if !strings.HasPrefix(path, f.root) {
		return "", errors.New("operation forbidden")
	}
	return path, nil
}
