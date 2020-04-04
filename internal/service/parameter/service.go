package parameter

import "github.com/hpifu/go-aliyun/internal/store"

type Service struct {
	ps *store.ParameterStore
}

func NewService(root string) (*Service, error) {
	ps, err := store.NewParameterStore(root)
	if err != nil {
		return nil, err
	}
	return &Service{ps: ps}, nil
}
