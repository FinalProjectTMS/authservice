package usecase

import (
	"github.com/FinalProjectTMS/authservice/internal/adapter/driven/broker"
	"github.com/FinalProjectTMS/authservice/internal/adapter/driven/dbstore"
	"github.com/FinalProjectTMS/authservice/internal/config"
	"github.com/FinalProjectTMS/authservice/internal/port/usecase"
	authenticate "github.com/FinalProjectTMS/authservice/internal/usecase/authenticator"
	usercreater "github.com/FinalProjectTMS/authservice/internal/usecase/user_creater"
)

type UseCases struct {
	UserCreater   usecase.UserCreater
	Authenticator usecase.Authenticate
}

func New(cfg config.Config, store *dbstore.DBStore, publisher *broker.MessagePublisher) *UseCases {
	return &UseCases{
		UserCreater:   usercreater.New(&cfg, store.UserStorage, publisher),
		Authenticator: authenticate.New(&cfg, store.UserStorage),
	}
}
