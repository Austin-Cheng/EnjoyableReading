package tag

import (
	"context"
)

type UseCase interface {
	Create(ctx context.Context, tag *CreateTagReq) error
}

type CreateTagReq struct {
	Name string
}
