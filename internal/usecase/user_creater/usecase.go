package usercreater

import (
	"context"
	"errors"
	"fmt"

	"github.com/FinalProjectTMS/authservice/internal/config"
	"github.com/FinalProjectTMS/authservice/internal/domain"
	"github.com/FinalProjectTMS/authservice/internal/errs"
	"github.com/FinalProjectTMS/authservice/internal/port/driven"
	"github.com/FinalProjectTMS/authservice/utils"
)

type UseCase struct {
	cfg              *config.Config
	userStorage      driven.UserStorage
	messagePublisher driven.MessagePublisher
}

func New(cfg *config.Config, userStorage driven.UserStorage, publisher driven.MessagePublisher) *UseCase {
	return &UseCase{
		cfg:              cfg,
		userStorage:      userStorage,
		messagePublisher: publisher,
	}
}

func (u *UseCase) CreateUser(ctx context.Context, user domain.User) (err error) {
	if !utils.IsValidEmail(user.Email) {
		return errs.ErrInvalidEmailFormat
	}

	// Check if username already exists
	if _, err := u.userStorage.GetUserByUsername(ctx, user.Username); err == nil {
		return errs.ErrUsernameAlreadyExists
	} else if !errors.Is(err, errs.ErrNotfound) {
		return err
	}

	// Check if email already exists
	if _, err := u.userStorage.GetUserByEmail(ctx, user.Email); err == nil {
		return errs.ErrEmailAlreadyExists
	} else if !errors.Is(err, errs.ErrNotfound) {
		return err
	}

	user.Password, err = utils.GenerateHash(user.Password)
	if err != nil {
		return err
	}

	user.Role = domain.RoleUser

	if err = u.userStorage.CreateUser(ctx, user); err != nil {
		return err
	}

	if u.messagePublisher != nil {
		message := domain.Message{
			Recipient: user.Email,
			Subject:   fmt.Sprintf("Welcome, %s!", user.Username),
			Body:      fmt.Sprintf("Hello, %s! Your account has been created successfully.", user.Username),
		}
		if err = u.messagePublisher.PublishMessage(message); err != nil {
			// Just log the error as a demo behavior
			fmt.Println(err)
		}
	}

	return nil
}
