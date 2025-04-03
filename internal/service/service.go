package service

import (
	"context"
	"fmt"

	"github.com/front-go/user/pkg/user"
)

type Service struct {
	user.UnimplementedUserServiceServer
	DbRepo DbRepo
}

func NewService(DbRepo DbRepo) *Service {
	return &Service{DbRepo: DbRepo}
}

func (s *Service) UpdateUserPersonalData(ctx context.Context, in *user.UpdateUserPersonalDataIn) (out *user.UpdateUserPersonalDataOut, err error) {
	fmt.Println("asdasdasda")
	return nil, nil
}
