package firebase

import (
	"context"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
)

var _ FirebaseAuth = (*firebaseAuth)(nil)

type FirebaseAuth interface {
	VerifyIDToken(ctx context.Context, idToken string) (*auth.Token, error)
}

type firebaseAuth struct {
	client *auth.Client
}

func NewAuthClient() (FirebaseAuth, error) {
	ctx := context.Background()
	app, err := firebase.NewApp(ctx, nil)
	if err != nil {
		return nil, err
	}

	client, err := app.Auth(ctx)
	if err != nil {
		return nil, err
	}

	return &firebaseAuth{client: client}, nil
}

func (f *firebaseAuth) VerifyIDToken(ctx context.Context, idToken string) (*auth.Token, error) {
	token, err := f.client.VerifyIDToken(ctx, idToken)
	if err != nil {
		return nil, err
	}

	return token, nil
}
