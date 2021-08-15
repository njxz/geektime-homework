package service

import (
	"context"
	v1 "geektime/api/mydemo/v1"
	"geektime/internal/data"
)

func NewDemo(data2 *data.Data) *Mydemo {
	return &Mydemo{
		dao: data2,
	}

}

type Mydemo struct {
	v1.UnimplementedHelloServer
	dao *data.Data
}

func (d Mydemo) CheckName(ctx context.Context, in *v1.IdRequest) (*v1.IdResponse, error) {
	name, err := d.dao.SelectName(in.GetId())
	if err != nil {
		return nil, err
	}
	return &v1.IdResponse{Name: name}, nil
}

func (d Mydemo) InsertName(ctx context.Context, in *v1.InsertReq) (*v1.InsertRes, error) {
	err := d.dao.CreateId(in.GetId(), in.GetName())
	if err != nil {
		return nil, err
	}
	return &v1.InsertRes{Status: true}, nil
}
