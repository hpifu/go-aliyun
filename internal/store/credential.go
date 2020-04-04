package store

type Credential struct {
	AccessKeyID     string `json:"accessKeyID"`
	AccessKeySecret string `json:"accessKeySecret"`
}

type CredentialStore struct {
	store *FileStore
}

func NewCredentialStore(root string) (*CredentialStore, error) {
	store, err := NewFileStore(root)
	if err != nil {
		return nil, err
	}
	return &CredentialStore{
		store: store,
	}, nil
}

func (cs *CredentialStore) Put(filename, accessKeyID, accessKeySecret string) error {
	return cs.store.Put(filename, &Credential{
		AccessKeyID:     accessKeyID,
		AccessKeySecret: accessKeySecret,
	})
}

func (cs *CredentialStore) Get(filename string) (*Credential, error) {
	c := &Credential{}
	if err := cs.store.Get(filename, c); err != nil {
		return nil, err
	}
	return c, nil
}

func (cs *CredentialStore) Del(filename string) error {
	return cs.store.Del(filename)
}

func (cs *CredentialStore) List() ([]string, error) {
	return cs.store.List("")
}
