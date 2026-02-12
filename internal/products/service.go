package products

import "context"

type Service interface {
	ListProduct(ctx context.Context) error
}

type svc struct {
}

func NewService() Service {
	return &svc{}
}

func (s *svc) ListProduct(ctx context.Context) error {
	return nil
}
