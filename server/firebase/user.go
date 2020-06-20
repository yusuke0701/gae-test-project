package firebase

import (
	"context"
	"fmt"
	"log"

	"firebase.google.com/go/auth"
)

// CreateUser create a user on the Firebase Auth.
func CreateUser(ctx context.Context, email, password string) (*auth.UserRecord, error) {
	c, err := getAuthClient(ctx)
	if err != nil {
		return nil, err
	}

	params := (&auth.UserToCreate{}).
		Email(email).
		Password(password)

	u, err := c.CreateUser(ctx, params)
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %v", err)
	}
	log.Printf("Successfully created user: %v\n", u)
	return u, nil
}

// GetUser get a user from th Firebase Auth.
func GetUser(ctx context.Context, uid string) (*auth.UserRecord, error) {
	c, err := getAuthClient(ctx)
	if err != nil {
		return nil, err
	}

	u, err := c.GetUser(ctx, uid)
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %v", err)
	}
	log.Printf("Successfully fetched user data: %v\n", u)
	return u, nil
}
