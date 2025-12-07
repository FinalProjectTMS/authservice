package driven

import "github.com/FinalProjectTMS/authservice/internal/domain"

type MessagePublisher interface {
	PublishMessage(message domain.Message) error
}
