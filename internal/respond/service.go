package respond

import (
	"errors"
	"github.com/matthewjamesboyle/google-chat-acronym-bot/internal/persist"
	"strings"
)

type Service struct {
	db persist.Persister
}

func NewService(db persist.Persister) (*Service, error) {
	if db == nil {
		return nil, errors.New("db cannot be nil")
	}
	return &Service{db: db}, nil
}

const (
	action_add     = "add"
	action_explain = "explain"
	action_help    = "help"
)

// Supported phrases:
// * help
// * add $word $definition
// * explain $word
func (s *Service) Respond(phrase string) string {

	if len(phrase) == 0 {
		return "whoops, looks like you sent me an empty phrase by accident."
	}

	a := strings.Split(phrase, " ")
	action := strings.ToLower(a[0])

	var acc, def string
	if action == action_explain {
		acc = strings.ToLower(a[1])
	}
	if action == action_add {
		acc = strings.ToLower(a[1])
		def = strings.Join(a[2:], " ")
	}

	switch action {
	case action_help:
		return s.help()
	case action_add:
		err := s.db.AddPhrase(acc, def)
		if err != nil {
			switch {
			case errors.Is(err, persist.ErrPhraseAlreadyExists):
				return "I already have a definition for that acronym and cannot store multiple. sorry about that!"
			default:
				return "something went wrong, Please try again!"
			}
		}
		return "Thanks, I added that."

	case action_explain:
		res, err := s.db.LookupPhrase(acc)
		if err != nil {
			switch {
			case errors.Is(err, persist.ErrNotFound):
				return "I don't have a definition for that acronym unfortunately!" +
					"When you find out please come back and add it with the add command to benefit future employees"
			default:
				return "something went wrong, Please try again!"
			}
		}
		return res

	default:
		return s.help()
	}
}

func (s *Service) help() string {
	return "Hi! Try saying explain WWW to me." +
		"If I dont support a phrase you think would be useful for others, you can add it with the add command."
}
