package endpoint

import "github.com/hpifu/go-aliyun/internal/store"

type Service struct {
	es *store.EndpointStore
}

func NewService(root string) (*Service, error) {
	es, err := store.NewEndpointStore(root)
	if err != nil {
		return nil, err
	}
	return &Service{es: es}, nil
}
