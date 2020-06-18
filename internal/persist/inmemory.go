package persist

import "fmt"

type InMemoryPersister struct {
	mem map[string]string
}

var knownPhrases = map[string]string{
	"item1": "definition1",
	"item2": "definition2",
}

func NewInMemoryPersister() *InMemoryPersister {
	return &InMemoryPersister{mem: knownPhrases}
}

func (i InMemoryPersister) LookupPhrase(tl string) (string, error) {
	if r, ok := i.mem[tl]; ok {
		return r, nil
	}
	return "", fmt.Errorf("%w: could not find phrase %s", ErrNotFound, tl)
}

func (i InMemoryPersister) AddPhrase(ta string, means string) error {
	if _, exists := i.mem[ta]; exists {
		return fmt.Errorf("%w: definition already exists for %s", ErrPhraseAlreadyExists, ta)
	}

	i.mem[ta] = means
	return nil
}
