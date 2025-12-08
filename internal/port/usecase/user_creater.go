package usecase

import (
	"context"

	"github.com/FinalProjectTMS/authservice/internal/domain"
)

type UserCreater interface {
	CreateUser(ctx context.Context, user domain.User) (err error)
}
