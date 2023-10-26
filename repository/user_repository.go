// Database operations for user and todo.
package repository

import (
	"context"
	"errors"

	"github.com/mustafaakilll/ent_todo/auth"
	"github.com/mustafaakilll/ent_todo/ent"
	"github.com/mustafaakilll/ent_todo/ent/user"
	"golang.org/x/crypto/bcrypt"
)

// UserRepository struct for accessing Ent.Client
type UserRepository struct {
	// Ent Client
	Client *ent.Client
}

// NewUserRepository function for creating new UserRepository with client for DI
func NewUserRepository(client *ent.Client) *UserRepository {
	return &UserRepository{
		Client: client,
	}
}

// GetUserById function for register user to db and GenerateJWT.
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

	_, err = auth.GenerateJWT(user.ID)
	if err != nil {
		return nil, err
	}

	return newUser, nil
}

// GetUserById function for log in to existing user.
// If there is no user or credentials are wrong, return error.
func (r UserRepository) LoginUser(user *ent.User) (*ent.User, error) {
	oldUser, err := r.GetUserByEmail(user.Email)

	if err != nil {
		return nil, err
	}

	err = CheckPassword(user.Password, oldUser.Password)
	if err != nil {
		return nil, err
	}

	_, err = auth.GenerateJWT(user.ID)
	if err != nil {
		return nil, err
	}

	return oldUser, nil
}

// GetUserByEmail function for getting user by email.
// If there is no user with email, return "User not found" error.
func (r UserRepository) GetUserByEmail(email string) (*ent.User, error) {
	ctx := context.Background()
	user, err := r.Client.User.Query().Where(user.Email(email)).First(ctx)

	switch {
	case ent.IsNotFound(err):
		return nil, errors.New("User not found")
	case err != nil:
		return nil, err
	}

	return user, nil
}

// HashPassword function for hashing password before saving password to db.
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// CheckPassword function for checking password is correct or not.
func CheckPassword(incomingPassword, password string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(password), []byte(incomingPassword)); err != nil {
		return err
	}
	return nil
}
