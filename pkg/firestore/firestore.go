package firestore

import (
	"context"
	"fmt"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/option"
	"nextbasis-service-v-0.1/db/repository/models"
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

// CreateClient ...
func (cred Credential) CreateClientWithCredentialFile() (*firestore.Client, error) {
	opt := option.WithCredentialsFile((cred.Key))
	client, err := firestore.NewClient(context.Background(), cred.DB, opt)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	// Close client when done with
	// defer client.Close()
	return client, nil
}

// IFireStore ...
type IFireStore interface {
	GetData(context.Context, string) ([]models.FireStoreUser, error)
	UpdateData(context.Context, models.FireStoreUser, string) (models.FireStoreUser, error)
	// Delete(bucketName, objectName string) error
}

type firestoreModel struct {
	Client *firestore.Client
}

const defaultDuration = 15

// NewFireStoreModel ...
func NewFireStoreModel(client *firestore.Client) IFireStore {
	return &firestoreModel{Client: client}
}

func (model firestoreModel) GetData(c context.Context, fcollections string) (res []models.FireStoreUser, err error) {
	userList := model.Client.Collection(fcollections).Where("dbsync", "==", "1")
	docsnap := userList.Documents(c)

	dataMap, err := docsnap.GetAll()

	var userData models.FireStoreUser

	for _, ds := range dataMap {
		if err := ds.DataTo(&userData); err != nil {
		}
		res = append(res, userData)
	}

	return res, err
}

func (model firestoreModel) UpdateData(c context.Context, data models.FireStoreUser, fcollection string) (res models.FireStoreUser, err error) {
	userList := model.Client.Collection(fcollection)
	user := userList.Doc(data.UID)
	// _, err = user.Set(c, models.FireStoreUser{
	// 	ID:     data.ID,
	// 	Name:   data.Name,
	// 	UID:    data.UID,
	// 	DBSync: "2",
	// })

	_, err = user.Update(c, []firestore.Update{{Path: "dbsync", Value: "2"}})

	return res, err
}
