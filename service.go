package api_utils

import "context"

type IDefaultService interface {
	Process(ctx context.Context, req interface{}) interface{}
}

type DefaultService struct {
}

func (ds *DefaultService) Process(ctx context.Context, req interface{}) interface{} {
	return "hello"
}
