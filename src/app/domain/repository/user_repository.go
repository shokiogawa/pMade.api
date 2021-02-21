package repository

import (
	"context"
	"pMade.api/src/app/domain/model"
)

type UserRepository interface {
	GetUser(ctx context.Context) ([]*model.User, error)
}
