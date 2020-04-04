package credential

import "github.com/hpifu/go-aliyun/internal/store"

type Service struct {
	cs *store.CredentialStore
}

func NewService(root string) (*Service, error) {
	cs, err := store.NewCredentialStore(root)
	if err != nil {
		return nil, err
	}
	return &Service{cs: cs}, nil
}
