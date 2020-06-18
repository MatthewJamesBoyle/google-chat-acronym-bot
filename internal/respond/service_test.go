package respond_test

import (
	"errors"
	"fmt"
	"github.com/matthewjamesboyle/google-chat-acronym-bot/internal/persist"
	"github.com/matthewjamesboyle/google-chat-acronym-bot/internal/respond"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

type nilTestDb struct{}

func (n nilTestDb) LookupPhrase(string) (string, error) {
	panic("implement me")
}

func (n nilTestDb) AddPhrase(string, string) error {
	panic("implement me")
}

func TestNewService(t *testing.T) {
	t.Run("error given an nil db", func(t *testing.T) {
		s, err := respond.NewService(nil)

		assert.Nil(t, s)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "db cannot be nil")
	})

	t.Run("given a valid db, a service is returned", func(t *testing.T) {
		s, err := respond.NewService(nilTestDb{})
		require.NoError(t, err)

		assert.NotNil(t, s)
	})
}

type testDb struct {
	err error
	res string
}

func (t testDb) LookupPhrase(string) (string, error) {
	return t.res, t.err
}

func (t testDb) AddPhrase(string, string) error {
	return t.err
}

func TestService_Respond(t *testing.T) {
	t.Run("error case given empty phrase", func(t *testing.T) {
		//given
		s := respond.Service{}

		//when
		res := s.Respond("")

		//then
		assert.Equal(t, res, "whoops, looks like you sent me an empty phrase by accident.")
	})

	t.Run("given I ask for help, I receive the help response", func(t *testing.T) {
		s := respond.Service{}
		validHelpPhrase := "help"
		res := s.Respond(validHelpPhrase)

		assert.Equal(t, res, "Hi! Try saying explain WWW to me."+
			"If I dont support a phrase you think would be useful for others, you can add it with the add command.")
	})

	t.Run("given I add a never before added acronym, I receive no error", func(t *testing.T) {
		validAddKeyword := "add"
		acctoAdd := "hello"
		validDeff := "a friendly greeting"
		i := persist.NewInMemoryPersister()

		s, err := respond.NewService(i)
		require.NoError(t, err)

		res := s.Respond(fmt.Sprintf("%s %s %s", validAddKeyword, acctoAdd, validDeff))
		assert.Equal(t, res, "Thanks, I added that.")

		m := persist.GetInMemoryDbForTest(*i)
		v := m[acctoAdd]

		assert.Equal(t, validDeff, v)
	})

	t.Run("given I add a phrase that already exists, I get an error", func(t *testing.T) {
		validAddKeyword := "add"
		acctoAdd := "gc"
		validDeff := "garbage collection"
		i := persist.NewInMemoryPersister()

		s, err := respond.NewService(i)
		require.NoError(t, err)

		res := s.Respond(fmt.Sprintf("%s %s %s", validAddKeyword, acctoAdd, validDeff))
		assert.Equal(t, res, "Thanks, I added that.")

		errString := s.Respond(fmt.Sprintf("%s %s %s", validAddKeyword, acctoAdd, validDeff))
		assert.Equal(t, errString, "I already have a definition for that acronym and cannot store multiple. sorry about that!")
	})

	t.Run("Given I try to add and get an unexpected error from my db, I get an error string", func(t *testing.T) {
		unexpectedErr := errors.New("some-error")
		errDb := testDb{
			err: unexpectedErr,
		}
		s, err := respond.NewService(errDb)
		require.NoError(t, err)

		res := s.Respond("add some-phrase")
		assert.Equal(t, "something went wrong, Please try again!", res)
	})

	t.Run("given I ask for a definition that has not been added, I get back an error phrase", func(t *testing.T) {
		s, err := respond.NewService(persist.NewInMemoryPersister())
		require.NoError(t, err)

		res := s.Respond("explain something-you-dont-have")
		assert.Equal(t, res, "I don't have a definition for that acronym unfortunately!"+
			"When you find out please come back and add it with the add command to benefit future employees")
	})

	t.Run("given an unexpected error at the db layer, I get back and error phrase", func(t *testing.T) {
		unexpectedErr := errors.New("some-error")
		errDb := testDb{
			err: unexpectedErr,
		}
		s, err := respond.NewService(errDb)
		require.NoError(t, err)

		res := s.Respond("explain something")
		assert.Equal(t, res, "something went wrong, Please try again!")
	})

	t.Run("given I ask for an acronym that has been added before, I get back the definition", func(t *testing.T) {
		acctoExplain := "internet"
		validDeff := "the whole world's knowledge is here,as are memes"
		i := persist.NewInMemoryPersister()

		s, err := respond.NewService(i)
		require.NoError(t, err)

		res := s.Respond(fmt.Sprintf("add %s %s", acctoExplain, validDeff))
		require.Equal(t, res, "Thanks, I added that.")

		res = s.Respond(fmt.Sprintf("explain %s", acctoExplain))
		assert.Equal(t, validDeff, res)
	})

	t.Run("given you don't use a valid keyword, you get the help response", func(t *testing.T) {
		s := respond.Service{}
		res := s.Respond("hi")

		assert.Equal(t, "Hi! Try saying explain WWW to me."+
			"If I dont support a phrase you think would be useful for others, you can add it with the add command.", res)
	})

}
