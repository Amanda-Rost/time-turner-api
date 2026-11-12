package user

import (
	"context"
	"errors"
	"strings"
	"time"
	"time-turner-api/internal/domain/user"
	"unicode"

	"golang.org/x/crypto/bcrypt"
)

// ============================================================================
// PATTERN 1: DTO (Data Transfer Object)
// ============================================================================
// O QUE É: Objeto anêmico que carrega dados entre processos sem conter lógica.
//
// POR QUE FOI ESCOLHIDO ESTE PATTERN (vs ALTERNATIVAS):
// ❌ Alternativa Descartada A (Usar Entidade do Domínio 'user.User'):
//    Vazaria campos como Password, ID gerado e DeletedAt para quem chama a API.
//    Também causaria risco de "Mass Assignment", permitindo que um atacante envie
//    campos protegidos no JSON de entrada.
// ❌ Alternativa Descartada B (Usar map[string]interface{}):
//    Perderia a checagem de tipos estáticos do Go em tempo de compilação,
//    deixando o código propenso a erros de runtime.
//
// MOTIVAÇÃO DA ESCOLHA:
// O DTO cria um contrato estrito. O 'Input' aceita APENAS o necessário para a operação,
// e o 'Output' devolve APENAS os dados sanitizados para o cliente HTTP.

type CreateUserInput struct {
	Name     string `json:"name"`
	Email 	 string `json:"email"`
	Password string `json:"password"`
}

type CreateUserOutput struct {
	Id        string 	`json:"id"`
	Name      string 	`json:"name"`
	Email     string 	`json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

// ============================================================================
// PATTERN 2: Command Pattern / Single-Purpose Use Case
// ============================================================================
// O QUE É: Encapsular uma ação (um comando do negócio) dentro de um objeto dedicado.
//
// POR QUE FOI ESCOLHIDO ESTE PATTERN (vs ALTERNATIVAS):
// ❌ Alternativa Descartada (Service Layer Monolítico - ex: 'UserService'):
//    O padrão clássico 'UserService' agrupa dezenas de métodos (Create, Update, Delete, 
//    SendEmail, ResetPassword) em uma única struct gigantesca. Isso viola o SRP 
//    (Single Responsibility Principle), infla as dependências e gera conflitos no Git.
//
// MOTIVAÇÃO DA ESCOLHA:
// Cada Use Case vira uma struct independente com apenas UM método 'Execute()'.
// Escolhemos isso porque facilita a leitura, permite testar cada caso de uso isoladamente
// e garante que a struct dependa APENAS dos repositórios que ela realmente usa.

type CreateUserUseCase struct {
	// ============================================================================
	// PRINCÍPIO SOLID: DIP (Dependency Inversion Principle)
	// ============================================================================
	// O QUE É: O módulo de alto nível (Use Case) depende de abstrações (Interfaces),
	// e não de módulos de baixo nível (GORM/Postgres).
	//
	// POR QUE FOI ESCOLHIDO (vs ALTERNATIVAS):
	// ❌ Alternativa Descartada (Depender da struct 'db.UserGormRepository'):
	//    Acoplaria a regra de negócio ao PostgreSQL e ao ORM GORM. Se amanhã o banco
	//    mudasse para MongoDB ou DynamoDB, teríamos que reescrever a aplicação inteira.
	//
	// MOTIVAÇÃO DA ESCOLHA:
	// Depender da interface 'user.Repository' nos permite injetar um repositório mock
	// em memória nos testes unitários, executando suítes de teste em milissegundos.

	userRepo user.UserRepository
}

// ============================================================================
// PATTERN 3: Factory Method (Constructor Function Idiomático em Go)
// ============================================================================
// O QUE É: Função encarregada de fabricar e inicializar uma instância da struct.
//
// POR QUE FOI ESCOLHIDO ESTE PATTERN (vs ALTERNATIVAS):
// ❌ Alternativa Descartada (Instanciação direta: '&CreateUserUseCase{userRepo: repo}'):
//    Permite que o chamador esqueça de passar o repositório, resultando em uma struct 
//    com ponteiro nulo (nil) que causaria um crash ('panic: runtime error') na aplicação.
//
// MOTIVAÇÃO DA ESCOLHA:
// Como o Go não possui construtores nativos como Java, a função 'New...' garante o encapsulamento: 
// o campo 'userRepo' permanece privado (com letra minúscula) e o Use Case NUNCA é criado
// em um estado inválido ou sem suas dependências.

func NewCreateUserUseCase(repo user.UserRepository) *CreateUserUseCase {
	return &CreateUserUseCase{
		userRepo: repo,
	}
}

func validatePasswordFormat(password string) error {
	if len(password) < 8 {
		return errors.New("A senha deve ter no minimo 8 caracteres")
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
		return errors.New("a senha deve conter: " + strings.Join(missingItems, ", "))
	}

	return nil
}

// O (uc *CreateUserUseCase) é o jeito de mostrar que esse metodo pertece aquela classe
func (uc *CreateUserUseCase) Execute(ctx context.Context, input CreateUserInput) (*CreateUserOutput, error) {
	// O strings.TrimSpace remove os espaços em branco do início e do fim de uma string
	name := strings.TrimSpace(input.Name)
	email := strings.TrimSpace(input.Email)

	if name == "" {
		return nil, errors.New("É obrigatorio o usuario ter um nome")
	}
	if email == "" || !strings.Contains(email,"@") {
		return nil, errors.New("É obrigatorio o usuario ter um email")
	}

	if err := validatePasswordFormat(input.Password); err !=nil {
		return nil, err
	}

	existingUser, _ := uc.userRepo.FindByEmail(ctx, email) 

	if existingUser != nil {
		return nil, errors.New("Email indisponivel")
	}

	// Gera um hash seguro da senha usando o algoritmo bcrypt com sal (salt) automático 
	// e custo computacional padrão.
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	
	if err != nil {
    	return nil, err
	}

	newUser := &user.User{
		Name: name,
		Email: email,
		Password: string(hashedPassword),
	}

	err = uc.userRepo.Create(ctx, newUser)

	if err != nil {
		return nil, err
	}

	return &CreateUserOutput {
		Id: newUser.Id,
		Name: newUser.Name,
		Email: newUser.Email,
		CreatedAt: newUser.CreatedAt,
	}, nil
}

