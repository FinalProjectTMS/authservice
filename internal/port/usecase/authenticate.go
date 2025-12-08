package usecase

import (
	"context"

	"github.com/FinalProjectTMS/authservice/internal/domain"
)

type Authenticate interface {
	Authenticate(ctx context.Context, user domain.User) (int, domain.Role, error)
}
