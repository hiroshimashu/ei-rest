package usecases

import (
	"reflect"
	"testing"

	"github.com/hiroshimashu/ei-rest/app/domain"
	databse "github.com/hiroshimashu/ei-rest/app/interfaces/database"
)

func TestFindAllUsers(t *testing.T) {
	var want domain.Users
	repository := databse.NewImMemoryUserRepositry()
	t.Run("returns users", func(t *testing.T) {
		got, err := repository.FindAll()
		if err != nil {
			t.Errorf("Failed to get users")
		}
		want = []domain.User{
			{
				ID:       "5555",
				Gender:   "M",
				Age:      26,
				Area:     "tokyo",
				Email:    "example1@g.com",
				Password: "1111",
				IsPaid:   true,
			},
			{
				ID:       "4444",
				Gender:   "W",
				Age:      28,
				Area:     "kyoto",
				Email:    "example2@g.com",
				Password: "2222",
				IsPaid:   false,
			},
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}

	})
}

func TestPostUser(t *testing.T) {
	repository := databse.NewImMemoryUserRepositry()
	t.Run("Correctly user has been posted", func(t *testing.T) {
		user := domain.User{
			ID:       "7777",
			Gender:   "M",
			Age:      26,
			Area:     "tokyo",
			Email:    "example3@g.com",
			Password: "1111",
			IsPaid:   true,
		}
		got, err := repository.Store(user)
		if err != nil {
			t.Errorf("Failed to post user")
		}
		want := domain.User{
			ID:       "7777",
			Gender:   "M",
			Age:      26,
			Area:     "tokyo",
			Email:    "example3@g.com",
			Password: "1111",
			IsPaid:   true,
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}

	})
}
