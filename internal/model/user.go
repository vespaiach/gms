package model

import (
	"time"

	"github.com/vespaiach/auth/internal/comtype"
)

// User model
type User struct {
	ID        uint      `json:"id"`
	FullName  string    `json:"full_name"`
	Username  string    `json:"username"`
	Hashed    string    `json:"-"`
	Email     string    `json:"email"`
	Active    int       `json:"active"`
	Verified  bool      `json:"verified"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Actions   []*Action `json:"actions"`
	Roles     []*Role   `json:"roles"`
}

// UserRepo defines user repo
type UserRepo interface {
	// GetByID gets user by user ID
	GetByID(id int64) (user *User, err error)

	// GetByEmail gets user by user's email
	GetByUsername(username string) (*User, error)

	// GetByEmail gets user by user's email
	GetByEmail(email string) (*User, error)

	// Create a new user
	Create(fullName string, username string, hashed string, email string) (*User, error)

	// Update user
	Update(id int64, fields map[string]interface{}) error

	// Query a list of users
	Query(page int, perPage int, filters map[string]interface{}, sorts map[string]comtype.SortDirection) ([]*User, int64, error)
}
