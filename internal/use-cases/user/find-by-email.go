package user

import (
	"context"
	"errors"
	"strings"
	"time-turner-api/internal/domain/user"
)

type FindByEmailUserInput struct {
	Email   string 	`json:"email"`
}

type FindByEmailUserOutput struct {
	Id        string 	`json:"id"`
	Name      string 	`json:"name"`
	Email     string 	`json:"email"`
}

type FindByEmailUserUseCase struct {
	userRepo user.UserRepository
}

func NewFindByEmailUserUseCase (repo user.UserRepository) *FindByEmailUserUseCase {
	return &FindByEmailUserUseCase{
		userRepo: repo,
	}
}

func (uf *FindByEmailUserUseCase) Execute(ctx context.Context, input FindByEmailUserInput) (*FindByEmailUserOutput, error) {
	email := strings.TrimSpace(strings.ToLower(input.Email))

	if email == "" || !strings.Contains(email,"@") {
		return nil, errors.New("Email invalido")
	}

	found, err := uf.userRepo.FindByEmail(ctx, email)

	if err != nil {
		return nil, err
	}

	if found == nil {
		return nil, errors.New("Usuario nao encontrado")
	}

	return &FindByEmailUserOutput{
		Id: found.Id,
		Name: found.Name,
		Email: found.Email,
	}, nil
}