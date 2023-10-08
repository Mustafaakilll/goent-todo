package repository

import (
	"context"

	"github.com/mustafaakilll/ent_todo/auth"
	"github.com/mustafaakilll/ent_todo/ent"
	"github.com/mustafaakilll/ent_todo/ent/user"
	"golang.org/x/crypto/bcrypt"
)

type UserRepository struct {
	Client *ent.Client
}

func NewUserRepository(client *ent.Client) *UserRepository {
	return &UserRepository{
		Client: client,
	}
}

func (r UserRepository) RegisterUser(user *ent.User) (*ent.User, error) {
	ctx := context.Background()
	hashedPassword, err := HashPassword(user.Password)

	if err != nil {
		return nil, err
	}

	newUser, err := r.Client.User.Create().
		SetEmail(user.Email).
		SetFullname(user.Fullname).
		SetPassword(hashedPassword).
		Save(ctx)

	if err != nil {
		return nil, err
	}

	_, err = auth.GenerateJWT(user.Email, user.Fullname, user.ID)
	if err != nil {
		return nil, err
	}

	return newUser, nil
}

func (r UserRepository) LoginUser(user *ent.User) (*ent.User, error) {
	oldUser, err := r.GetUserByEmail(user.Email)
	if err != nil {
		return nil, err
	}

	err = CheckPassword(user.Password, oldUser.Password)
	if err != nil {
		return nil, err
	}

	_, err = auth.GenerateJWT(user.Email, user.Fullname, user.ID)
	if err != nil {
		return nil, err
	}

	return oldUser, nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func CheckPassword(incomingPassword, password string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(password), []byte(incomingPassword)); err != nil {
		return err
	}
	return nil
}

func (r UserRepository) GetUserByEmail(email string) (*ent.User, error) {
	ctx := context.Background()
	user, err := r.Client.User.Query().Where(user.Email(email)).First(ctx)

	if err != nil {
		return nil, err
	}

	return user, nil
}
