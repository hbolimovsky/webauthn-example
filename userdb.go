package main

import (
	"fmt"
	"sync"
)

type userdb struct {
	users map[string]*User
	mu    sync.RWMutex
}

var db *userdb

// DB returns a userdb singleton
func DB() *userdb {

	if db == nil {
		db = &userdb{
			users: make(map[string]*User),
		}
	}

	return db
}

// GetUser returns a *User by the user's username
func (db *userdb) GetUser(name string) (*User, error) {

	db.mu.Lock()
	defer db.mu.Unlock()
	user, ok := db.users[name]
	if !ok {
		return &User{}, fmt.Errorf("error getting user '%s': does not exist", name)
	}

	return user, nil
}

// PutUser stores a new user by the user's username
func (db *userdb) PutUser(user *User) {

	db.mu.Lock()
	defer db.mu.Unlock()
	db.users[user.name] = user
}
