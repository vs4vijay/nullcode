package nullcode

import (
	"context"
)

// "github.com/google/uuid"

type Service interface {
	Add(ctx context.Context) (string, error)
	// Get(ctx context.Context) (string, error)
	// GetAll(ctx context.Context) (string, error)
}

type CodeService struct{}

func NewCodeService() Service {
	return CodeService{}
}

func (CodeService) Add(ctx context.Context) (string, error) {
	return "ok", nil
}

// func (CodeService) Get(ctx context.Context) (string, error) {
// 	return "ok", nil
// }
//
// func (CodeService) GetAll(ctx context.Context) (string, error) {
// 	return "ok", nil
// }
