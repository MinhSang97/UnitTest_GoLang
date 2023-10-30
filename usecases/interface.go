package usecases

import (
	"app/model"
	"context"
)

type UseCases interface {
	GetOneByID(ctx context.Context, id int) (model.Student, error)
	GetAll(ctx context.Context) ([]model.Student, error)
	InsertOne(ctx context.Context, c *model.Student) error
}
