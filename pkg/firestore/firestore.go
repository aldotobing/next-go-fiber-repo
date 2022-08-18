package firestore

import (
	"context"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/option"
)

// Credential ...
type Credential struct {
	Key string
	DB  string
}

// CreateClient ...
func (cred Credential) CreateClient() (*firestore.Client, error) {
	opt := option.WithCredentialsJSON([]byte(cred.Key))
	client, err := firestore.NewClient(context.Background(), cred.DB, opt)
	if err != nil {
		return nil, err
	}
	// Close client when done with
	// defer client.Close()
	return client, nil
}
