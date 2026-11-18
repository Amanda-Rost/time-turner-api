package user

import (
	"errors"
	"strings"
	"unicode"
)

// ============================================================================
// PATTERN: Value Object (DDD - Domain-Driven Design)
// ============================================================================
// O QUE É: Um objeto de domínio definido pelo seu valor (sem ID) que encapsula
// suas próprias regras de validação e garante que nunca exista em estado inválido.
//
// POR QUE EVOLUIR 'validatePasswordFormat' PARA ESTE PATTERN (vs ALTERNATIVAS):
// ❌ Alternativa Descartada (Exportar 'ValidatePasswordFormat(pwd string)' no domínio):
//    Manteria a senha trafegando como 'string' comum (Primitive Obsession). Nada
//    impediria um desenvolvedor de esquecer de chamar a validação no 'UpdateUserUseCase'
//    e salvar uma senha fraca no banco.
//
// MOTIVAÇÃO DA ESCOLHA:
// Com a struct 'Password' tendo o campo 'value' privado, a única forma de criar uma
// senha válida no sistema é passando pelo construtor 'NewPassword'. Isso centraliza
// a regra de negócio para o 'Create', 'Update' e qualquer fluxo futuro.

type Password struct {
	value string
}

// Value retorna a senha em texto plano já validada para ser processada pelo bcrypt no Use Case.
func (p Password) Value() string {
	return p.value
}

// ============================================================================
// PATTERN: Factory Method (Self-Validating Constructor)
// ============================================================================
func NewPassword(password string) (Password, error) {
	if len(password) < 8 {
		return Password{}, errors.New("A senha deve ter no minimo 8 caracteres")
	}

	var (
		hasUpper     bool = false
		hasLower     bool = false
		hasNumber    bool = false
		hasSpecial   bool = false
		missingItems []string
	)

	// O range percorre a coleção entregando o índice e o valor a cada iteração, enquanto
	// o _ descarta o índice retornado para evitar erros de compilação por variáveis não
	// utilizadas.
	for _, char := range password {
		switch {
		// O unicode é um pacote nativo da biblioteca padrão do Go que fornece funções
		// para analisar a categoria e as propriedades de um caractere
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsNumber(char):
			hasNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}

	if !hasUpper {
		missingItems = append(missingItems, "pelo menos uma letra maiuscula")
	}
	if !hasLower {
		missingItems = append(missingItems, "pelo menos uma letra minuscula")
	}
	if !hasNumber {
		missingItems = append(missingItems, "pelo menos um numero")
	}
	if !hasSpecial {
		missingItems = append(missingItems, "pelo menos um caractere especial (ex: !@#$%^&*)")
	}

	if len(missingItems) > 0 {
		return Password{}, errors.New("A senha deve conter: " + strings.Join(missingItems, ", "))
	}

	return Password{value: password}, nil
}

