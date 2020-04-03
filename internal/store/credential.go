package store

import (
	"bufio"
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
)

type Credential struct {
	AccessKeyID     string `json:"accessKeyID"`
	AccessKeySecret string `json:"accessKeySecret"`
}

type CredentialStore struct {
	directory string
}

func NewCredentialStore(directory string) (*CredentialStore, error) {
	if err := os.MkdirAll(directory, 0755); err != nil {
		return nil, err
	}
	return &CredentialStore{
		directory: directory,
	}, nil
}

func (cs *CredentialStore) Put(filename, accessKeyID, accessKeySecret string) error {
	fp, err := os.OpenFile(filepath.Join(cs.directory, filename), os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	defer fp.Close()
	if err != nil {
		return err
	}
	w := bufio.NewWriter(fp)
	buf, err := json.Marshal(&Credential{
		AccessKeyID:     accessKeyID,
		AccessKeySecret: accessKeySecret,
	})
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

func (cs *CredentialStore) Get(filename string) (*Credential, error) {
	buf, err := ioutil.ReadFile(filepath.Join(cs.directory, filename))
	if err != nil {
		return nil, err
	}
	c := &Credential{}
	if err = json.Unmarshal(buf, c); err != nil {
		return nil, err
	}
	return c, nil
}

func (cs *CredentialStore) Del(filename string) error {
	return os.RemoveAll(filepath.Join(cs.directory, filename))
}

func (cs *CredentialStore) List() ([]string, error) {
	var fns []string
	infos, err := ioutil.ReadDir(cs.directory)
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
