package user

import (
	"context"
	"errors"
	"strings"
	"time"
	"time-turner-api/internal/domain/user"

	"golang.org/x/crypto/bcrypt"
)

type UpdateUserInput struct {
	Id       string  `json:"id"`
	Name     *string `json:"name"`
	Email    *string `json:"email"`
	Password *string `json:"password"`
}

type UpdateUserOutput struct {
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UpdateUserUseCase struct {
	userRepo user.UserRepository
}

func NewUpdateUserUseCase (repo user.UserRepository) *UpdateUserUseCase {
	return &UpdateUserUseCase{
		userRepo: repo,
	}
}

func (uu *UpdateUserUseCase) Execute(ctx context.Context, input UpdateUserInput) (*UpdateUserOutput, error) {
	id := strings.TrimSpace(input.Id)
	if id == "" {
		return nil, errors.New("o id do usuario é obrigatório")
	}

	// ============================================================================
	// PATTERN: Read-Modify-Write (Carregar Estado Atual antes de Mutar)
	// ============================================================================
	// POR QUÊ? Garante que o usuário existe e preserva os campos que não foram
	// enviados na requisição de atualização.
	existingUser, err := uu.userRepo.FindById(ctx, id)
	if err != nil {
		return nil, err
	}

	if existingUser == nil {
		return nil, errors.New("usuário não encontrado")
	}

	if input.Name != nil {
		name := strings.TrimSpace(*input.Name)

		if name == ""{
			return nil, errors.New("O nome não pode ficar em branco")
		}

		existingUser.Name = name 
	}
	
	if input.Email != nil {
		email := strings.TrimSpace(strings.ToLower(*input.Email))
		if email == "" || !strings.Contains(email, "@") {
			return nil, errors.New("email inválido")
		}

		if email != existingUser.Email {
			userWithSameEmail, _ := uu.userRepo.FindByEmail(ctx, email)
			if userWithSameEmail != nil {
				return nil, errors.New("email indisponível")
			}
			existingUser.Email = email
		}
	}

	if input.Password != nil {
		validPassword, err := user.NewPassword(*input.Password)
		if err != nil {
			return nil, err
		}

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(validPassword.Value()), bcrypt.DefaultCost)
		if err != nil {
			return nil, err
		}
		existingUser.Password = string(hashedPassword)
	}

	// Executa o Update no repositório, cria a variável 'err' limitada apenas ao escopo
	// deste 'if' e, caso ocorra alguma falha, interrompe a função retornando o erro.
	if err := uu.userRepo.Update(ctx, existingUser); err != nil {
		return nil, err
	}	

	return &UpdateUserOutput{
		Id:        existingUser.Id,
		Name:      existingUser.Name,
		Email:     existingUser.Email,
		UpdatedAt: existingUser.UpdatedAt,
	}, nil
}