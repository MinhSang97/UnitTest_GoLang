package usecases

import (
	"app/model"
	"context"
)

type UseOfCases struct {
	repo UseCases
}

func NewStudentUseCase(repo UseCases) *UseOfCases {
	return &UseOfCases{
		repo: repo,
	}
}

func (uc *UseOfCases) GetStudentByID(ctx context.Context, id int) (model.Student, error) {
	return uc.repo.GetOneByID(ctx, id)
}

func (uc *UseOfCases) GetAllStudents(ctx context.Context) ([]model.Student, error) {
	return uc.repo.GetAll(ctx)
}

func (uc *UseOfCases) CreateStudent(ctx context.Context, student *model.Student) error {
	return uc.repo.InsertOne(ctx, student)
}
