package biz

import "context"

type (
	Data struct {
		Id int64
		Data string
	}

	DataRepo interface {
		GetData(ctx context.Context, id int64) (*Data, error)
	}

	DataUsecase struct {
		repo DataRepo
	}
)

func NewDataUsecase(repo DataRepo) *DataUsecase {
	return &DataUsecase{repo: repo}
}

func (u *DataUsecase) Get(ctx context.Context, id int64) (*Data, error) {
	return u.repo.GetData(ctx, id)
}