package data

import (
	"context"
	"go-practice-homework/week04/internal/app/biz"
)

type dataRepo struct {
	dao *dao
}

func NewDataRepo() biz.DataRepo {
	return &dataRepo{}
}

func (d *dataRepo) GetData(ctx context.Context, id int64) (*biz.Data, error) {
	return &biz.Data{Id: id, Data: d.dao.Get(id)}, nil
}
