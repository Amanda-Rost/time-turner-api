package user

import (
	"context"
	"errors"
	"strings"
	"time-turner-api/internal/domain/user"
)

type FindByIdUserInput struct {
	Id  string 	`json:"id"`
}

type FindByIdUserOutput struct {
	Id        string 	`json:"id"`
	Name      string 	`json:"name"`
	Email     string 	`json:"email"`
}

type FindByIdUserUseCase struct {
	userRepo user.UserRepository
}

func NewFindByUserUseCase (repo user.UserRepository) *FindByIdUserUseCase {
	return &FindByIdUserUseCase{
		userRepo: repo,
	}
}

func (uf *FindByIdUserUseCase) Execute(ctx context.Context, input FindByIdUserInput) (*FindByIdUserOutput, error) {
	id := strings.TrimSpace(input.Id)

	if id == "" {
		return nil, errors.New("Nenhum id inserido")
	}

	found, err := uf.userRepo.FindById(ctx, id)

	if err != nil {
		return nil, err
	}

	if found == nil {
		return nil, errors.New("Usuario nao encontrado")
	}

	return &FindByIdUserOutput {
		Id: found.Id,
		Name: found.Name,
		Email: found.Email,
	}, nil
}