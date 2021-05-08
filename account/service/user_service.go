package service

import (
	"context"
	"log"
	"memrizr/apperrors"
	"memrizr/model"
	"memrizr/util"

	"github.com/google/uuid"
)

type UserService struct {
	UserRepository model.UserRepository
}

func NewUserService(r model.UserRepository) model.UserService {
	return &UserService{
		UserRepository: r,
	}
}

func (s *UserService) Get(ctx context.Context, uid uuid.UUID) (*model.User, error) {
	return s.UserRepository.FindByID(ctx, uid)
}

func (s *UserService) Signup(ctx context.Context, u *model.User) error {
	pw, err := util.HashPassword(u.Password)
	if err != nil {
		log.Printf("Unable to signup user for email: %v\n", u.Email)
		return apperrors.NewInternal()
	}
	u.Password = pw
	return s.UserRepository.Create(ctx, u)
}
